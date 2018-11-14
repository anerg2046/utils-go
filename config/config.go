package config

import (
	"utils-go/helper"

	"github.com/jinzhu/configor"
)

// Config 一些配置
var Config = struct {
	DB struct {
		Name string
		Host string
		User string
		Pass string
		Port string `default:"3306"`
	}
}{}

func init() {
	configor.Load(&Config, helper.GetCurrentDirectory()+"/config/config.yml")
}
