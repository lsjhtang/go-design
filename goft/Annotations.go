package goft

import (
	"reflect"
)

var Annotations []Annotation

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

func init()  {
	Annotations = make([]Annotation, 0)
	Annotations = append(Annotations, new(Value))
}

type Value struct {
	tag reflect.StructTag
}

func(this *Value) SetTag(tag reflect.StructTag)  {
	this.tag = tag
}

func(this *Value) String() string {
	return "18"
}