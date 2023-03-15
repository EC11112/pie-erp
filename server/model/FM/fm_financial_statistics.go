// 自动生成模板FmFinancialStatistics
package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FmFinancialStatistics 结构体
type FmFinancialStatistics struct {
      global.GVA_MODEL
      Mean  string `json:"mean" form:"mean" gorm:"column:mean;comment:规则名;size:64;"`
      Price  *int `json:"price" form:"price" gorm:"column:price;comment:规则值;"`
}


// TableName FmFinancialStatistics 表名
func (FmFinancialStatistics) TableName() string {
  return "fm_financial_statistics"
}

