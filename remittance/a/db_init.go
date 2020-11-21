package a

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type GormAdapter struct {
	*gorm.DB
}

func GetDb() *GormAdapter {
	db, err := gorm.Open("mysql",
		"root:root@tcp(192.168.87.128:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true) //调试模式
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	return &GormAdapter{DB: db}
}
