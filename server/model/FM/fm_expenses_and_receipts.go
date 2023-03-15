// 自动生成模板FmExpensesAndReceipts
package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FmExpensesAndReceipts 结构体
type FmExpensesAndReceipts struct {
      global.GVA_MODEL
      Expenses_or_receipts  string `json:"expenses_or_receipts" form:"expenses_or_receipts" gorm:"column:expenses_or_receipts;type:enum('收入','支出');comment:收或支;"`
      Amount  *int `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`
}


// TableName FmExpensesAndReceipts 表名
func (FmExpensesAndReceipts) TableName() string {
  return "fm_expenses_and_receipts"
}

