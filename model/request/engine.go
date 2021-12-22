package request

import "fmt"

type AgentRegisterReq struct {
	Name             string `json:"name"`
	Language         string `json:"language"`
	Version          string `json:"version"`
	ProjectName      string `json:"projectName"`
	Hostname         string `json:"hostname"`
	Network          string `json:"network"`
	ContainerName    string `json:"container_name"`
	ContainerVersion string `json:"container_version"`
	ServerAddr       string `json:"serverAddr"`
	ServerPort       string `json:"serverPort"`
	ServerPath       string `json:"serverPath"`
	ServerEnv        string `json:"serverEnv"`
	Pid              string `json:"pid"`
}

type HookRuleReq struct {
	Language string `json:"language"`
}

// UploadReq
/* Type
数据类型，可选择：
1 - Agent心跳数据
17 - 依赖组件数据
36 - 方法调用数据
*/
type UploadReq struct {
	Type     int    `json:"type"`
	Detail   Detail `json:"detail"`
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
	Disk        string `json:"disk"`
	Memory      string `json:"memory"`
	Cpu         string `json:"cpu"`
	MethodQueue int    `json:"methodQueue"`
	ReplayQueue int    `json:"replayQueue"`
	ReqCount    int    `json:"reqCount"`
	ReportQueue int    `json:"reportQueue"`
}

type Component struct {
	PackagePath      string `json:"packagePath"`
	PackageSignature string `json:"packageSignature"`
	PackageName      string `json:"packageName"`
	PackageAlgorithm string `json:"packageAlgorithm"`
}

type Function struct {
	Uri           string `json:"uri"`
	Url           string `json:"url"`
	Protocol      string `json:"protocol"`
	ContextPath   string `json:"contextPath"`
	Pool          []Pool `json:"pool"`
	Language      string `json:"language"`
	ClientIp      string `json:"clientIp"`
	Secure        bool   `json:"secure"`
	QueryString   string `json:"queryString"`
	ReplayRequest bool   `json:"replayRequest"`
	Method        string `json:"method"`
	ReqHeader     string `json:"reqHeader"`
	ReqBody       string `json:"reqBody"`
	ResBody       string `json:"resBody"`
	Scheme        string `json:"scheme"`
	ResHeader     string `json:"resHeader"`
}

type Pool struct {
	InvokeId         int           `json:"invokeId"`
	Interfaces       []interface{} `json:"interfaces"`
	TargetHash       []string      `json:"targetHash"`
	TargetValues     string        `json:"targetValues"`
	Signature        string        `json:"signature"`
	OriginClassName  string        `json:"originClassName"`
	SourceValues     string        `json:"sourceValues"`
	MethodName       string        `json:"methodName"`
	ClassName        string        `json:"className"`
	Source           bool          `json:"source"`
	CallerLineNumber int           `json:"callerLineNumber"`
	CallerClass      string        `json:"callerClass"`
	Args             string        `json:"args"`
	CallerMethod     string        `json:"callerMethod"`
	SourceHash       []string      `json:"sourceHash"`
	RetClassName     string        `json:"retClassName"`
}

type PoolReq struct {
	Source          bool
	OriginClassName string
	MethodName      string
	ClassName       string
	Args            []interface{}
	Reqs            []interface{}
}

type Log struct {
	Log string `json:"log"`
}

type Api struct {
	ApiData []ApiData `json:"apiData"`
}

type ApiData struct {
	Uri         string      `json:"uri"`
	Method      []string    `json:"method"`
	Class       string      `json:"class"`
	Parameters  []Parameter `json:"parameters"`
	ReturnType  string      `json:"returnType"`
	File        string      `json:"file"`
	Controller  string      `json:"controller"`
	Description string      `json:"description"`
}

type Parameter struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Annotation string `json:"annotation"`
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

func (p *PoolTree) FMT(pools *[]Pool, onlyKey int, goroutineIDs map[string]bool) {
	p.Pool.InvokeId = onlyKey
	*pools = append(*pools, *p.Pool)
	goroutineIDs[p.GoroutineID] = true
	fmt.Println(p.Pool.SourceValues, "===>", p.Pool.TargetValues)
	fmt.Println(p.Pool.SourceHash, "===>", p.Pool.TargetHash)
	for k, _ := range p.Children {
		onlyKey += 1
		p.Children[k].FMT(pools, onlyKey, goroutineIDs)
	}
}
