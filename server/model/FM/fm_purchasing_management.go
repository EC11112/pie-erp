// 自动生成模板FmPurchasingManagement
package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FmPurchasingManagement 结构体
type FmPurchasingManagement struct {
      global.GVA_MODEL
      Warehouse_id  *int `json:"warehouse_id" form:"warehouse_id" gorm:"column:warehouse_id;comment:仓库名称;"`
      Item_id  *int `json:"item_id" form:"item_id" gorm:"column:item_id;comment:物品名称;"`
      Item_number  *int `json:"item_number" form:"item_number" gorm:"column:item_number;comment:物品数量;"`
}


// TableName FmPurchasingManagement 表名
func (FmPurchasingManagement) TableName() string {
  return "fm_purchasing_management"
}

