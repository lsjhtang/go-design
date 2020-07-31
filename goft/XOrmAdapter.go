package goft

import (
	"log"
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type XOrmAdapter struct {
	*xorm.Engine
}
func NewXOrmAdapter() *XOrmAdapter {
	engine, err := xorm.NewEngine("mysql",
		"root:root@tcp(localhost:3306	)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		log.Fatal(err)
	}
	engine.DB().SetMaxIdleConns(5)
	engine.DB().SetMaxOpenConns(10)
	return &XOrmAdapter{Engine:engine}
}