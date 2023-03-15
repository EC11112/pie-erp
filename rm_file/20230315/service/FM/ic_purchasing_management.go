package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
)

type IcPurchasingManagementService struct {
}

// CreateIcPurchasingManagement 创建IcPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (icPurchasingManagementService *IcPurchasingManagementService) CreateIcPurchasingManagement(icPurchasingManagement FM.IcPurchasingManagement) (err error) {
	err = global.GVA_DB.Create(&icPurchasingManagement).Error
	return err
}

// DeleteIcPurchasingManagement 删除IcPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (icPurchasingManagementService *IcPurchasingManagementService)DeleteIcPurchasingManagement(icPurchasingManagement FM.IcPurchasingManagement) (err error) {
	err = global.GVA_DB.Delete(&icPurchasingManagement).Error
	return err
}

// DeleteIcPurchasingManagementByIds 批量删除IcPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (icPurchasingManagementService *IcPurchasingManagementService)DeleteIcPurchasingManagementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]FM.IcPurchasingManagement{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIcPurchasingManagement 更新IcPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (icPurchasingManagementService *IcPurchasingManagementService)UpdateIcPurchasingManagement(icPurchasingManagement FM.IcPurchasingManagement) (err error) {
	err = global.GVA_DB.Save(&icPurchasingManagement).Error
	return err
}

// GetIcPurchasingManagement 根据id获取IcPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (icPurchasingManagementService *IcPurchasingManagementService)GetIcPurchasingManagement(id uint) (icPurchasingManagement FM.IcPurchasingManagement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&icPurchasingManagement).Error
	return
}

// GetIcPurchasingManagementInfoList 分页获取IcPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (icPurchasingManagementService *IcPurchasingManagementService)GetIcPurchasingManagementInfoList(info FMReq.IcPurchasingManagementSearch) (list []FM.IcPurchasingManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&FM.IcPurchasingManagement{})
    var icPurchasingManagements []FM.IcPurchasingManagement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Warehouse_id != nil {
        db = db.Where("warehouse_id = ?",info.Warehouse_id)
    }
    if info.Item_id != nil {
        db = db.Where("item_id = ?",info.Item_id)
    }
    if info.Item_number != nil {
        db = db.Where("item_number = ?",info.Item_number)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["warehouse_id"] = true
         	orderMap["item_id"] = true
         	orderMap["item_number"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&icPurchasingManagements).Error
	return  icPurchasingManagements, total, err
}
