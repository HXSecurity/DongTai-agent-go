//go:build go1.18
// +build go1.18

package service

import (
	"debug/buildinfo"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

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
	bi, err := buildinfo.ReadFile(file)
	if err != nil {
		if mustPrint {
			if pathErr := (*os.PathError)(nil); errors.As(err, &pathErr) && filepath.Clean(pathErr.Path) == filepath.Clean(file) {
				fmt.Fprintf(os.Stderr, "%v\n", file)
			} else {
				fmt.Fprintf(os.Stderr, "%s: %v\n", file, err)
			}
		}
		return packages, agentVersion
	}
	fmt.Printf("%s: %s\n", file, bi.GoVersion)
	bi.GoVersion = "" // suppress printing go version again
	mod := bi.String()

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
