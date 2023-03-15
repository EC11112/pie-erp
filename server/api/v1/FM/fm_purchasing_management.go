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

type FmPurchasingManagementApi struct {
}

var fmPurchasingManagementService = service.ServiceGroupApp.FMServiceGroup.FmPurchasingManagementService


// CreateFmPurchasingManagement 创建FmPurchasingManagement
// @Tags FmPurchasingManagement
// @Summary 创建FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmPurchasingManagement true "创建FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmPurchasingManagement/createFmPurchasingManagement [post]
func (fmPurchasingManagementApi *FmPurchasingManagementApi) CreateFmPurchasingManagement(c *gin.Context) {
	var fmPurchasingManagement FM.FmPurchasingManagement
	err := c.ShouldBindJSON(&fmPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Warehouse_id":{utils.NotEmpty()},
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(fmPurchasingManagement, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := fmPurchasingManagementService.CreateFmPurchasingManagement(fmPurchasingManagement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFmPurchasingManagement 删除FmPurchasingManagement
// @Tags FmPurchasingManagement
// @Summary 删除FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmPurchasingManagement true "删除FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmPurchasingManagement/deleteFmPurchasingManagement [delete]
func (fmPurchasingManagementApi *FmPurchasingManagementApi) DeleteFmPurchasingManagement(c *gin.Context) {
	var fmPurchasingManagement FM.FmPurchasingManagement
	err := c.ShouldBindJSON(&fmPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmPurchasingManagementService.DeleteFmPurchasingManagement(fmPurchasingManagement); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFmPurchasingManagementByIds 批量删除FmPurchasingManagement
// @Tags FmPurchasingManagement
// @Summary 批量删除FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fmPurchasingManagement/deleteFmPurchasingManagementByIds [delete]
func (fmPurchasingManagementApi *FmPurchasingManagementApi) DeleteFmPurchasingManagementByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmPurchasingManagementService.DeleteFmPurchasingManagementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFmPurchasingManagement 更新FmPurchasingManagement
// @Tags FmPurchasingManagement
// @Summary 更新FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmPurchasingManagement true "更新FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmPurchasingManagement/updateFmPurchasingManagement [put]
func (fmPurchasingManagementApi *FmPurchasingManagementApi) UpdateFmPurchasingManagement(c *gin.Context) {
	var fmPurchasingManagement FM.FmPurchasingManagement
	err := c.ShouldBindJSON(&fmPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Warehouse_id":{utils.NotEmpty()},
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(fmPurchasingManagement, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := fmPurchasingManagementService.UpdateFmPurchasingManagement(fmPurchasingManagement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFmPurchasingManagement 用id查询FmPurchasingManagement
// @Tags FmPurchasingManagement
// @Summary 用id查询FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FM.FmPurchasingManagement true "用id查询FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmPurchasingManagement/findFmPurchasingManagement [get]
func (fmPurchasingManagementApi *FmPurchasingManagementApi) FindFmPurchasingManagement(c *gin.Context) {
	var fmPurchasingManagement FM.FmPurchasingManagement
	err := c.ShouldBindQuery(&fmPurchasingManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refmPurchasingManagement, err := fmPurchasingManagementService.GetFmPurchasingManagement(fmPurchasingManagement.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refmPurchasingManagement": refmPurchasingManagement}, c)
	}
}

// GetFmPurchasingManagementList 分页获取FmPurchasingManagement列表
// @Tags FmPurchasingManagement
// @Summary 分页获取FmPurchasingManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FMReq.FmPurchasingManagementSearch true "分页获取FmPurchasingManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmPurchasingManagement/getFmPurchasingManagementList [get]
func (fmPurchasingManagementApi *FmPurchasingManagementApi) GetFmPurchasingManagementList(c *gin.Context) {
	var pageInfo FMReq.FmPurchasingManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fmPurchasingManagementService.GetFmPurchasingManagementInfoList(pageInfo); err != nil {
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
