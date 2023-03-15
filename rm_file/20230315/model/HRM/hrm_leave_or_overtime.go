// 自动生成模板HrmLeaveOrOvertime
package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HrmLeaveOrOvertime 结构体
type HrmLeaveOrOvertime struct {
      global.GVA_MODEL
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:事件申请者ID;"`
      Is_leave  *bool `json:"is_leave" form:"is_leave" gorm:"column:is_leave;comment:请假还是加班;"`
      Start_date  *time.Time `json:"start_date" form:"start_date" gorm:"column:start_date;comment:起始日期;"`
      End_date  *time.Time `json:"end_date" form:"end_date" gorm:"column:end_date;comment:终止日期;"`
      Is_approved  *bool `json:"is_approved" form:"is_approved" gorm:"column:is_approved;comment:是否批准;"`
      Explanation  string `json:"explanation" form:"explanation" gorm:"column:explanation;comment:说明;size:400;"`
}


// TableName HrmLeaveOrOvertime 表名
func (HrmLeaveOrOvertime) TableName() string {
  return "hrm_leave_or_overtime"
}

