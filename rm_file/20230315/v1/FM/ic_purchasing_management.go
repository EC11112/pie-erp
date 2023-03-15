package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/FM"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    FMReq "github.com/flipped-aurora/gin-vue-admin/server/model/FM/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type IcPurchasingManagementApi struct {
}

var icPurchasingManagementService = service.ServiceGroupApp.FMServiceGroup.IcPurchasingManagementService


// CreateIcPurchasingManagement 创建IcPurchasingManagement
// @Tags IcPurchasingManagement
// @Summary 创建IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.IcPurchasingManagement true "创建IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icPurchasingManagement/createIcPurchasingManagement [post]
func (icPurchasingManagementApi *IcPurchasingManagementApi) CreateIcPurchasingManagement(c *gin.Context) {
	var icPurchasingManagement FM.IcPurchasingManagement
	err := c.ShouldBindJSON(&icPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Warehouse_id":{utils.NotEmpty()},
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(icPurchasingManagement, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := icPurchasingManagementService.CreateIcPurchasingManagement(icPurchasingManagement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIcPurchasingManagement 删除IcPurchasingManagement
// @Tags IcPurchasingManagement
// @Summary 删除IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.IcPurchasingManagement true "删除IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icPurchasingManagement/deleteIcPurchasingManagement [delete]
func (icPurchasingManagementApi *IcPurchasingManagementApi) DeleteIcPurchasingManagement(c *gin.Context) {
	var icPurchasingManagement FM.IcPurchasingManagement
	err := c.ShouldBindJSON(&icPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icPurchasingManagementService.DeleteIcPurchasingManagement(icPurchasingManagement); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIcPurchasingManagementByIds 批量删除IcPurchasingManagement
// @Tags IcPurchasingManagement
// @Summary 批量删除IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /icPurchasingManagement/deleteIcPurchasingManagementByIds [delete]
func (icPurchasingManagementApi *IcPurchasingManagementApi) DeleteIcPurchasingManagementByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icPurchasingManagementService.DeleteIcPurchasingManagementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIcPurchasingManagement 更新IcPurchasingManagement
// @Tags IcPurchasingManagement
// @Summary 更新IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.IcPurchasingManagement true "更新IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icPurchasingManagement/updateIcPurchasingManagement [put]
func (icPurchasingManagementApi *IcPurchasingManagementApi) UpdateIcPurchasingManagement(c *gin.Context) {
	var icPurchasingManagement FM.IcPurchasingManagement
	err := c.ShouldBindJSON(&icPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Warehouse_id":{utils.NotEmpty()},
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(icPurchasingManagement, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := icPurchasingManagementService.UpdateIcPurchasingManagement(icPurchasingManagement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIcPurchasingManagement 用id查询IcPurchasingManagement
// @Tags IcPurchasingManagement
// @Summary 用id查询IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FM.IcPurchasingManagement true "用id查询IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icPurchasingManagement/findIcPurchasingManagement [get]
func (icPurchasingManagementApi *IcPurchasingManagementApi) FindIcPurchasingManagement(c *gin.Context) {
	var icPurchasingManagement FM.IcPurchasingManagement
	err := c.ShouldBindQuery(&icPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reicPurchasingManagement, err := icPurchasingManagementService.GetIcPurchasingManagement(icPurchasingManagement.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reicPurchasingManagement": reicPurchasingManagement}, c)
	}
}

// GetIcPurchasingManagementList 分页获取IcPurchasingManagement列表
// @Tags IcPurchasingManagement
// @Summary 分页获取IcPurchasingManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FMReq.IcPurchasingManagementSearch true "分页获取IcPurchasingManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icPurchasingManagement/getIcPurchasingManagementList [get]
func (icPurchasingManagementApi *IcPurchasingManagementApi) GetIcPurchasingManagementList(c *gin.Context) {
	var pageInfo FMReq.IcPurchasingManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := icPurchasingManagementService.GetIcPurchasingManagementInfoList(pageInfo); err != nil {
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
