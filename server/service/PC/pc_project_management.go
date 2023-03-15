package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    PCReq "github.com/flipped-aurora/gin-vue-admin/server/model/PC/request"
)

type PcProjectManagementService struct {
}

// CreatePcProjectManagement 创建PcProjectManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProjectManagementService *PcProjectManagementService) CreatePcProjectManagement(pcProjectManagement PC.PcProjectManagement) (err error) {
	err = global.GVA_DB.Create(&pcProjectManagement).Error
	return err
}

// DeletePcProjectManagement 删除PcProjectManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProjectManagementService *PcProjectManagementService)DeletePcProjectManagement(pcProjectManagement PC.PcProjectManagement) (err error) {
	err = global.GVA_DB.Delete(&pcProjectManagement).Error
	return err
}

// DeletePcProjectManagementByIds 批量删除PcProjectManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProjectManagementService *PcProjectManagementService)DeletePcProjectManagementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]PC.PcProjectManagement{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePcProjectManagement 更新PcProjectManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProjectManagementService *PcProjectManagementService)UpdatePcProjectManagement(pcProjectManagement PC.PcProjectManagement) (err error) {
	err = global.GVA_DB.Save(&pcProjectManagement).Error
	return err
}

// GetPcProjectManagement 根据id获取PcProjectManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProjectManagementService *PcProjectManagementService)GetPcProjectManagement(id uint) (pcProjectManagement PC.PcProjectManagement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pcProjectManagement).Error
	return
}

// GetPcProjectManagementInfoList 分页获取PcProjectManagement记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcProjectManagementService *PcProjectManagementService)GetPcProjectManagementInfoList(info PCReq.PcProjectManagementSearch) (list []PC.PcProjectManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&PC.PcProjectManagement{})
    var pcProjectManagements []PC.PcProjectManagement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.User_id != nil {
        db = db.Where("user_id = ?",info.User_id)
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
         	orderMap["user_id"] = true
         	orderMap["explanation"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&pcProjectManagements).Error
	return  pcProjectManagements, total, err
}
