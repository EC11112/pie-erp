package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    HRMReq "github.com/flipped-aurora/gin-vue-admin/server/model/HRM/request"
)

type HrmFinesOrBonusService struct {
}

// CreateHrmFinesOrBonus 创建HrmFinesOrBonus记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmFinesOrBonusService *HrmFinesOrBonusService) CreateHrmFinesOrBonus(hrmFinesOrBonus HRM.HrmFinesOrBonus) (err error) {
	err = global.GVA_DB.Create(&hrmFinesOrBonus).Error
	return err
}

// DeleteHrmFinesOrBonus 删除HrmFinesOrBonus记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmFinesOrBonusService *HrmFinesOrBonusService)DeleteHrmFinesOrBonus(hrmFinesOrBonus HRM.HrmFinesOrBonus) (err error) {
	err = global.GVA_DB.Delete(&hrmFinesOrBonus).Error
	return err
}

// DeleteHrmFinesOrBonusByIds 批量删除HrmFinesOrBonus记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmFinesOrBonusService *HrmFinesOrBonusService)DeleteHrmFinesOrBonusByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]HRM.HrmFinesOrBonus{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHrmFinesOrBonus 更新HrmFinesOrBonus记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmFinesOrBonusService *HrmFinesOrBonusService)UpdateHrmFinesOrBonus(hrmFinesOrBonus HRM.HrmFinesOrBonus) (err error) {
	err = global.GVA_DB.Save(&hrmFinesOrBonus).Error
	return err
}

// GetHrmFinesOrBonus 根据id获取HrmFinesOrBonus记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmFinesOrBonusService *HrmFinesOrBonusService)GetHrmFinesOrBonus(id uint) (hrmFinesOrBonus HRM.HrmFinesOrBonus, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hrmFinesOrBonus).Error
	return
}

// GetHrmFinesOrBonusInfoList 分页获取HrmFinesOrBonus记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmFinesOrBonusService *HrmFinesOrBonusService)GetHrmFinesOrBonusInfoList(info HRMReq.HrmFinesOrBonusSearch) (list []HRM.HrmFinesOrBonus, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&HRM.HrmFinesOrBonus{})
    var hrmFinesOrBonuss []HRM.HrmFinesOrBonus
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Action != "" {
        db = db.Where("action = ?",info.Action)
    }
    if info.User_id != nil {
        db = db.Where("user_id = ?",info.User_id)
    }
    if info.Amount != nil {
        db = db.Where("amount = ?",info.Amount)
    }
    if info.Explanation != "" {
        db = db.Where("explanation LIKE ?","%"+ info.Explanation+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["action"] = true
         	orderMap["user_id"] = true
         	orderMap["amount"] = true
         	orderMap["explanation"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&hrmFinesOrBonuss).Error
	return  hrmFinesOrBonuss, total, err
}
