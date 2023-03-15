// 自动生成模板IcViewInventory
package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// IcViewInventory 结构体
type IcViewInventory struct {
      global.GVA_MODEL
      Item_id  *int `json:"item_id" form:"item_id" gorm:"column:item_id;comment:物品名称;"`
      Item_number  *int `json:"item_number" form:"item_number" gorm:"column:item_number;comment:物品数量;"`
      Warehouse_id  *int `json:"warehouse_id" form:"warehouse_id" gorm:"column:warehouse_id;comment:仓库名称;"`
}


// TableName IcViewInventory 表名
func (IcViewInventory) TableName() string {
  return "ic_View_inventory"
}

