package htmlTemplateExecuteTemplate

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"html/template"
)

func init() {
	model.HookMap["htmlTemplateExecuteTemplate"] = new(HtmlTemplateExecuteTemplate)
}

type HtmlTemplateExecuteTemplate struct {
}

func (h *HtmlTemplateExecuteTemplate) Hook() {
	t := &template.Template{}
	err := gohook.HookMethod(t, "ExecuteTemplate", ExecuteTemplate, ExecuteTemplateT)
	if err != nil {
		fmt.Println(err, "HtmlTemplateExecuteTemplate")
	} else {
		fmt.Println("HtmlTemplateExecuteTemplate")
	}
}

func (h *HtmlTemplateExecuteTemplate) UnHook() {
	t := &template.Template{}
	gohook.UnHookMethod(t, "ExecuteTemplate")
}
