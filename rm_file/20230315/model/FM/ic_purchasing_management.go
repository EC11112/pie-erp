// 自动生成模板IcPurchasingManagement
package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// IcPurchasingManagement 结构体
type IcPurchasingManagement struct {
      global.GVA_MODEL
      Warehouse_id  *int `json:"warehouse_id" form:"warehouse_id" gorm:"column:warehouse_id;comment:仓库名称;"`
      Item_id  *int `json:"item_id" form:"item_id" gorm:"column:item_id;comment:物品名称;"`
      Item_number  *int `json:"item_number" form:"item_number" gorm:"column:item_number;comment:物品数量;"`
}


// TableName IcPurchasingManagement 表名
func (IcPurchasingManagement) TableName() string {
  return "ic_purchasing_management"
}

