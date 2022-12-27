package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"net"
	"os"
	"runtime"

	// "github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
)

var live bool

func CreateCircuitBreaker() func() (err error) {
	var count int
	var circuit bool
	return func() (err error) {
		limit := api.Limit()
		cpuLimit, ferr := strconv.ParseFloat(limit["cpu_limit"], 32)
		if ferr != nil || cpuLimit == 0 {
			cpuLimit = 100
			return
		}
		s, err := getServerInfo()
		var cpuNum float64
		cpus := s.Cpu.Cpus
		for _, k := range cpus {
			cpuNum += k
		}
		cpu := cpuNum / float64(len(cpus))
		if cpu > cpuLimit {
			if count >= 5 {
				circuit = true
				count = 0
				StopAgent()
				return
			}
			count++
		}
		if cpu <= cpuLimit {
			count = 0
			if circuit {
				RunAgent()
			}
		}
		return nil
	}
}

func StopAgent() {
	live = false
	hook.HookAll(global.AllHooks...)
}

func RunAgent() {
	live = true
	hook.UnHookAll(global.AllHooks...)
}

func AgentRegister() (err error) {
	max := 6
	breaker := CreateCircuitBreaker()
	live = true
	OS := runtime.GOOS
	hostname, _ := os.Hostname()
	version := "1.0.0"
	packages, agentVersion := GetMod()
	if global.Config.DongtaiGoProjectVersion != "" {
		version = global.Config.DongtaiGoProjectVersion
	}
	projectName := "project Name"
	if global.Config.DongtaiGoProjectVersion != "" {
		projectName = global.Config.DongtaiGoProjectName
	}
	name := OS + "-" + hostname + "-" + version + "-" + global.Config.DongtaiGoAgentToken
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
		Version:           agentVersion,
		ProjectVersion:    version,
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
			max = max - 1
			fmt.Printf("等待当前程序当前程序启动完成，剩余%d次", max)
			time.Sleep(1 * time.Second)
			ip, err := utils.ExternalIP()
			if err != nil {
				fmt.Println(err)
			}
			var cmd *exec.Cmd
			var strErr bytes.Buffer
			var out bytes.Buffer
			if OS == "windows" {
				cmd = exec.Command("netstat", "-ano")
				cmd.Stderr = &strErr
				cmd.Stdout = &out
			} else {
				cmd = exec.Command("netstat", "-antup", "|grep", req.Pid)
				cmd.Stderr = &strErr
				cmd.Stdout = &out
			}
			fmt.Println("等待当前程序端口返回")
			err = cmd.Run()
			fmt.Println("等待当前程序端口返回成功")
			output := out.String()
			if err != nil {
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
			if matches != nil || max == 0 {
				if matches[0] != nil || max == 0 {
					if matches[0][1] != "" || max == 0 {
						fmt.Println("当前程序启动完成")
						if max == 0 {
							req.ServerPort = ""
						} else {
							req.ServerPort = matches[0][1]
						}
						req.ServerAddr = ip.String()
						agentId, err := api.AgentRegister(req)
						if err != nil {
							fmt.Println(err)
							break
						}
						global.AgentId = agentId
						UploadSca(packages)
						go func() {
							for {
								time.Sleep(10 * time.Second)
								PingPang()
								breaker()
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
	req.Detail.Pant.IsCoreRunning = 1
	req.Detail.Pant.IsCoreInstalled = 1
	req.Detail.AgentId = global.AgentId
	api.ReportUpload(req)
}
func UploadSca(packages []request.Component) {
	var req request.UploadReq
	req.Type = 18
	req.Detail.AgentId = global.AgentId
	req.Detail.Packages = packages
	api.ReportUpload(req)
}
