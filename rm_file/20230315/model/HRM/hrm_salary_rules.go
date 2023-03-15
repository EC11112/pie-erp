// 自动生成模板HrmSalaryRules
package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// HrmSalaryRules 结构体
type HrmSalaryRules struct {
      global.GVA_MODEL
      Mean  string `json:"mean" form:"mean" gorm:"column:mean;comment:数据含义;size:10;"`
      Price  *int `json:"price" form:"price" gorm:"column:price;comment:和工资有关的数据;"`
}


// TableName HrmSalaryRules 表名
func (HrmSalaryRules) TableName() string {
  return "hrm_salary_rules"
}

