package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"wserver/goft"
	"wserver/models"
)

type User struct {
	*goft.GormAdapter
	Age *goft.Value `prefix:"user.age"`
}

func NewUser() *User {
	return &User{}
}

func (this *User) GetUser() string {
	/*return func(context *gin.Context) {
		context.JSON(200, gin.H{"data": map[string]string{"msg":"获取用户信息成功"}})
	}*/
	return this.Age.String()
}

func (this *User) UserDetail(context *gin.Context) goft.Model {
	ctx := models.NewUserModel()
	goft.Error(context.BindUri(ctx))
	return ctx
}

func (this *User) UpdateViews(params ...interface{}) {
	this.Table("users").Where("id=?", params[0]).Update("views", gorm.Expr("views+1"))
}

func (this *User) UserList(context *gin.Context) goft.Model {
	//users := []*models.UserModel{{ID: 202, UserName: "李四"},{ID: 303, UserName: "王五"}}
	//return goft.MakeModels(users)
	userModel := models.NewUserModel()
	err := context.BindUri(userModel)
	goft.Error(err)
	this.Table("users").Find(userModel)
	goft.Task(this.UpdateViews, func() {
		this.after(userModel.ID)
	}, userModel.ID) //任务入队
	return userModel

}

func (this *User) after(p ...interface{}) {
	log.Print("after callback", p[0].(int))
}

func (this *User) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.GetUser)
	goft.Handle("GET", "/user_detail/:id", this.UserDetail)
	goft.Handle("GET", "/user_list/:id", this.UserList)
}
