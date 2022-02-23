package global

import (
	"sync"
)

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
	AgentId           = 0
	PoolTreeMap       = sync.Map{}
	ResponseMap       = sync.Map{}
	ResponseHeaderMap = sync.Map{}
)
