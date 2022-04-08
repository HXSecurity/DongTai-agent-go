package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	//"github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"net"
	"os"
	"runtime"
)

func AgentRegister() (err error) {
	OS := runtime.GOOS
	hostname, _ := os.Hostname()
	version := "1.0.0"
	if global.Config.DongtaiGoProjectVersion != "" {
		version = global.Config.DongtaiGoProjectVersion
	}
	projectName := "project Name"
	if global.Config.DongtaiGoProjectVersion != "" {
		projectName = global.Config.DongtaiGoProjectName
	}
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
		AutoCreateProject: global.Config.DongtaiGoProjectCreate,
		Name:              name,
		Language:          "GO",
		Version:           version,
		ProjectName:       projectName,
		Hostname:          hostname,
		Network:           ips,
		ContainerName:     "GO",
		ContainerVersion:  "GO",
		ServerAddr:        "",
		ServerPort:        "",
		ServerPath:        filePath,
		ServerEnv:         encodeEnv,
		Pid:               strconv.Itoa(pid),
	}
	go func() {
		for {
			fmt.Println("等待当前程序http启动完成")
			time.Sleep(1 * time.Second)
			ip, err := utils.ExternalIP()
			if err != nil {
				fmt.Println(err)
			}
			var cmd *exec.Cmd
			var strErr bytes.Buffer
			var out bytes.Buffer
			fmt.Println(req.Pid)
			if OS == "windows" {
				cmd = exec.Command("netstat", "-ano")
				cmd.Stderr = &strErr
				cmd.Stdout = &out
			} else {
				cmd = exec.Command("netstat", "-antup", "|grep", req.Pid)
				cmd.Stderr = &strErr
				cmd.Stdout = &out
			}
			err = cmd.Run()
			output := out.String()
			if err != nil {
				netstatOperate()
				return
			}
			var matches [][]string
			if OS == "windows" {
				str := ""
				for {
					line, setErr := out.ReadBytes('\n')
					if setErr != nil {
						break
					}
					if strings.Index(string(line), " "+req.Pid) > -1 {
						str += string(line) + "\n"
					}
				}
				r := regexp.MustCompile(`0.0.0.0:\s*(.*?)\s* `)
				matches = r.FindAllStringSubmatch(str, -1)
			} else {
				r := regexp.MustCompile(`:::\s*(.*?)\s* `)
				matches = r.FindAllStringSubmatch(output, -1)
			}
			if matches != nil {
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

		}
	}()
	return nil
}

// 判断如果有netstat则继续等待，否则就是安装，安装失败则panic，提醒用户。
func netstatOperate() {
	var look = func(add func()) {
		cmd := exec.Command("netstat", "-e")
		err := cmd.Run()
		if err != nil {
			add()
		}
	}
	var add = func() {
		install := make(map[string][]string)
		install["apt-get"] = []string{"install", "net-tools"}
		install["yum"] = []string{"-y", "install", "net-tools"}

		var err error

		var execCommand = func(tool string, params []string) {
			cmd := exec.Command(tool, params...) // 拼接命令
			err = cmd.Start()
		}

		for key, value := range install {
			execCommand(key, value)
			if err == nil {
				return
			}
		}
		if err != nil {
			panic("请自行安装netstat工具")
		}
	}
	look(add)
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
	var req request.UploadReq

	cpuMap := make(map[string]string)
	memoryMap := make(map[string]string)
	var cpus float64 = 0
	for _, v := range s.Cpu.Cpus {
		cpus += v
	}
	cpuRate := fmt.Sprintf("%.2f", cpus/float64(len(s.Cpu.Cpus)))
	memoryRate := fmt.Sprintf("%.2f", float64(s.Rrm.UsedMB)/float64(s.Rrm.TotalMB))
	total := strconv.Itoa(s.Rrm.TotalMB) + "MB"
	use := strconv.Itoa(s.Rrm.UsedMB) + "MB"
	cpuMap["rate"] = cpuRate
	memoryMap["total"] = total
	memoryMap["use"] = use
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
