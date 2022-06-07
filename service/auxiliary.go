package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/service/version"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var buildInfoMagic = []byte("\xff Go buildinf:")

// An exe is a generic interface to an OS executable (ELF, Mach-O, PE, XCOFF).
type exe interface {
	// Close closes the underlying file.
	Close() error

	// ReadData reads and returns up to size byte starting at virtual address addr.
	ReadData(addr, size uint64) ([]byte, error)

	// DataStart returns the writable data segment start address.
	DataStart() uint64
}

// GenAQLForGolang 为golang组件生成aql
func GenAQLForGolang(packageName, version string) string {
	return fmt.Sprintf("golang:%s:%s:", packageName, version)
}

// 获取包
func GetMod() ([]request.Component, string) {
	//fmt.Println(getCurrentPath())
	path, _ := os.Executable()
	return scanFile(path, true)
}

// 判断是否是exe文件
func isExe(file string, info fs.FileInfo) bool {
	if runtime.GOOS == "windows" {
		return strings.HasSuffix(strings.ToLower(file), ".exe")
	}
	return info.Mode().IsRegular() && info.Mode()&0111 != 0
}

// 从二进制文件读取包信息
func scanFile(file string, mustPrint bool) (packages []request.Component, agentVersion string) {

	i, err := os.Stat(file)
	info := i

	if !isExe(file, info) {
		if mustPrint {
			fmt.Fprintf(os.Stderr, "%s: not executable file\n", file)
		}
		return
	}

	x, err := version.OpenExe(file)
	if err != nil {
		if mustPrint {
			fmt.Fprintf(os.Stderr, "%s: %v\n", file, err)
		}
		return
	}
	defer x.Close()

	vers, mod := findVers(x)
	if vers == "" {
		if mustPrint {
			fmt.Fprintf(os.Stderr, "%s: go version not found\n", file)
		}
		return
	}

	li := strings.Split(mod[:len(mod)-1], "\n")
	for i := range li {
		licl := strings.Split(li[i], "\t")
		if licl[0] == "dep" {
			fmt.Printf("依赖:%s\t版本:%s\n", licl[1], licl[2])
			aql := GenAQLForGolang(licl[1], licl[2])
			if licl[1] == "github.com/HXSecurity/DongTai-agent-go" {
				fmt.Println("当前探针版本为：" + licl[2])
				agentVersion = licl[2]
			}
			packages = append(packages, request.Component{
				PackageName:      aql,
				PackageAlgorithm: "SHA-1",
				PackagePath:      file,
				PackageVersion:   licl[2],
				PackageSignature: utils.SHA1(aql),
			},
			)
		}
	}
	return packages, agentVersion
}

// 获取服务信息
func getServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		fmt.Println(err.Error())
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		fmt.Println(err.Error())
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		fmt.Println(err.Error())
		return &s, err
	}

	return &s, nil
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func findVers(x exe) (vers, mod string) {
	// Read the first 64kB of text to find the build info blob.
	text := x.DataStart()
	data, err := x.ReadData(text, 64*1024)
	if err != nil {
		return
	}
	for ; !bytes.HasPrefix(data, buildInfoMagic); data = data[32:] {
		if len(data) < 32 {
			return
		}
	}

	// Decode the blob.
	ptrSize := int(data[14])
	bigEndian := data[15] != 0
	var bo binary.ByteOrder
	if bigEndian {
		bo = binary.BigEndian
	} else {
		bo = binary.LittleEndian
	}
	var readPtr func([]byte) uint64
	if ptrSize == 4 {
		readPtr = func(b []byte) uint64 { return uint64(bo.Uint32(b)) }
	} else {
		readPtr = bo.Uint64
	}
	vers = readString(x, ptrSize, readPtr, readPtr(data[16:]))
	if vers == "" {
		return
	}
	mod = readString(x, ptrSize, readPtr, readPtr(data[16+ptrSize:]))
	if len(mod) >= 33 && mod[len(mod)-17] == '\n' {
		// Strip module framing.
		mod = mod[16 : len(mod)-16]
	} else {
		mod = ""
	}
	return
}

// readString returns the string at address addr in the executable x.
func readString(x exe, ptrSize int, readPtr func([]byte) uint64, addr uint64) string {
	hdr, err := x.ReadData(addr, uint64(2*ptrSize))
	if err != nil || len(hdr) < 2*ptrSize {
		return ""
	}
	dataAddr := readPtr(hdr)
	dataLen := readPtr(hdr[ptrSize:])
	data, err := x.ReadData(dataAddr, dataLen)
	if err != nil || uint64(len(data)) < dataLen {
		return ""
	}
	return string(data)
}
