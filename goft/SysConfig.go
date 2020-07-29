package goft

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ServerConfig struct {
	Port int
	Name string
}

//系统配置
type SysConfig struct {
	Server *ServerConfig
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port:8081, Name:"web"}}
}

func InitConfig() *SysConfig {
	config := NewSysConfig()
	if b := LoadConfigFile() ; b !=nil{
		if err:=yaml.Unmarshal(b, config) ; err != nil {
			panic("请检查配置文件")
		}
	}
	return config
}

func LoadConfigFile() []byte  {
	dir,_ := os.Getwd()
	file:=dir+"/application.yaml"
	b,err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return b
}
