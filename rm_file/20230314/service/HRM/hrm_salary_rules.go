package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    HRMReq "github.com/flipped-aurora/gin-vue-admin/server/model/HRM/request"
)

type HrmSalaryRulesService struct {
}

// CreateHrmSalaryRules 创建HrmSalaryRules记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmSalaryRulesService *HrmSalaryRulesService) CreateHrmSalaryRules(hrmSalaryRules HRM.HrmSalaryRules) (err error) {
	err = global.GVA_DB.Create(&hrmSalaryRules).Error
	return err
}

// DeleteHrmSalaryRules 删除HrmSalaryRules记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmSalaryRulesService *HrmSalaryRulesService)DeleteHrmSalaryRules(hrmSalaryRules HRM.HrmSalaryRules) (err error) {
	err = global.GVA_DB.Delete(&hrmSalaryRules).Error
	return err
}

// DeleteHrmSalaryRulesByIds 批量删除HrmSalaryRules记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmSalaryRulesService *HrmSalaryRulesService)DeleteHrmSalaryRulesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]HRM.HrmSalaryRules{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHrmSalaryRules 更新HrmSalaryRules记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmSalaryRulesService *HrmSalaryRulesService)UpdateHrmSalaryRules(hrmSalaryRules HRM.HrmSalaryRules) (err error) {
	err = global.GVA_DB.Save(&hrmSalaryRules).Error
	return err
}

// GetHrmSalaryRules 根据id获取HrmSalaryRules记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmSalaryRulesService *HrmSalaryRulesService)GetHrmSalaryRules(id uint) (hrmSalaryRules HRM.HrmSalaryRules, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hrmSalaryRules).Error
	return
}

// GetHrmSalaryRulesInfoList 分页获取HrmSalaryRules记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmSalaryRulesService *HrmSalaryRulesService)GetHrmSalaryRulesInfoList(info HRMReq.HrmSalaryRulesSearch) (list []HRM.HrmSalaryRules, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&HRM.HrmSalaryRules{})
    var hrmSalaryRuless []HRM.HrmSalaryRules
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Mean != "" {
        db = db.Where("mean = ?",info.Mean)
    }
    if info.Price != nil {
        db = db.Where("price = ?",info.Price)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hrmSalaryRuless).Error
	return  hrmSalaryRuless, total, err
}
