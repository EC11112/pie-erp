// 自动生成模板CrmCustomerInformation
package CRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// CrmCustomerInformation 结构体
type CrmCustomerInformation struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:客户姓名;size:32;"`
      Organizational_context  string `json:"organizational_context" form:"organizational_context" gorm:"column:organizational_context;comment:组织背景;size:64;"`
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户ID;"`
      Phone_number  *int `json:"phone_number" form:"phone_number" gorm:"column:phone_number;comment:联系电话;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:电子邮箱;size:64;"`
      Explanation  string `json:"explanation" form:"explanation" gorm:"column:explanation;comment:说明;size:400;"`
}


// TableName CrmCustomerInformation 表名
func (CrmCustomerInformation) TableName() string {
  return "crm_customer_information"
}

