package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
)

type FmFinancialStatisticsService struct {
}

// CreateFmFinancialStatistics 创建FmFinancialStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmFinancialStatisticsService *FmFinancialStatisticsService) CreateFmFinancialStatistics(fmFinancialStatistics FM.FmFinancialStatistics) (err error) {
	err = global.GVA_DB.Create(&fmFinancialStatistics).Error
	return err
}

// DeleteFmFinancialStatistics 删除FmFinancialStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmFinancialStatisticsService *FmFinancialStatisticsService)DeleteFmFinancialStatistics(fmFinancialStatistics FM.FmFinancialStatistics) (err error) {
	err = global.GVA_DB.Delete(&fmFinancialStatistics).Error
	return err
}

// DeleteFmFinancialStatisticsByIds 批量删除FmFinancialStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmFinancialStatisticsService *FmFinancialStatisticsService)DeleteFmFinancialStatisticsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]FM.FmFinancialStatistics{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFmFinancialStatistics 更新FmFinancialStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmFinancialStatisticsService *FmFinancialStatisticsService)UpdateFmFinancialStatistics(fmFinancialStatistics FM.FmFinancialStatistics) (err error) {
	err = global.GVA_DB.Save(&fmFinancialStatistics).Error
	return err
}

// GetFmFinancialStatistics 根据id获取FmFinancialStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmFinancialStatisticsService *FmFinancialStatisticsService)GetFmFinancialStatistics(id uint) (fmFinancialStatistics FM.FmFinancialStatistics, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fmFinancialStatistics).Error
	return
}

// GetFmFinancialStatisticsInfoList 分页获取FmFinancialStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (fmFinancialStatisticsService *FmFinancialStatisticsService)GetFmFinancialStatisticsInfoList(info FMReq.FmFinancialStatisticsSearch) (list []FM.FmFinancialStatistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&FM.FmFinancialStatistics{})
    var fmFinancialStatisticss []FM.FmFinancialStatistics
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Mean != "" {
        db = db.Where("mean LIKE ?","%"+ info.Mean+"%")
    }
    if info.Price != nil {
        db = db.Where("price = ?",info.Price)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["mean"] = true
         	orderMap["price"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&fmFinancialStatisticss).Error
	return  fmFinancialStatisticss, total, err
}
