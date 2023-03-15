package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    PCReq "github.com/flipped-aurora/gin-vue-admin/server/model/PC/request"
)

type PcProductionProcessService struct {
}

// CreatePcProductionProcess 创建PcProductionProcess记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProductionProcessService *PcProductionProcessService) CreatePcProductionProcess(pcProductionProcess PC.PcProductionProcess) (err error) {
	err = global.GVA_DB.Create(&pcProductionProcess).Error
	return err
}

// DeletePcProductionProcess 删除PcProductionProcess记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProductionProcessService *PcProductionProcessService)DeletePcProductionProcess(pcProductionProcess PC.PcProductionProcess) (err error) {
	err = global.GVA_DB.Delete(&pcProductionProcess).Error
	return err
}

// DeletePcProductionProcessByIds 批量删除PcProductionProcess记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProductionProcessService *PcProductionProcessService)DeletePcProductionProcessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]PC.PcProductionProcess{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePcProductionProcess 更新PcProductionProcess记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProductionProcessService *PcProductionProcessService)UpdatePcProductionProcess(pcProductionProcess PC.PcProductionProcess) (err error) {
	err = global.GVA_DB.Save(&pcProductionProcess).Error
	return err
}

// GetPcProductionProcess 根据id获取PcProductionProcess记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProductionProcessService *PcProductionProcessService)GetPcProductionProcess(id uint) (pcProductionProcess PC.PcProductionProcess, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pcProductionProcess).Error
	return
}

// GetPcProductionProcessInfoList 分页获取PcProductionProcess记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProductionProcessService *PcProductionProcessService)GetPcProductionProcessInfoList(info PCReq.PcProductionProcessSearch) (list []PC.PcProductionProcess, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&PC.PcProductionProcess{})
    var pcProductionProcesss []PC.PcProductionProcess
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Project_id != nil {
        db = db.Where("project_id = ?",info.Project_id)
    }
    if info.User_id != nil {
        db = db.Where("user_id = ?",info.User_id)
    }
    if info.Production_cycle != nil {
        db = db.Where("production_cycle = ?",info.Production_cycle)
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
         	orderMap["name"] = true
         	orderMap["project_id"] = true
         	orderMap["user_id"] = true
         	orderMap["production_cycle"] = true
         	orderMap["explanation"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&pcProductionProcesss).Error
	return  pcProductionProcesss, total, err
}
