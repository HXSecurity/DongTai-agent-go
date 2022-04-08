package request

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"reflect"
	"strconv"
)

type AgentRegisterReq struct {
	AutoCreateProject bool   `json:"autoCreateProject"`
	Name              string `json:"name"`
	Language          string `json:"language"`
	Version           string `json:"version"`
	ProjectName       string `json:"projectName"`
	Hostname          string `json:"hostname"`
	Network           string `json:"network"`
	ContainerName     string `json:"containerName"`
	ContainerVersion  string `json:"containerVersion"`
	ServerAddr        string `json:"serverAddr"`
	ServerPort        string `json:"serverPort"`
	ServerPath        string `json:"serverPath"`
	ServerEnv         string `json:"serverEnv"`
	Pid               string `json:"pid"`
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
	InvokeId int    `json:"invokeId"`
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
	TraceId       string `json:"traceId"`
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
	TraceId          string        `json:"traceId"`
}

type PoolReq struct {
	Source          bool
	OriginClassName string
	MethodName      string
	ClassName       string
	Args            []interface{}
	Reqs            []interface{}
	NeedHook        []interface{}
	NeedCatch       []interface{}
	ArgsStr         string
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

func (p *PoolTree) FMT(pools *[]Pool, w *utils.Worker, goroutineIDs map[string]bool, TraceId string) {
	p.Pool.InvokeId = int(w.GetId())
	p.Pool.TraceId = TraceId
	*pools = append(*pools, *p.Pool)
	goroutineIDs[p.GoroutineID] = true
	fmt.Println(p.Pool.ClassName, p.Pool.MethodName)
	fmt.Println(p.Pool.SourceValues, "====>", p.Pool.TargetValues)
	fmt.Println(p.Pool.SourceHash, "===>", p.Pool.TargetHash)
	for k, v := range p.Children {
		global.PoolTreeMap.Delete(&v.Pool.TargetHash)
		p.Children[k].FMT(pools, w, goroutineIDs, TraceId)
	}
}

func Collect(args ...interface{}) []interface{} {
	return args
}

func FmtHookPool(p PoolReq) Pool {
	signature, callerClass, callerMethod, callerLineNumber := utils.FmtStack()
	var sourceHash global.HashKeys = make(global.HashKeys, 0)
	var SourceValues string = ""
	if len(p.NeedHook) == 0 {
		utils.RangeSource(p.Args, &p.NeedHook)
	}
	for _, v := range p.NeedHook {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, utils.GetSource(v))
			SourceValues = utils.StringAdd(SourceValues, v.(string), " ")
			break
		case uintptr:
			sourceHash = append(sourceHash, strconv.Itoa(int(v.(uintptr))))
			SourceValues = utils.StringAdd(SourceValues, strconv.Itoa(int(v.(uintptr))), " ")
			break
		default:
			t := reflect.TypeOf(v)
			value := reflect.ValueOf(v)
			if t.Kind() == reflect.String {
				sourceHash = append(sourceHash, utils.GetSource(value.String()))
				SourceValues = utils.StringAdd(SourceValues, value.String(), " ")
			}
			break
		}
	}
	var targetHash global.HashKeys = make(global.HashKeys, 0)
	var targetValues string = ""
	var RetClassNames string = ""
	if len(p.NeedCatch) == 0 {
		utils.RangeSource(p.Reqs, &p.NeedCatch)
	}
	for _, v := range p.Reqs {
		if reflect.ValueOf(v).IsValid() {
			RetClassNames = utils.StringAdd(RetClassNames, reflect.ValueOf(v).Type().String(), " ")
		}
	}
	for _, v := range p.NeedCatch {
		switch v.(type) {
		case string:
			targetHash = append(targetHash, utils.GetSource(v))
			targetValues = utils.StringAdd(targetValues, v.(string), " ")
		case uintptr:
			targetHash = append(targetHash, strconv.Itoa(int(v.(uintptr))))
			targetValues = utils.StringAdd(targetValues, strconv.Itoa(int(v.(uintptr))), " ")
		}
	}
	var ArgsStr string
	if p.ArgsStr != "" {
		ArgsStr = p.ArgsStr
	} else {
		ArgsStr = utils.Strval(p.Args)
	}
	var pool = Pool{
		Source:           p.Source,
		Interfaces:       []interface{}{},
		SourceValues:     SourceValues,
		SourceHash:       sourceHash,
		TargetValues:     targetValues,
		TargetHash:       targetHash,
		Signature:        signature,
		OriginClassName:  p.OriginClassName,
		MethodName:       p.MethodName,
		ClassName:        p.ClassName,
		CallerLineNumber: callerLineNumber,
		CallerClass:      callerClass,
		CallerMethod:     callerMethod,
		RetClassName:     RetClassNames,
		Args:             ArgsStr,
	}

	poolTree := PoolTree{
		Begin:       p.Source,
		Pool:        &pool,
		Children:    []*PoolTree{},
		GoroutineID: utils.CatGoroutineID(),
	}

	if !p.Source {
		global.PoolTreeMap.Range(func(key, value interface{}) bool {
			if key.(*global.HashKeys).Some(sourceHash) {
				poolTree.GoroutineID = value.(*PoolTree).GoroutineID
				value.(*PoolTree).Children = append(value.(*PoolTree).Children, &poolTree)
				global.PoolTreeMap.Store(&targetHash, &poolTree)
			}
			return true
		})
	} else {
		global.PoolTreeMap.Store(&targetHash, &poolTree)
	}
	return pool
}

func RunMapGCbYGoroutineID(goroutineIDs map[string]bool) {
	global.PoolTreeMap.Range(func(key, value interface{}) bool {
		for id, _ := range goroutineIDs {
			if value.(*PoolTree).GoroutineID == id {
				global.PoolTreeMap.Delete(key)
			}
		}
		return true
	})
}
