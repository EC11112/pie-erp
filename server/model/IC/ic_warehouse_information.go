// 自动生成模板IcWarehouseInformation
package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// IcWarehouseInformation 结构体
type IcWarehouseInformation struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:仓库名称;size:64;"`
}


// TableName IcWarehouseInformation 表名
func (IcWarehouseInformation) TableName() string {
  return "ic_warehouse_information"
}

