package model

type Config struct {
	DongtaiGoOpenapi        string `yaml:"DongtaiGoOpenapi"`
	DongtaiGoToken          string `yaml:"DongtaiGoToken"`
	DongtaiGoProjectName    string `yaml:"DongtaiGoProjectName"`
	DongtaiGoProjectVersion string `yaml:"DongtaiGoProjectVersion"`
	DongtaiGoProjectCreate  bool   `yaml:"DongtaiGoProjectCreate"`
	DongtaiGoAgentToken     string `yaml:"DongtaiGoAgentToken"`
}
