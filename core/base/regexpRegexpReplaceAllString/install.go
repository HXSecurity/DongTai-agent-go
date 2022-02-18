package regexpRegexpReplaceAllString

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"regexp"
)

func init() {
	model.HookMap["regexpRegexpReplaceAllString"] = new(RegexpRegexpReplaceAllString)
}

type RegexpRegexpReplaceAllString struct {
}

func (h *RegexpRegexpReplaceAllString) Hook() {
	var r *regexp.Regexp
	err := gohook.HookMethod(r, "ReplaceAllString", ReplaceAllString, ReplaceAllStringT)
	if err != nil {
		fmt.Println(err, "RegexpRegexpReplaceAllString")
	} else {
		fmt.Println("RegexpRegexpReplaceAllString")
	}
}

func (h *RegexpRegexpReplaceAllString) UnHook() {
	var r *regexp.Regexp
	err := gohook.UnHookMethod(r, "ReplaceAllString")
	if err != nil {
		fmt.Println(err, "RegexpRegexpReplaceAllString")
	} else {
		fmt.Println("RegexpRegexpReplaceAllString")
	}
}
