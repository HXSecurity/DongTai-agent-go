package hook

type Base struct {
}

func (b *Base) GetHook() []string {
	return []string{
		"sqlDBQuery",
		"fmtSprintf",
		"jsonUnmarshal",
		"jsonDecoderDecode",
		"jsonNewDecoder",
		"runtimeConcatstrings",
		"execCommand",
		"execCmdStart",
		"bufioWriterWrite",
		"bufioWriterWriteString",
		"runtimesSringtoslicebyte",
		"htmlTemplateExecuteTemplate",
		"httpClientDo",
		"httpNewRequest",
		"stringsBuilderWriteString",
		"osOpenFile",
		"ioReadAll",
		"regexpRegexpReplaceAllString",
		"urlUrlString",
	}
}

func (b *Base) HookAll() {
	Hook(b.GetHook())
}

func (b *Base) UnHookAll() {
	UnHook(b.GetHook())
}
