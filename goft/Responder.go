package goft

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var RsponderList []Responder

func init()  {
	RsponderList = []Responder{	new(StringResponder),
							   	new(ModelResponder),
								new(ModelsResponder),
	}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc  {
	Href := reflect.ValueOf(handler)
	for _,r := range RsponderList {
		Rref := reflect.ValueOf(r).Elem()
		if Href.Type().ConvertibleTo(Rref.Type()) {
			Rref.Set(Href)
			return Rref.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, this(context))
	}
}

type ModelResponder func(*gin.Context) Model

func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, this(context))
	}
}

type ModelsResponder func(*gin.Context) Models

func (this ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-type", "application/json")
		_,err :=context.Writer.WriteString(string(this(context)))
		if err != nil{
			Error(err)
		}
	}
}