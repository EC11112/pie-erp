package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
)

type FmExpensesAndReceiptsService struct {
}

// CreateFmExpensesAndReceipts 创建FmExpensesAndReceipts记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmExpensesAndReceiptsService *FmExpensesAndReceiptsService) CreateFmExpensesAndReceipts(fmExpensesAndReceipts FM.FmExpensesAndReceipts) (err error) {
	err = global.GVA_DB.Create(&fmExpensesAndReceipts).Error
	return err
}

// DeleteFmExpensesAndReceipts 删除FmExpensesAndReceipts记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmExpensesAndReceiptsService *FmExpensesAndReceiptsService)DeleteFmExpensesAndReceipts(fmExpensesAndReceipts FM.FmExpensesAndReceipts) (err error) {
	err = global.GVA_DB.Delete(&fmExpensesAndReceipts).Error
	return err
}

// DeleteFmExpensesAndReceiptsByIds 批量删除FmExpensesAndReceipts记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmExpensesAndReceiptsService *FmExpensesAndReceiptsService)DeleteFmExpensesAndReceiptsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]FM.FmExpensesAndReceipts{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFmExpensesAndReceipts 更新FmExpensesAndReceipts记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmExpensesAndReceiptsService *FmExpensesAndReceiptsService)UpdateFmExpensesAndReceipts(fmExpensesAndReceipts FM.FmExpensesAndReceipts) (err error) {
	err = global.GVA_DB.Save(&fmExpensesAndReceipts).Error
	return err
}

// GetFmExpensesAndReceipts 根据id获取FmExpensesAndReceipts记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmExpensesAndReceiptsService *FmExpensesAndReceiptsService)GetFmExpensesAndReceipts(id uint) (fmExpensesAndReceipts FM.FmExpensesAndReceipts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fmExpensesAndReceipts).Error
	return
}

// GetFmExpensesAndReceiptsInfoList 分页获取FmExpensesAndReceipts记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmExpensesAndReceiptsService *FmExpensesAndReceiptsService)GetFmExpensesAndReceiptsInfoList(info FMReq.FmExpensesAndReceiptsSearch) (list []FM.FmExpensesAndReceipts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&FM.FmExpensesAndReceipts{})
    var fmExpensesAndReceiptss []FM.FmExpensesAndReceipts
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Expenses_or_receipts != "" {
        db = db.Where("expenses_or_receipts = ?",info.Expenses_or_receipts)
    }
    if info.Amount != nil {
        db = db.Where("amount = ?",info.Amount)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["expenses_or_receipts"] = true
         	orderMap["amount"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&fmExpensesAndReceiptss).Error
	return  fmExpensesAndReceiptss, total, err
}
