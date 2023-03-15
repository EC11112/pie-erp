package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ICReq "github.com/flipped-aurora/gin-vue-admin/server/model/IC/request"
)

type IcItemInformationService struct {
}

// CreateIcItemInformation 创建IcItemInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icItemInformationService *IcItemInformationService) CreateIcItemInformation(icItemInformation IC.IcItemInformation) (err error) {
	err = global.GVA_DB.Create(&icItemInformation).Error
	return err
}

// DeleteIcItemInformation 删除IcItemInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icItemInformationService *IcItemInformationService)DeleteIcItemInformation(icItemInformation IC.IcItemInformation) (err error) {
	err = global.GVA_DB.Delete(&icItemInformation).Error
	return err
}

// DeleteIcItemInformationByIds 批量删除IcItemInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icItemInformationService *IcItemInformationService)DeleteIcItemInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]IC.IcItemInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIcItemInformation 更新IcItemInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icItemInformationService *IcItemInformationService)UpdateIcItemInformation(icItemInformation IC.IcItemInformation) (err error) {
	err = global.GVA_DB.Save(&icItemInformation).Error
	return err
}

// GetIcItemInformation 根据id获取IcItemInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icItemInformationService *IcItemInformationService)GetIcItemInformation(id uint) (icItemInformation IC.IcItemInformation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&icItemInformation).Error
	return
}

// GetIcItemInformationInfoList 分页获取IcItemInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icItemInformationService *IcItemInformationService)GetIcItemInformationInfoList(info ICReq.IcItemInformationSearch) (list []IC.IcItemInformation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&IC.IcItemInformation{})
    var icItemInformations []IC.IcItemInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Unit_price != nil {
        db = db.Where("unit_price = ?",info.Unit_price)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
         	orderMap["unit_price"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&icItemInformations).Error
	return  icItemInformations, total, err
}
