package global

import (
	"flag"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Config model.Config

func InitViper() {
	config, err := ioutil.ReadFile("dongtai-go-agent-config.yaml")
	fmt.Println(string(config))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(config, &Config)
	if err != nil {
		fmt.Println(err)
		return
	}

	var version string
	var name string
	var auto bool

	flag.StringVar(&version, "DongtaiGoProjectVersion", "v1.0.0", "Project Version")
	flag.StringVar(&name, "DongtaiGoProjectName", "GO Project", "Project Name")
	flag.BoolVar(&auto, "DongtaiGoProjectCreate", true, "Auto Create Project")
	flag.Parse()
	if version != "v1.0.0" {
		Config.DongtaiGoProjectVersion = version
	}
	if name != "GO Project" {
		Config.DongtaiGoProjectName = name
	}
	if auto != true {
		Config.DongtaiGoProjectCreate = auto
	}
	fmt.Println(Config)
}
