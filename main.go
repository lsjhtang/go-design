package main

import (
	. "wserver/classes"
	. "wserver/goft"
	. "wserver/middlewares"
	"wserver/remittance/a"
)

func main() {

	Ignite().
		//可以在控制器 同时使用两个ORM
		Beans(NewGormAdapter(), NewXOrmAdapter()). // 实现简单的依赖注入
		Attach(NewUserMiddle()).                   //带生命周期的中间件
		Mount("v1", NewUser()).
		Mount("v3", NewBook()).
		Mount("va", a.NewA()).
		/*Task("0/5 * * * * *", func() {
			log.Println("定时任务启动")
		}).*/
		Launch()

}
