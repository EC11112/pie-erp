package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ICReq "github.com/flipped-aurora/gin-vue-admin/server/model/IC/request"
)

type IcWarehouseInformationService struct {
}

// CreateIcWarehouseInformation 创建IcWarehouseInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icWarehouseInformationService *IcWarehouseInformationService) CreateIcWarehouseInformation(icWarehouseInformation IC.IcWarehouseInformation) (err error) {
	err = global.GVA_DB.Create(&icWarehouseInformation).Error
	return err
}

// DeleteIcWarehouseInformation 删除IcWarehouseInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icWarehouseInformationService *IcWarehouseInformationService)DeleteIcWarehouseInformation(icWarehouseInformation IC.IcWarehouseInformation) (err error) {
	err = global.GVA_DB.Delete(&icWarehouseInformation).Error
	return err
}

// DeleteIcWarehouseInformationByIds 批量删除IcWarehouseInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icWarehouseInformationService *IcWarehouseInformationService)DeleteIcWarehouseInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]IC.IcWarehouseInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIcWarehouseInformation 更新IcWarehouseInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icWarehouseInformationService *IcWarehouseInformationService)UpdateIcWarehouseInformation(icWarehouseInformation IC.IcWarehouseInformation) (err error) {
	err = global.GVA_DB.Save(&icWarehouseInformation).Error
	return err
}

// GetIcWarehouseInformation 根据id获取IcWarehouseInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icWarehouseInformationService *IcWarehouseInformationService)GetIcWarehouseInformation(id uint) (icWarehouseInformation IC.IcWarehouseInformation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&icWarehouseInformation).Error
	return
}

// GetIcWarehouseInformationInfoList 分页获取IcWarehouseInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (icWarehouseInformationService *IcWarehouseInformationService)GetIcWarehouseInformationInfoList(info ICReq.IcWarehouseInformationSearch) (list []IC.IcWarehouseInformation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&IC.IcWarehouseInformation{})
    var icWarehouseInformations []IC.IcWarehouseInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&icWarehouseInformations).Error
	return  icWarehouseInformations, total, err
}
