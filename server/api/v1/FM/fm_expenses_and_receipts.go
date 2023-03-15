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

type FmExpensesAndReceiptsApi struct {
}

var fmExpensesAndReceiptsService = service.ServiceGroupApp.FMServiceGroup.FmExpensesAndReceiptsService


// CreateFmExpensesAndReceipts 创建FmExpensesAndReceipts
// @Tags FmExpensesAndReceipts
// @Summary 创建FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmExpensesAndReceipts true "创建FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmExpensesAndReceipts/createFmExpensesAndReceipts [post]
func (fmExpensesAndReceiptsApi *FmExpensesAndReceiptsApi) CreateFmExpensesAndReceipts(c *gin.Context) {
	var fmExpensesAndReceipts FM.FmExpensesAndReceipts
	err := c.ShouldBindJSON(&fmExpensesAndReceipts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Expenses_or_receipts":{utils.NotEmpty()},
        "Amount":{utils.NotEmpty()},
    }
	if err := utils.Verify(fmExpensesAndReceipts, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := fmExpensesAndReceiptsService.CreateFmExpensesAndReceipts(fmExpensesAndReceipts); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFmExpensesAndReceipts 删除FmExpensesAndReceipts
// @Tags FmExpensesAndReceipts
// @Summary 删除FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmExpensesAndReceipts true "删除FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmExpensesAndReceipts/deleteFmExpensesAndReceipts [delete]
func (fmExpensesAndReceiptsApi *FmExpensesAndReceiptsApi) DeleteFmExpensesAndReceipts(c *gin.Context) {
	var fmExpensesAndReceipts FM.FmExpensesAndReceipts
	err := c.ShouldBindJSON(&fmExpensesAndReceipts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmExpensesAndReceiptsService.DeleteFmExpensesAndReceipts(fmExpensesAndReceipts); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFmExpensesAndReceiptsByIds 批量删除FmExpensesAndReceipts
// @Tags FmExpensesAndReceipts
// @Summary 批量删除FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fmExpensesAndReceipts/deleteFmExpensesAndReceiptsByIds [delete]
func (fmExpensesAndReceiptsApi *FmExpensesAndReceiptsApi) DeleteFmExpensesAndReceiptsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmExpensesAndReceiptsService.DeleteFmExpensesAndReceiptsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFmExpensesAndReceipts 更新FmExpensesAndReceipts
// @Tags FmExpensesAndReceipts
// @Summary 更新FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmExpensesAndReceipts true "更新FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmExpensesAndReceipts/updateFmExpensesAndReceipts [put]
func (fmExpensesAndReceiptsApi *FmExpensesAndReceiptsApi) UpdateFmExpensesAndReceipts(c *gin.Context) {
	var fmExpensesAndReceipts FM.FmExpensesAndReceipts
	err := c.ShouldBindJSON(&fmExpensesAndReceipts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Expenses_or_receipts":{utils.NotEmpty()},
          "Amount":{utils.NotEmpty()},
      }
    if err := utils.Verify(fmExpensesAndReceipts, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := fmExpensesAndReceiptsService.UpdateFmExpensesAndReceipts(fmExpensesAndReceipts); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFmExpensesAndReceipts 用id查询FmExpensesAndReceipts
// @Tags FmExpensesAndReceipts
// @Summary 用id查询FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FM.FmExpensesAndReceipts true "用id查询FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmExpensesAndReceipts/findFmExpensesAndReceipts [get]
func (fmExpensesAndReceiptsApi *FmExpensesAndReceiptsApi) FindFmExpensesAndReceipts(c *gin.Context) {
	var fmExpensesAndReceipts FM.FmExpensesAndReceipts
	err := c.ShouldBindQuery(&fmExpensesAndReceipts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refmExpensesAndReceipts, err := fmExpensesAndReceiptsService.GetFmExpensesAndReceipts(fmExpensesAndReceipts.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refmExpensesAndReceipts": refmExpensesAndReceipts}, c)
	}
}

// GetFmExpensesAndReceiptsList 分页获取FmExpensesAndReceipts列表
// @Tags FmExpensesAndReceipts
// @Summary 分页获取FmExpensesAndReceipts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FMReq.FmExpensesAndReceiptsSearch true "分页获取FmExpensesAndReceipts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmExpensesAndReceipts/getFmExpensesAndReceiptsList [get]
func (fmExpensesAndReceiptsApi *FmExpensesAndReceiptsApi) GetFmExpensesAndReceiptsList(c *gin.Context) {
	var pageInfo FMReq.FmExpensesAndReceiptsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fmExpensesAndReceiptsService.GetFmExpensesAndReceiptsInfoList(pageInfo); err != nil {
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
