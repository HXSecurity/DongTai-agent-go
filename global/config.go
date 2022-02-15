package global

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Config model.Config

func InitViper() {
	v := viper.New()
	v.SetConfigFile("dongtai-go-agent-config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	pflag.String("DongtaiGoProjectVersion", "v1.0.0", "Project Version")
	pflag.String("DongtaiGoProjectName", "GO Project", "Project Name")
	pflag.Bool("DongtaiGoProjectCreate", true, "Auto Create Project")
	pflag.Parse()
	FErr := v.BindPFlags(pflag.CommandLine)
	if FErr != nil {
		fmt.Println("[waring]You may have used parameters other than Dongtai-GO-Agent. If it is a parameter carried by your system, please ignore this prompt.")
	}
	configErr := v.Unmarshal(&Config)
	if err != nil {
		fmt.Println(configErr)
		return
	}
}
