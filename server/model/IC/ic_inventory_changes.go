// 自动生成模板IcInventoryChanges
package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// IcInventoryChanges 结构体
type IcInventoryChanges struct {
      global.GVA_MODEL
      Warehouse_id  *int `json:"warehouse_id" form:"warehouse_id" gorm:"column:warehouse_id;comment:仓库名称;"`
      Change_type  string `json:"change_type" form:"change_type" gorm:"column:change_type;type:enum('出库','入库');comment:变更类型;"`
      Item_id  *int `json:"item_id" form:"item_id" gorm:"column:item_id;comment:物品名称;"`
      Item_number  *int `json:"item_number" form:"item_number" gorm:"column:item_number;comment:物品数量;"`
}


// TableName IcInventoryChanges 表名
func (IcInventoryChanges) TableName() string {
  return "ic_inventory_changes"
}

