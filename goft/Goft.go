package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
	beanFactory *BeanFactory
}

func Ignite() *Goft {
	g:= &Goft{Engine:gin.New(), beanFactory: NewBeanFactory()}
	g.Use(ErrorHandler())
	g.beanFactory.setBean(InitConfig())  //整个配置加载进bean中
	return g
}

func (this *Goft) Launch()  {
	var port int=8080
	if config:=this.beanFactory.GetBean(new(SysConfig));config!=nil {
		port = config.(*SysConfig).Server.Port
	}
	this.Run(fmt.Sprintf(":%d",port))
}

//设置数据库连接
func (this *Goft) Beans(beans ...interface{}) *Goft {
	this.beanFactory.setBean(beans)
	return this
}

func(this *Goft) Attach(f Fairing) *Goft   {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			panic(err)
		}else {
			context.Next()
		}
	})
	return this
}

func(this *Goft) Handle(httpMethod, relativePath string, handlers interface{}) *Goft  {
	if h := Convert(handlers); h != nil{
		this.g.Handle(httpMethod, relativePath, h)
	}else {
		this.g.Handle(httpMethod, relativePath,  handlers.(gin.HandlerFunc))
	}
	return this
}

func (this *Goft) Mount(gorup string, classes ...IClass) *Goft {
	this.g = this.Group(gorup)
	for _,class := range classes{
		class.Build(this)
		this.beanFactory.inject(class)
	}
	return this
}

