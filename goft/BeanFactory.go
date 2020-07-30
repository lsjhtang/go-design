package goft

import "reflect"

type BeanFactory struct {
	beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	b := &BeanFactory{beans: make([]interface{},0)}
	b.beans = append(b.beans, b)
	return b
}

func(this *BeanFactory) setBean(beans ...interface{})  {
	this.beans = append(this.beans, beans...)
}

func(this *BeanFactory) GetBean(bean interface{}) interface{} {
	return this.getBen(reflect.TypeOf(bean))
}

func(this *BeanFactory) getBen(t reflect.Type) interface{} {
	for _,p := range this.beans{
		if t == reflect.TypeOf(p)  {
			return p
		}
	}
	return nil
}

func(this *BeanFactory) Inject(object interface{}) {
	vObject := reflect.ValueOf(object)
	if vObject.Kind() == reflect.Ptr{//只处理指针类型
		vObject = vObject.Elem()
	}
	for i:=0;i<vObject.NumField();i++ {
		f:=vObject.Field(i)
		if f.Kind() != reflect.Ptr || !f.IsNil() {
			continue
		}
		if p := this.getBen(f.Type()); p != nil && f.CanInterface() {//类型一致则做处理
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

func(this *BeanFactory) inject(class IClass)  {
	vClass := reflect.ValueOf(class).Elem()
	for i:=0;i<vClass.NumField();i++ {
		f := vClass.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {//nil则无需重新初始化, 以及是否为指针类型
			continue
		}else {

			if IsAnnotation(f.Type()) {//处理注解
				f.Set(reflect.New(f.Type().Elem()))
				f.Interface().(Annotation).SetTag(reflect.TypeOf(class).Elem().Field(i).Tag)
				this.Inject(f.Interface())
				continue
			}

			if p := this.getBen(f.Type()); p != nil {//类型一致则做处理
				f.Set(reflect.New(f.Type().Elem()))
				f.Elem().Set(reflect.ValueOf(p).Elem())
			}

		}

	}
}