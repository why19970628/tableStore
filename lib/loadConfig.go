package lib

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Conf 配置文件全局变量
var Conf Config

// Config 配置文件结构
type Config struct {
	Port            string
	Endpoint        string
	InstanceName    string
	AccessKeyID     string
	AccessKeySecret string
}

// LoadConfig 加载
func LoadConfig() {
	file, _ := ioutil.ReadFile("./config.yml")
	if len(file) == 0 {
		panic("未发现config.yml文件")
	}
	config := Config{}
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	Conf = config
}
