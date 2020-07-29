package classes

import (
	"github.com/gin-gonic/gin"
	"wserver/goft"
	"wserver/models"
)

type User struct {
	*goft.GormAdapter
}

func NewUser() *User {
	return &User{}
}

func(this *User) GetUser(context *gin.Context) string {
	/*return func(context *gin.Context) {
		context.JSON(200, gin.H{"data": map[string]string{"msg":"获取用户信息成功"}})
	}*/
	return "123"
}

func(this *User) UserDetail(context *gin.Context) goft.Model {
	ctx := models.NewUserModel()
	err := context.BindUri(ctx)
	goft.Error(err)
	return ctx
}

func(this *User) UserList(context *gin.Context) goft.Model {
	//users := []*models.UserModel{{ID: 202, UserName: "李四"},{ID: 303, UserName: "王五"}}
	//return goft.MakeModels(users)
	userModel := models.NewUserModel()
	err := context.BindUri(userModel)
	goft.Error(err)
	this.Table("users").Find(userModel)
	return userModel

}


func(this *User) Build(goft *goft.Goft)  {
	goft.Handle("GET", "/user", this.GetUser)
	goft.Handle("GET", "/user_detail/:id", this.UserDetail)
	goft.Handle("GET", "/user_list/:id", this.UserList)
}