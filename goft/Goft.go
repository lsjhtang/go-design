package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
	props []interface{}
}

func Ignite() *Goft {
	g:= &Goft{Engine:gin.New(), props:make([]interface{},0)}
	g.Use(ErrorHandler())
	return g
}

func (this *Goft) Launch()  {
	config := InitConfig()
	this.Run(fmt.Sprintf(":%d",config.Server.Port))
}

func (this *Goft) DB(db interface{}) *Goft {
	this.props = append(this.props, db)
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
		this.setProp(class)
	}
	return this
}

func(this *Goft) getProp(f reflect.Type) interface{}  {
	for _,p:=range this.props{
		if reflect.TypeOf(p) == f {
			return p
		}
	}
	return nil
}

func(this *Goft) setProp(class IClass)  {
	vClass := reflect.ValueOf(class).Elem()
	for i:=0;i<vClass.NumField();i++ {
		f := vClass.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {//nil则无需重新初始化, 以及是否为指针类型
			continue
		}else {
			if p := this.getProp(f.Type()); p != nil {//类型一致则做处理
				f.Set(reflect.New(f.Type().Elem()))
				f.Elem().Set(reflect.ValueOf(p).Elem())
			}
		}

	}
}
