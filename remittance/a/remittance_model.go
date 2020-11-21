package a

import "fmt"

type remittanceModel struct {
	From  string `json:"from" gorm:"from;type:varchar(255)"`
	To    string `json:"to" gorm:"to;type:varchar(255)"`
	Money int    `json:"money" gorm:"money;type:int"`
}

func newRemittanceModel() *remittanceModel {
	return &remittanceModel{}
}

func (this *remittanceModel) String() string {
	return fmt.Sprintf("%s转账给%s, 金额是%d\n", this.From, this.To, this.Money)
}
