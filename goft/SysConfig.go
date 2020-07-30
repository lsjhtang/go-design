package goft

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type UserConfig map[interface{}]interface{}
//递归读取用户配置文件
func GetConfigValue(m UserConfig,prefix []string,index int) interface{}  {
	key:=prefix[index]
	if v,ok:=m[key];ok{
		if index==len(prefix)-1{ //到了最后一个
			return v
		}else{
			index=index+1
			if mv,ok:=v.(UserConfig);ok{ //值必须是UserConfig类型
				return GetConfigValue(mv,prefix,index)
			}else{
				return  nil
			}

		}
	}
	return  nil
}

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
	Config UserConfig
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
