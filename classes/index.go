package classes

import (
	"github.com/gin-gonic/gin"
	"wserver/goft"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (this *Index) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"data": map[string]string{"msg": "获取主页成功"}})
	}
}

func (this *Index) Build(goft *goft.Goft) {
	goft.Handle("GET", "/index", this.GetIndex())
}
