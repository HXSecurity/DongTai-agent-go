package global

import "go-agent/model/request"

var (
	AgentId   = 0
	HookGroup = make(map[string]request.UploadReq)
)
