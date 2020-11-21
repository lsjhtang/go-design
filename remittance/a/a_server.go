package a

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"wserver/goft"
)

type A struct {
	*goft.GormAdapter
}

func NewA() *A {
	return &A{}
}

func (this *A) TransMoney(context *gin.Context) goft.Model {
	rm := newRemittanceModel()
	err := context.BindJSON(rm)
	goft.Error(err)
	tx := GetDb().BeginTx(context, nil)
	row := tx.Table("usermoney").Where("username = ? and usermoney >= ?", rm.From, rm.Money).UpdateColumn("usermoney", gorm.Expr("usermoney - ?", rm.Money)).RowsAffected
	if row == 0 {
		tx.Rollback()
		goft.Error(fmt.Errorf("扣款失败"))
	}

	err = tx.Table("money_log").Create(rm).Error
	if err != nil {
		tx.Rollback()
		goft.Error(err, "插入日志失败")
	}

	tx.Commit()

	return rm
}

func (this *A) Build(goft *goft.Goft) {
	goft.Handle("POST", "/atob", this.TransMoney)
}
