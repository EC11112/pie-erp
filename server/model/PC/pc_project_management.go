// 自动生成模板PcProjectManagement
package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// PcProjectManagement 结构体
type PcProjectManagement struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:项目名称;size:64;"`
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:项目经理;"`
      Explanation  string `json:"explanation" form:"explanation" gorm:"column:explanation;comment:说明;size:400;"`
}


// TableName PcProjectManagement 表名
func (PcProjectManagement) TableName() string {
  return "pc_project_management"
}

