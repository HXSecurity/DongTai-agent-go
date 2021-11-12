package response

type ResBase struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type AgentRegisterRes struct {
	ResBase
	Data struct {
		Id            string `json:"id"`
		CoreAutoStart int    `json:"coreAutoStart"`
	} `json:"data"`
}
