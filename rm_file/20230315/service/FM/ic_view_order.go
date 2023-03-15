package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
)

type IcViewOrderService struct {
}

// CreateIcViewOrder 创建IcViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewOrderService *IcViewOrderService) CreateIcViewOrder(icViewOrder FM.IcViewOrder) (err error) {
	err = global.GVA_DB.Create(&icViewOrder).Error
	return err
}

// DeleteIcViewOrder 删除IcViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewOrderService *IcViewOrderService)DeleteIcViewOrder(icViewOrder FM.IcViewOrder) (err error) {
	err = global.GVA_DB.Delete(&icViewOrder).Error
	return err
}

// DeleteIcViewOrderByIds 批量删除IcViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewOrderService *IcViewOrderService)DeleteIcViewOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]FM.IcViewOrder{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIcViewOrder 更新IcViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewOrderService *IcViewOrderService)UpdateIcViewOrder(icViewOrder FM.IcViewOrder) (err error) {
	err = global.GVA_DB.Save(&icViewOrder).Error
	return err
}

// GetIcViewOrder 根据id获取IcViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewOrderService *IcViewOrderService)GetIcViewOrder(id uint) (icViewOrder FM.IcViewOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&icViewOrder).Error
	return
}

// GetIcViewOrderInfoList 分页获取IcViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (icViewOrderService *IcViewOrderService)GetIcViewOrderInfoList(info FMReq.IcViewOrderSearch) (list []FM.IcViewOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&FM.IcViewOrder{})
    var icViewOrders []FM.IcViewOrder
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
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["item_id"] = true
         	orderMap["item_number"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&icViewOrders).Error
	return  icViewOrders, total, err
}
