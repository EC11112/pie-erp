// 自动生成模板PcInputAndOutput
package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// PcInputAndOutput 结构体
type PcInputAndOutput struct {
      global.GVA_MODEL
      Process_id  *int `json:"process_id" form:"process_id" gorm:"column:process_id;comment:生产流程;"`
      Input_or_output  string `json:"input_or_output" form:"input_or_output" gorm:"column:input_or_output;type:enum('消耗','产出');comment:消耗产出;"`
      Item_id  *int `json:"item_id" form:"item_id" gorm:"column:item_id;comment:物品名称;"`
      Item_number  *int `json:"item_number" form:"item_number" gorm:"column:item_number;comment:物品数量;"`
}


// TableName PcInputAndOutput 表名
func (PcInputAndOutput) TableName() string {
  return "pc_input_and_output"
}

