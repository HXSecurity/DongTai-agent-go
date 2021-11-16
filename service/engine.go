package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go-agent/api"
	"go-agent/global"
	"go-agent/utils"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	//"go-agent/api"
	"go-agent/model/request"
	"net"
	"os"
	"runtime"
)

func AgentRegister() (err error) {
	OS := runtime.GOOS
	hostname, _ := os.Hostname()
	version := "v1.0"
	id := xid.New().String()
	name := OS + "-" + hostname + "-" + version + "-" + id
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	ips := ""

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			return err
		}
		addresses, err := byName.Addrs()
		for _, v := range addresses {
			if ips == "" {
				ips = "{\"name\":" + byName.Name + ",\"ip\":" + v.String() + "}"
			} else {
				ips += ",{\"name\":" + byName.Name + ",\"ip\":" + v.String() + "}"
			}
		}
	}
	pid := os.Getpid()
	envMap := make(map[string]string)
	envs := os.Environ()
	for _, v := range envs {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			envMap[parts[0]] = parts[1]
		}
	}
	envB, err := json.Marshal(envMap)
	if err != nil {
		return err
	}
	encodeEnv := base64.StdEncoding.EncodeToString(envB)

	filePath, err := getCurrentPath()
	if err != nil {
		return err
	}
	req := request.AgentRegisterReq{
		Name:             name,
		Language:         "GO",
		Version:          version,
		ProjectName:      "testProject",
		Hostname:         hostname,
		Network:          ips,
		ContainerName:    "GO",
		ContainerVersion: "GO",
		ServerAddr:       "",
		ServerPort:       "",
		ServerPath:       filePath,
		ServerEnv:        encodeEnv,
		Pid:              strconv.Itoa(pid),
	}

	go func() {
		for {
			fmt.Println("等待当前程序http启动完成")
			time.Sleep(1 * time.Second)
			ip, err := utils.ExternalIP()
			if err != nil {
				fmt.Println(err)
			}
			cmd := exec.Command("netstat", "-antup", "|grep", req.Pid)
			output, err := cmd.Output()
			if err != nil {
				return
			}
			r := regexp.MustCompile(`:::\s*(.*?)\s* `)
			matches := r.FindAllStringSubmatch(string(output), -1)
			if matches[0] != nil {
				if matches[0][1] != "" {
					req.ServerPort = matches[0][1]
					req.ServerAddr = ip.String()
					agentId, err := api.AgentRegister(req)
					if err != nil {
						fmt.Println(err)
						break
					}
					global.AgentId = agentId
					go func() {
						for {
							time.Sleep(5 * time.Second)
							PingPang()
						}
					}()
					break
				}
			}
		}
	}()
	return nil
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

func PingPang() {
	s, err := getServerInfo()
	if err != nil {
		return
	}
	fmt.Println(s)
	var req request.UploadReq

	cpuMap := make(map[string]string)
	memoryMap := make(map[string]string)
	var cpus float64 = 0
	for _, v := range s.Cpu.Cpus {
		cpus += v
	}
	cpuRate := fmt.Sprintf("%.2f", cpus)
	memoryRate := fmt.Sprintf("%.2f", float64(s.Rrm.UsedMB)/float64(s.Rrm.TotalMB))
	cpuMap["rate"] = cpuRate
	memoryMap["total"] = strconv.Itoa(s.Rrm.TotalMB) + "MB"
	memoryMap["use"] = strconv.Itoa(s.Rrm.UsedMB) + "MB"
	memoryMap["rate"] = memoryRate

	cpuByte, _ := json.Marshal(cpuMap)
	memoryByte, _ := json.Marshal(memoryMap)

	req.Type = 1
	req.Detail.Pant.Disk = "{}"
	req.Detail.Pant.Cpu = string(cpuByte)
	req.Detail.Pant.Memory = string(memoryByte)
	req.Detail.AgentId = global.AgentId
	api.ReportUpload(req)
}

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

func UploadMethodCall(interface{}) {

}
