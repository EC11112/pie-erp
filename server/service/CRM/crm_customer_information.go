package CRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/CRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    CRMReq "github.com/flipped-aurora/gin-vue-admin/server/model/CRM/request"
)

type CrmCustomerInformationService struct {
}

// CreateCrmCustomerInformation 创建CrmCustomerInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerInformationService *CrmCustomerInformationService) CreateCrmCustomerInformation(crmCustomerInformation CRM.CrmCustomerInformation) (err error) {
	err = global.GVA_DB.Create(&crmCustomerInformation).Error
	return err
}

// DeleteCrmCustomerInformation 删除CrmCustomerInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerInformationService *CrmCustomerInformationService)DeleteCrmCustomerInformation(crmCustomerInformation CRM.CrmCustomerInformation) (err error) {
	err = global.GVA_DB.Delete(&crmCustomerInformation).Error
	return err
}

// DeleteCrmCustomerInformationByIds 批量删除CrmCustomerInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerInformationService *CrmCustomerInformationService)DeleteCrmCustomerInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]CRM.CrmCustomerInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCrmCustomerInformation 更新CrmCustomerInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerInformationService *CrmCustomerInformationService)UpdateCrmCustomerInformation(crmCustomerInformation CRM.CrmCustomerInformation) (err error) {
	err = global.GVA_DB.Save(&crmCustomerInformation).Error
	return err
}

// GetCrmCustomerInformation 根据id获取CrmCustomerInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerInformationService *CrmCustomerInformationService)GetCrmCustomerInformation(id uint) (crmCustomerInformation CRM.CrmCustomerInformation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&crmCustomerInformation).Error
	return
}

// GetCrmCustomerInformationInfoList 分页获取CrmCustomerInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerInformationService *CrmCustomerInformationService)GetCrmCustomerInformationInfoList(info CRMReq.CrmCustomerInformationSearch) (list []CRM.CrmCustomerInformation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&CRM.CrmCustomerInformation{})
    var crmCustomerInformations []CRM.CrmCustomerInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Organizational_context != "" {
        db = db.Where("organizational_context LIKE ?","%"+ info.Organizational_context+"%")
    }
    if info.User_id != nil {
        db = db.Where("user_id = ?",info.User_id)
    }
    if info.Phone_number != nil {
        db = db.Where("phone_number = ?",info.Phone_number)
    }
    if info.Email != "" {
        db = db.Where("email LIKE ?","%"+ info.Email+"%")
    }
    if info.Explanation != "" {
        db = db.Where("explanation LIKE ?","%"+ info.Explanation+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
         	orderMap["organizational_context"] = true
         	orderMap["user_id"] = true
         	orderMap["phone_number"] = true
         	orderMap["email"] = true
         	orderMap["explanation"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&crmCustomerInformations).Error
	return  crmCustomerInformations, total, err
}
