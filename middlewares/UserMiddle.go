package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserMiddle struct {

}

func NewUserMiddle() *UserMiddle {
	return &UserMiddle{}
}

func(this UserMiddle) OnRequest(context *gin.Context) error {
	log.Println("这是用户中间件")
	return nil
}
