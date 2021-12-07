package global

import "go-agent/model/request"

type HashKeys []string

func (h *HashKeys) Some(source []string) bool {
	for i := 0; i < len(*h); i++ {
		for k := 0; k < len(source); k++ {
			if (*h)[i] == source[k] {
				return true
			}
		}
	}
	return false
}

var (
	AgentId     = 0
	HookGroup   = make(map[string]*request.UploadReq)
	PoolTreeMap = make(map[*HashKeys]*request.PoolTree)
)
