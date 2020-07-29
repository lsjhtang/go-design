package goft

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormAdapter struct {
	*gorm.DB
}

func NewGormAdapter() *GormAdapter {
	config := InitConfig().Drive
	db, err := gorm.Open(fmt.Sprintf("%s",config.Connection),
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.UserName, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		panic(err)
	}
	//db.SingularTable(true)//表是否默认添加s
	db.DB().SetMaxIdleConns(10)//最大空闲连接
	db.DB().SetMaxOpenConns(100)//最大连接

	db.LogMode(true)//debug模式
	return &GormAdapter{DB:db}
}

