package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/IC"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ICReq "github.com/flipped-aurora/gin-vue-admin/server/model/IC/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type IcViewInventoryApi struct {
}

var icViewInventoryService = service.ServiceGroupApp.ICServiceGroup.IcViewInventoryService


// CreateIcViewInventory 创建IcViewInventory
// @Tags IcViewInventory
// @Summary 创建IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcViewInventory true "创建IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewInventory/createIcViewInventory [post]
func (icViewInventoryApi *IcViewInventoryApi) CreateIcViewInventory(c *gin.Context) {
	var icViewInventory IC.IcViewInventory
	err := c.ShouldBindJSON(&icViewInventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
        "Warehouse_id":{utils.NotEmpty()},
    }
	if err := utils.Verify(icViewInventory, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := icViewInventoryService.CreateIcViewInventory(icViewInventory); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIcViewInventory 删除IcViewInventory
// @Tags IcViewInventory
// @Summary 删除IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcViewInventory true "删除IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icViewInventory/deleteIcViewInventory [delete]
func (icViewInventoryApi *IcViewInventoryApi) DeleteIcViewInventory(c *gin.Context) {
	var icViewInventory IC.IcViewInventory
	err := c.ShouldBindJSON(&icViewInventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icViewInventoryService.DeleteIcViewInventory(icViewInventory); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIcViewInventoryByIds 批量删除IcViewInventory
// @Tags IcViewInventory
// @Summary 批量删除IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /icViewInventory/deleteIcViewInventoryByIds [delete]
func (icViewInventoryApi *IcViewInventoryApi) DeleteIcViewInventoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icViewInventoryService.DeleteIcViewInventoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIcViewInventory 更新IcViewInventory
// @Tags IcViewInventory
// @Summary 更新IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcViewInventory true "更新IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icViewInventory/updateIcViewInventory [put]
func (icViewInventoryApi *IcViewInventoryApi) UpdateIcViewInventory(c *gin.Context) {
	var icViewInventory IC.IcViewInventory
	err := c.ShouldBindJSON(&icViewInventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
          "Warehouse_id":{utils.NotEmpty()},
      }
    if err := utils.Verify(icViewInventory, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := icViewInventoryService.UpdateIcViewInventory(icViewInventory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIcViewInventory 用id查询IcViewInventory
// @Tags IcViewInventory
// @Summary 用id查询IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query IC.IcViewInventory true "用id查询IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icViewInventory/findIcViewInventory [get]
func (icViewInventoryApi *IcViewInventoryApi) FindIcViewInventory(c *gin.Context) {
	var icViewInventory IC.IcViewInventory
	err := c.ShouldBindQuery(&icViewInventory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reicViewInventory, err := icViewInventoryService.GetIcViewInventory(icViewInventory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reicViewInventory": reicViewInventory}, c)
	}
}

// GetIcViewInventoryList 分页获取IcViewInventory列表
// @Tags IcViewInventory
// @Summary 分页获取IcViewInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ICReq.IcViewInventorySearch true "分页获取IcViewInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewInventory/getIcViewInventoryList [get]
func (icViewInventoryApi *IcViewInventoryApi) GetIcViewInventoryList(c *gin.Context) {
	var pageInfo ICReq.IcViewInventorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := icViewInventoryService.GetIcViewInventoryInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
