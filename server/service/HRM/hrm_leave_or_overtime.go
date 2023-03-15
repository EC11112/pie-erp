package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    HRMReq "github.com/flipped-aurora/gin-vue-admin/server/model/HRM/request"
)

type HrmLeaveOrOvertimeService struct {
}

// CreateHrmLeaveOrOvertime 创建HrmLeaveOrOvertime记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmLeaveOrOvertimeService *HrmLeaveOrOvertimeService) CreateHrmLeaveOrOvertime(hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime) (err error) {
	err = global.GVA_DB.Create(&hrmLeaveOrOvertime).Error
	return err
}

// DeleteHrmLeaveOrOvertime 删除HrmLeaveOrOvertime记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmLeaveOrOvertimeService *HrmLeaveOrOvertimeService)DeleteHrmLeaveOrOvertime(hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime) (err error) {
	err = global.GVA_DB.Delete(&hrmLeaveOrOvertime).Error
	return err
}

// DeleteHrmLeaveOrOvertimeByIds 批量删除HrmLeaveOrOvertime记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmLeaveOrOvertimeService *HrmLeaveOrOvertimeService)DeleteHrmLeaveOrOvertimeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]HRM.HrmLeaveOrOvertime{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHrmLeaveOrOvertime 更新HrmLeaveOrOvertime记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmLeaveOrOvertimeService *HrmLeaveOrOvertimeService)UpdateHrmLeaveOrOvertime(hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime) (err error) {
	err = global.GVA_DB.Save(&hrmLeaveOrOvertime).Error
	return err
}

// GetHrmLeaveOrOvertime 根据id获取HrmLeaveOrOvertime记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmLeaveOrOvertimeService *HrmLeaveOrOvertimeService)GetHrmLeaveOrOvertime(id uint) (hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hrmLeaveOrOvertime).Error
	return
}

// GetHrmLeaveOrOvertimeInfoList 分页获取HrmLeaveOrOvertime记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrmLeaveOrOvertimeService *HrmLeaveOrOvertimeService)GetHrmLeaveOrOvertimeInfoList(info HRMReq.HrmLeaveOrOvertimeSearch) (list []HRM.HrmLeaveOrOvertime, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&HRM.HrmLeaveOrOvertime{})
    var hrmLeaveOrOvertimes []HRM.HrmLeaveOrOvertime
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.User_id != nil {
        db = db.Where("user_id = ?",info.User_id)
    }
    if info.Is_leave != "" {
        db = db.Where("is_leave = ?",info.Is_leave)
    }
        if info.StartStart_date != nil && info.EndStart_date != nil {
            db = db.Where("start_date BETWEEN ? AND ? ",info.StartStart_date,info.EndStart_date)
        }
        if info.StartEnd_date != nil && info.EndEnd_date != nil {
            db = db.Where("end_date BETWEEN ? AND ? ",info.StartEnd_date,info.EndEnd_date)
        }
    if info.Is_approved != "" {
        db = db.Where("is_approved = ?",info.Is_approved)
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
         	orderMap["user_id"] = true
         	orderMap["is_leave"] = true
         	orderMap["start_date"] = true
         	orderMap["end_date"] = true
         	orderMap["is_approved"] = true
         	orderMap["explanation"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&hrmLeaveOrOvertimes).Error
	return  hrmLeaveOrOvertimes, total, err
}
