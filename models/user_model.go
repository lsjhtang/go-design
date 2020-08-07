package models

import "fmt"

type UserModel struct {
	ID       int    `uri:"id" banding:"required,gt=0" gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string `json:"UserName" gorm:"column:user_name;type:varchar(255)"`
	Age      int    ` json:"Age" gorm:"column:age;type:int"`
	Views    int    `gorm:"column:views;type:int"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this UserModel) String() string {
	return fmt.Sprint("userid:%d, username:%s", this.ID, this.UserName)
}
