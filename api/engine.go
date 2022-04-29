package api

import (
	"encoding/json"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/model/response"
	"github.com/pkg/errors"
)

/*	AgentRegister
 	Agent 启动时，向 Server 端注册的接口
	req: api.AgentRegisterReq
*/
func AgentRegister(req request.AgentRegisterReq) (AgentId int, err error) {
	resp, body, errs := POST("/api/v1/agent/register", req).End()
	if len(errs) > 0 {
		for _, v := range errs {
			fmt.Println(v)
		}
		err = errs[0]
		return 0, err
	}
	if resp.StatusCode == 200 {
		var res response.AgentRegisterRes
		err := json.Unmarshal([]byte(body), &res)
		if err != nil {
			return 0, err
		}
		if res.Status == 201 {
			AgentId = res.Data.Id
			fmt.Println("注册成功，探针ID为", res.Data.Id)
			return AgentId, nil
		} 
		return 0, errors.New("注册失败，失败原因" + res.Msg)

	}
	return 0, errors.New("状态码为" + resp.Status)
}

func Limit() map[string]string {
	var limit = make(map[string]string)
	resp, body, errs := GET("/api/v1/agent/limit", nil).End()
	if len(errs) > 0 {
		for _, v := range errs {
			fmt.Println(v)
		}
		return limit
	}
	if resp.StatusCode == 200 {
		var res response.LimitRes
		err := json.Unmarshal([]byte(body), &res)
		if err != nil {
			fmt.Println(err)
		}
		for _, item := range res.Data {
			limit[item.Key] = item.Value
		}
	}
	return limit
}

/*	Profiles
	Engine 运行时，从 OpenAPI 服务获取Hook规则
*/
func Profiles(req request.HookRuleReq) {
	resp, body, errs := GET("/api/v1/profiles", req).End()
	if len(errs) > 0 {
		for _, v := range errs {
			fmt.Println(v)
		}
		return
	}
	if resp.StatusCode == 200 {
		var res response.AgentRegisterRes
		err := json.Unmarshal([]byte(body), &res)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ReportUpload(req request.UploadReq) {
	resp, body, errs := POST("/api/v1/report/upload", req).End()
	if len(errs) > 0 {
		for _, v := range errs {
			fmt.Println(v)
		}
		return
	}
	if resp.StatusCode == 200 {
		var res response.ResBase
		err := json.Unmarshal([]byte(body), &res)
		if res.Status == 201 {
			//fmt.Println(res.Msg)
		} else {
			//fmt.Println(res.Msg)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
