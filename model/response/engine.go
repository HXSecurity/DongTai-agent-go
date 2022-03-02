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

type LimitRes struct {
	ResBase
	Data []struct {
		Id    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"data"`
}
