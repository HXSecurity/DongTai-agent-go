package request

import "fmt"

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

// UploadReq
/* Type
数据类型，可选择：
1 - Agent心跳数据
17 - 依赖组件数据
36 - 方法调用数据
*/
type UploadReq struct {
	Type     int    `json:"type,omitempty"`
	Detail   Detail `json:"detail,omitempty"`
	InvokeId int    `json:"invoke_id"`
}

type Detail struct {
	AgentId int `json:"agentId"`
	Pant
	Component
	Function
	Pool
	Log
	Api
}

type Pant struct {
	Disk        string `json:"disk,omitempty"`
	Memory      string `json:"memory,omitempty"`
	Cpu         string `json:"cpu,omitempty"`
	MethodQueue int    `json:"methodQueue"`
	ReplayQueue int    `json:"replayQueue"`
	ReqCount    int    `json:"reqCount"`
	ReportQueue int    `json:"reportQueue"`
}

type Component struct {
	PackagePath      string `json:"packagePath,omitempty"`
	PackageSignature string `json:"packageSignature,omitempty"`
	PackageName      string `json:"packageName,omitempty"`
	PackageAlgorithm string `json:"packageAlgorithm,omitempty"`
}

type Function struct {
	Uri           string `json:"uri,omitempty"`
	Url           string `json:"url,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	ContextPath   string `json:"contextPath,omitempty"`
	Pool          []Pool `json:"pool,omitempty"`
	Language      string `json:"language,omitempty"`
	ClientIp      string `json:"clientIp,omitempty"`
	Secure        bool   `json:"secure,omitempty"`
	QueryString   string `json:"queryString,omitempty"`
	ReplayRequest bool   `json:"replayRequest,omitempty"`
	Method        string `json:"method,omitempty"`
	ReqHeader     string `json:"reqHeader,omitempty"`
	ReqBody       string `json:"reqBody,omitempty"`
	ResBody       string `json:"resBody,omitempty"`
	Scheme        string `json:"scheme,omitempty"`
	ResHeader     string `json:"resHeader,omitempty"`
}

type Pool struct {
	InvokeId         int           `json:"invokeId,omitempty"`
	Interfaces       []interface{} `json:"interfaces,omitempty"`
	TargetHash       []string      `json:"targetHash,omitempty"`
	TargetValues     string        `json:"targetValues,omitempty"`
	Signature        string        `json:"signature,omitempty"`
	OriginClassName  string        `json:"originClassName,omitempty"`
	SourceValues     string        `json:"sourceValues,omitempty"`
	MethodName       string        `json:"methodName,omitempty"`
	ClassName        string        `json:"className,omitempty"`
	Source           bool          `json:"source,omitempty"`
	CallerLineNumber int           `json:"callerLineNumber,omitempty"`
	CallerClass      string        `json:"callerClass,omitempty"`
	Args             string        `json:"args,omitempty"`
	CallerMethod     string        `json:"callerMethod,omitempty"`
	SourceHash       []string      `json:"sourceHash,omitempty"`
	RetClassName     string        `json:"retClassName,omitempty"`
}

type Log struct {
	Log string `json:"log,omitempty"`
}

type Api struct {
	ApiData []ApiData `json:"apiData,omitempty"`
}

type ApiData struct {
	Uri         string      `json:"uri,omitempty"`
	Method      []string    `json:"method,omitempty"`
	Class       string      `json:"class,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
	ReturnType  string      `json:"returnType,omitempty"`
	File        string      `json:"file,omitempty"`
	Controller  string      `json:"controller,omitempty"`
	Description string      `json:"description,omitempty"`
}

type Parameter struct {
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Annotation string `json:"annotation,omitempty"`
}

type PoolTree struct {
	*Pool
	Begin       bool
	GoroutineID string
	Children    []*PoolTree
}

func (p *PoolTree) IsThisBegin(GoroutineID string) bool {
	return GoroutineID == p.GoroutineID && p.Begin
}

func (p *PoolTree) FMT() {
	fmt.Println(p.Pool.SourceHash, "===>", p.Pool.TargetHash)
	for k, _ := range p.Children {
		p.Children[k].FMT()
	}
}
