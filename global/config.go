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
	var projectGroupId int
	var projectTemplateId int

	flag.StringVar(&version, "DongtaiGoProjectVersion", "v1.0.0", "Project Version")
	flag.StringVar(&name, "DongtaiGoProjectName", "GO Project", "Project Name")
	flag.BoolVar(&auto, "DongtaiGoProjectCreate", true, "Auto Create Project")
	flag.IntVar(&projectGroupId, "DongtaiGoProjectGroupId", 1, "Group ID")
	flag.IntVar(&projectTemplateId, "DongtaiGoProjectTemplateId", 1, "Template ID")

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
	if projectGroupId != 1 {
		Config.DongtaiGoProjectGroupId = projectGroupId
	}
	if projectTemplateId != 1 {
		Config.DongtaiGoProjectTemplateId = projectTemplateId
	}

	fmt.Println(Config)
}
