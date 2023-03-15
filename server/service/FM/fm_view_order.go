package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
)

type FmViewOrderService struct {
}

// CreateFmViewOrder 创建FmViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmViewOrderService *FmViewOrderService) CreateFmViewOrder(fmViewOrder FM.FmViewOrder) (err error) {
	err = global.GVA_DB.Create(&fmViewOrder).Error
	return err
}

// DeleteFmViewOrder 删除FmViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmViewOrderService *FmViewOrderService)DeleteFmViewOrder(fmViewOrder FM.FmViewOrder) (err error) {
	err = global.GVA_DB.Delete(&fmViewOrder).Error
	return err
}

// DeleteFmViewOrderByIds 批量删除FmViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmViewOrderService *FmViewOrderService)DeleteFmViewOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]FM.FmViewOrder{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFmViewOrder 更新FmViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmViewOrderService *FmViewOrderService)UpdateFmViewOrder(fmViewOrder FM.FmViewOrder) (err error) {
	err = global.GVA_DB.Save(&fmViewOrder).Error
	return err
}

// GetFmViewOrder 根据id获取FmViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmViewOrderService *FmViewOrderService)GetFmViewOrder(id uint) (fmViewOrder FM.FmViewOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fmViewOrder).Error
	return
}

// GetFmViewOrderInfoList 分页获取FmViewOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmViewOrderService *FmViewOrderService)GetFmViewOrderInfoList(info FMReq.FmViewOrderSearch) (list []FM.FmViewOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&FM.FmViewOrder{})
    var fmViewOrders []FM.FmViewOrder
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

	err = db.Limit(limit).Offset(offset).Find(&fmViewOrders).Error
	return  fmViewOrders, total, err
}
