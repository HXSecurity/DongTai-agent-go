package global

import (
	"fmt"
	"github.com/spf13/viper"
	"go-agent/model"
)

var Config model.Config

func InitViper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
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
