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

type IcInventoryChangesApi struct {
}

var icInventoryChangesService = service.ServiceGroupApp.ICServiceGroup.IcInventoryChangesService


// CreateIcInventoryChanges 创建IcInventoryChanges
// @Tags IcInventoryChanges
// @Summary 创建IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcInventoryChanges true "创建IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icInventoryChanges/createIcInventoryChanges [post]
func (icInventoryChangesApi *IcInventoryChangesApi) CreateIcInventoryChanges(c *gin.Context) {
	var icInventoryChanges IC.IcInventoryChanges
	err := c.ShouldBindJSON(&icInventoryChanges)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Warehouse_id":{utils.NotEmpty()},
        "Change_type":{utils.NotEmpty()},
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(icInventoryChanges, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := icInventoryChangesService.CreateIcInventoryChanges(icInventoryChanges); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIcInventoryChanges 删除IcInventoryChanges
// @Tags IcInventoryChanges
// @Summary 删除IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcInventoryChanges true "删除IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icInventoryChanges/deleteIcInventoryChanges [delete]
func (icInventoryChangesApi *IcInventoryChangesApi) DeleteIcInventoryChanges(c *gin.Context) {
	var icInventoryChanges IC.IcInventoryChanges
	err := c.ShouldBindJSON(&icInventoryChanges)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icInventoryChangesService.DeleteIcInventoryChanges(icInventoryChanges); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIcInventoryChangesByIds 批量删除IcInventoryChanges
// @Tags IcInventoryChanges
// @Summary 批量删除IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /icInventoryChanges/deleteIcInventoryChangesByIds [delete]
func (icInventoryChangesApi *IcInventoryChangesApi) DeleteIcInventoryChangesByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icInventoryChangesService.DeleteIcInventoryChangesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIcInventoryChanges 更新IcInventoryChanges
// @Tags IcInventoryChanges
// @Summary 更新IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcInventoryChanges true "更新IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icInventoryChanges/updateIcInventoryChanges [put]
func (icInventoryChangesApi *IcInventoryChangesApi) UpdateIcInventoryChanges(c *gin.Context) {
	var icInventoryChanges IC.IcInventoryChanges
	err := c.ShouldBindJSON(&icInventoryChanges)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Warehouse_id":{utils.NotEmpty()},
          "Change_type":{utils.NotEmpty()},
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(icInventoryChanges, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := icInventoryChangesService.UpdateIcInventoryChanges(icInventoryChanges); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIcInventoryChanges 用id查询IcInventoryChanges
// @Tags IcInventoryChanges
// @Summary 用id查询IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query IC.IcInventoryChanges true "用id查询IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icInventoryChanges/findIcInventoryChanges [get]
func (icInventoryChangesApi *IcInventoryChangesApi) FindIcInventoryChanges(c *gin.Context) {
	var icInventoryChanges IC.IcInventoryChanges
	err := c.ShouldBindQuery(&icInventoryChanges)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reicInventoryChanges, err := icInventoryChangesService.GetIcInventoryChanges(icInventoryChanges.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reicInventoryChanges": reicInventoryChanges}, c)
	}
}

// GetIcInventoryChangesList 分页获取IcInventoryChanges列表
// @Tags IcInventoryChanges
// @Summary 分页获取IcInventoryChanges列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ICReq.IcInventoryChangesSearch true "分页获取IcInventoryChanges列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icInventoryChanges/getIcInventoryChangesList [get]
func (icInventoryChangesApi *IcInventoryChangesApi) GetIcInventoryChangesList(c *gin.Context) {
	var pageInfo ICReq.IcInventoryChangesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := icInventoryChangesService.GetIcInventoryChangesInfoList(pageInfo); err != nil {
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
