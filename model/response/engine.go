package response

type ResBase struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type AgentRegisterRes struct {
	ResBase
	Data struct {
		Id            int `json:"id"`
		CoreAutoStart int `json:"coreAutoStart"`
	} `json:"data"`
}
