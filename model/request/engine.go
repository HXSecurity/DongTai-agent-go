package request

type AgentRegisterReq struct {
	Name             string `json:"name,omitempty"`
	Language         string `json:"language,omitempty"`
	Version          string `json:"version,omitempty"`
	ProjectName      string `json:"projectName,omitempty"`
	Hostname         string `json:"hostname,omitempty"`
	Network          string `json:"network,omitempty"`
	ContainerName    string `json:"container_name,omitempty"`
	ContainerVersion string `json:"container_version,omitempty"`
	ServerAddr       string `json:"serverAddr,omitempty"`
	ServerPort       string `json:"serverPort,omitempty"`
	ServerPath       string `json:"serverPath,omitempty"`
	ServerEnv        string `json:"serverEnv,omitempty"`
	Pid              string `json:"pid,omitempty"`
}

type HookRuleReq struct {
	Language string `json:"language,omitempty"`
}
