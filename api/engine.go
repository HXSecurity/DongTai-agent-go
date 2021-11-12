package api

import (
	"encoding/json"
	"fmt"
	"go-agent/model/request"
	"go-agent/model/response"
)

/*	AgentRegister
 	Agent 启动时，向 Server 端注册的接口
	req: api.AgentRegisterReq
*/
func AgentRegister(req request.AgentRegisterReq) {
	resp, body, errs := POST("/api/v1/agent/register", req).End()
	if len(errs) > 0 {
		for _, v := range errs {
			fmt.Println(v)
		}
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println(body)
		var res response.AgentRegisterRes
		err := json.Unmarshal([]byte(body), &res)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
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
		fmt.Println(body)
		var res response.AgentRegisterRes
		err := json.Unmarshal([]byte(body), &res)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
