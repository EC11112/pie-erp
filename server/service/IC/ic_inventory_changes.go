package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ICReq "github.com/flipped-aurora/gin-vue-admin/server/model/IC/request"
)

type IcInventoryChangesService struct {
}

// CreateIcInventoryChanges 创建IcInventoryChanges记录
// Author [piexlmax](https://github.com/piexlmax)
func (icInventoryChangesService *IcInventoryChangesService) CreateIcInventoryChanges(icInventoryChanges IC.IcInventoryChanges) (err error) {
	err = global.GVA_DB.Create(&icInventoryChanges).Error
	return err
}

// DeleteIcInventoryChanges 删除IcInventoryChanges记录
// Author [piexlmax](https://github.com/piexlmax)
func (icInventoryChangesService *IcInventoryChangesService)DeleteIcInventoryChanges(icInventoryChanges IC.IcInventoryChanges) (err error) {
	err = global.GVA_DB.Delete(&icInventoryChanges).Error
	return err
}

// DeleteIcInventoryChangesByIds 批量删除IcInventoryChanges记录
// Author [piexlmax](https://github.com/piexlmax)
func (icInventoryChangesService *IcInventoryChangesService)DeleteIcInventoryChangesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]IC.IcInventoryChanges{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIcInventoryChanges 更新IcInventoryChanges记录
// Author [piexlmax](https://github.com/piexlmax)
func (icInventoryChangesService *IcInventoryChangesService)UpdateIcInventoryChanges(icInventoryChanges IC.IcInventoryChanges) (err error) {
	err = global.GVA_DB.Save(&icInventoryChanges).Error
	return err
}

// GetIcInventoryChanges 根据id获取IcInventoryChanges记录
// Author [piexlmax](https://github.com/piexlmax)
func (icInventoryChangesService *IcInventoryChangesService)GetIcInventoryChanges(id uint) (icInventoryChanges IC.IcInventoryChanges, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&icInventoryChanges).Error
	return
}

// GetIcInventoryChangesInfoList 分页获取IcInventoryChanges记录
// Author [piexlmax](https://github.com/piexlmax)
func (icInventoryChangesService *IcInventoryChangesService)GetIcInventoryChangesInfoList(info ICReq.IcInventoryChangesSearch) (list []IC.IcInventoryChanges, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&IC.IcInventoryChanges{})
    var icInventoryChangess []IC.IcInventoryChanges
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Warehouse_id != nil {
        db = db.Where("warehouse_id = ?",info.Warehouse_id)
    }
    if info.Change_type != "" {
        db = db.Where("change_type = ?",info.Change_type)
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
         	orderMap["change_type"] = true
         	orderMap["item_id"] = true
         	orderMap["item_number"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&icInventoryChangess).Error
	return  icInventoryChangess, total, err
}
