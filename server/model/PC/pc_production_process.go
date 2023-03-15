// 自动生成模板PcProductionProcess
package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// PcProductionProcess 结构体
type PcProductionProcess struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:流程名称;size:64;"`
      Project_id  *int `json:"project_id" form:"project_id" gorm:"column:project_id;comment:所属项目;"`
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:车间主任;"`
      Production_cycle  *time.Time `json:"production_cycle" form:"production_cycle" gorm:"column:production_cycle;comment:生产周期;"`
      Explanation  string `json:"explanation" form:"explanation" gorm:"column:explanation;comment:说明;size:400;"`
}


// TableName PcProductionProcess 表名
func (PcProductionProcess) TableName() string {
  return "pc_production_process"
}

