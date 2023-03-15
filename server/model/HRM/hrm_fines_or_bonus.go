// 自动生成模板HrmFinesOrBonus
package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// HrmFinesOrBonus 结构体
type HrmFinesOrBonus struct {
      global.GVA_MODEL
      Action  string `json:"action" form:"action" gorm:"column:action;type:enum('罚款','奖金');comment:动作;"`
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:相关人员;"`
      Amount  *int `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`
      Explanation  string `json:"explanation" form:"explanation" gorm:"column:explanation;comment:说明;size:400;"`
}


// TableName HrmFinesOrBonus 表名
func (HrmFinesOrBonus) TableName() string {
  return "hrm_fines_or_bonus"
}

