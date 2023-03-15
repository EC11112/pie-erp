package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
)

type FmPurchasingManagementService struct {
}

// CreateFmPurchasingManagement 创建FmPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmPurchasingManagementService *FmPurchasingManagementService) CreateFmPurchasingManagement(fmPurchasingManagement FM.FmPurchasingManagement) (err error) {
	err = global.GVA_DB.Create(&fmPurchasingManagement).Error
	return err
}

// DeleteFmPurchasingManagement 删除FmPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmPurchasingManagementService *FmPurchasingManagementService)DeleteFmPurchasingManagement(fmPurchasingManagement FM.FmPurchasingManagement) (err error) {
	err = global.GVA_DB.Delete(&fmPurchasingManagement).Error
	return err
}

// DeleteFmPurchasingManagementByIds 批量删除FmPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmPurchasingManagementService *FmPurchasingManagementService)DeleteFmPurchasingManagementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]FM.FmPurchasingManagement{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFmPurchasingManagement 更新FmPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmPurchasingManagementService *FmPurchasingManagementService)UpdateFmPurchasingManagement(fmPurchasingManagement FM.FmPurchasingManagement) (err error) {
	err = global.GVA_DB.Save(&fmPurchasingManagement).Error
	return err
}

// GetFmPurchasingManagement 根据id获取FmPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmPurchasingManagementService *FmPurchasingManagementService)GetFmPurchasingManagement(id uint) (fmPurchasingManagement FM.FmPurchasingManagement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fmPurchasingManagement).Error
	return
}

// GetFmPurchasingManagementInfoList 分页获取FmPurchasingManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmPurchasingManagementService *FmPurchasingManagementService)GetFmPurchasingManagementInfoList(info FMReq.FmPurchasingManagementSearch) (list []FM.FmPurchasingManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&FM.FmPurchasingManagement{})
    var fmPurchasingManagements []FM.FmPurchasingManagement
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

	err = db.Limit(limit).Offset(offset).Find(&fmPurchasingManagements).Error
	return  fmPurchasingManagements, total, err
}
