package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ICReq "github.com/flipped-aurora/gin-vue-admin/server/model/IC/request"
)

type IcViewInventoryService struct {
}

// CreateIcViewInventory 创建IcViewInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewInventoryService *IcViewInventoryService) CreateIcViewInventory(icViewInventory IC.IcViewInventory) (err error) {
	err = global.GVA_DB.Create(&icViewInventory).Error
	return err
}

// DeleteIcViewInventory 删除IcViewInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewInventoryService *IcViewInventoryService)DeleteIcViewInventory(icViewInventory IC.IcViewInventory) (err error) {
	err = global.GVA_DB.Delete(&icViewInventory).Error
	return err
}

// DeleteIcViewInventoryByIds 批量删除IcViewInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewInventoryService *IcViewInventoryService)DeleteIcViewInventoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]IC.IcViewInventory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIcViewInventory 更新IcViewInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewInventoryService *IcViewInventoryService)UpdateIcViewInventory(icViewInventory IC.IcViewInventory) (err error) {
	err = global.GVA_DB.Save(&icViewInventory).Error
	return err
}

// GetIcViewInventory 根据id获取IcViewInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewInventoryService *IcViewInventoryService)GetIcViewInventory(id uint) (icViewInventory IC.IcViewInventory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&icViewInventory).Error
	return
}

// GetIcViewInventoryInfoList 分页获取IcViewInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewInventoryService *IcViewInventoryService)GetIcViewInventoryInfoList(info ICReq.IcViewInventorySearch) (list []IC.IcViewInventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&IC.IcViewInventory{})
    var icViewInventorys []IC.IcViewInventory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Item_id != nil {
        db = db.Where("item_id = ?",info.Item_id)
    }
    if info.Item_number != nil {
        db = db.Where("item_number = ?",info.Item_number)
    }
    if info.Warehouse_id != nil {
        db = db.Where("warehouse_id = ?",info.Warehouse_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["item_id"] = true
         	orderMap["item_number"] = true
         	orderMap["warehouse_id"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&icViewInventorys).Error
	return  icViewInventorys, total, err
}
