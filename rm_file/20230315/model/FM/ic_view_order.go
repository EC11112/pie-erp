// 自动生成模板IcViewOrder
package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// IcViewOrder 结构体
type IcViewOrder struct {
      global.GVA_MODEL
      Item_id  *int `json:"item_id" form:"item_id" gorm:"column:item_id;comment:物品名称;"`
      Item_number  *int `json:"item_number" form:"item_number" gorm:"column:item_number;comment:物品数量;"`
}


// TableName IcViewOrder 表名
func (IcViewOrder) TableName() string {
  return "ic_view_order"
}

