package global

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
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
	configErr := v.Unmarshal(&Config)
	if err != nil {
		fmt.Println(configErr)
		return
	}
}
