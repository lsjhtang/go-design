package goft

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormAdapter struct {
	*gorm.DB
}

func NewGormAdapter() *GormAdapter {
	db, err := gorm.Open("mysql", "root:root@(172.20.10.13:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//db.SingularTable(true)//表是否默认添加s
	db.DB().SetMaxIdleConns(10)//最大空闲连接
	db.DB().SetMaxOpenConns(100)//最大连接

	db.LogMode(true)//debug模式
	return &GormAdapter{DB:db}
}

