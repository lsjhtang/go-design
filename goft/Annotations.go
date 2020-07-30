package goft

import (
	"fmt"
	"reflect"
	"strings"
)

var Annotations []Annotation

func init()  {
	Annotations = make([]Annotation, 0)
	Annotations = append(Annotations, new(Value))
}

type Annotation interface {
	SetTag(tag reflect.StructTag)
	String() string
}

func IsAnnotation(t reflect.Type) bool {
	for _,v := range Annotations {
		if reflect.TypeOf(v) == t {
			return true
		}
	}
	return false
}


type Value struct {
	tag reflect.StructTag
	BeanFactory *BeanFactory
}

func(this *Value) SetTag(tag reflect.StructTag)  {
	this.tag = tag
}

func(this *Value) String() string {
	get_prefix:=this.tag.Get("prefix")
	if get_prefix==""{
		return ""
	}
	prefix:=strings.Split(get_prefix,".")
	if config:=this.BeanFactory.GetBean(new(SysConfig));config!=nil{
		get_value:=GetConfigValue(config.(*SysConfig).Config,prefix,0)
		if get_value!=nil{
			return fmt.Sprintf("%v",get_value)
		}else{
			return ""
		}
	}else{
		return ""
	}
}