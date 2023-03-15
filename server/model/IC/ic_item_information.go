// 自动生成模板IcItemInformation
package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// IcItemInformation 结构体
type IcItemInformation struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:货物名称;size:64;"`
      Unit_price  *int `json:"unit_price" form:"unit_price" gorm:"column:unit_price;comment:单位价格;"`
}


// TableName IcItemInformation 表名
func (IcItemInformation) TableName() string {
  return "ic_item_information"
}

