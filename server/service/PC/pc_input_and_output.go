package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    PCReq "github.com/flipped-aurora/gin-vue-admin/server/model/PC/request"
)

type PcInputAndOutputService struct {
}

// CreatePcInputAndOutput 创建PcInputAndOutput记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcInputAndOutputService *PcInputAndOutputService) CreatePcInputAndOutput(pcInputAndOutput PC.PcInputAndOutput) (err error) {
	err = global.GVA_DB.Create(&pcInputAndOutput).Error
	return err
}

// DeletePcInputAndOutput 删除PcInputAndOutput记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcInputAndOutputService *PcInputAndOutputService)DeletePcInputAndOutput(pcInputAndOutput PC.PcInputAndOutput) (err error) {
	err = global.GVA_DB.Delete(&pcInputAndOutput).Error
	return err
}

// DeletePcInputAndOutputByIds 批量删除PcInputAndOutput记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcInputAndOutputService *PcInputAndOutputService)DeletePcInputAndOutputByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]PC.PcInputAndOutput{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePcInputAndOutput 更新PcInputAndOutput记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcInputAndOutputService *PcInputAndOutputService)UpdatePcInputAndOutput(pcInputAndOutput PC.PcInputAndOutput) (err error) {
	err = global.GVA_DB.Save(&pcInputAndOutput).Error
	return err
}

// GetPcInputAndOutput 根据id获取PcInputAndOutput记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcInputAndOutputService *PcInputAndOutputService)GetPcInputAndOutput(id uint) (pcInputAndOutput PC.PcInputAndOutput, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pcInputAndOutput).Error
	return
}

// GetPcInputAndOutputInfoList 分页获取PcInputAndOutput记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcInputAndOutputService *PcInputAndOutputService)GetPcInputAndOutputInfoList(info PCReq.PcInputAndOutputSearch) (list []PC.PcInputAndOutput, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&PC.PcInputAndOutput{})
    var pcInputAndOutputs []PC.PcInputAndOutput
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Process_id != nil {
        db = db.Where("process_id = ?",info.Process_id)
    }
    if info.Input_or_output != "" {
        db = db.Where("input_or_output = ?",info.Input_or_output)
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
         	orderMap["process_id"] = true
         	orderMap["input_or_output"] = true
         	orderMap["item_id"] = true
         	orderMap["item_number"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&pcInputAndOutputs).Error
	return  pcInputAndOutputs, total, err
}
