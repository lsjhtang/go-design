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

type Drive struct {
	Connection string
	Port int
	Host string
	Database string
	UserName string
	Password string
}

//系统配置
type SysConfig struct {
	Server *ServerConfig
	Drive *Drive
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port:8081, Name:"web"},
					 Drive: &Drive{Connection: "mysql", Port: 3306, Host: "127.0.0.1", Database: "test", UserName: "root", Password: "root"},
		}
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
