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

type FmViewOrderApi struct {
}

var fmViewOrderService = service.ServiceGroupApp.FMServiceGroup.FmViewOrderService


// CreateFmViewOrder 创建FmViewOrder
// @Tags FmViewOrder
// @Summary 创建FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmViewOrder true "创建FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmViewOrder/createFmViewOrder [post]
func (fmViewOrderApi *FmViewOrderApi) CreateFmViewOrder(c *gin.Context) {
	var fmViewOrder FM.FmViewOrder
	err := c.ShouldBindJSON(&fmViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(fmViewOrder, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := fmViewOrderService.CreateFmViewOrder(fmViewOrder); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFmViewOrder 删除FmViewOrder
// @Tags FmViewOrder
// @Summary 删除FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmViewOrder true "删除FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmViewOrder/deleteFmViewOrder [delete]
func (fmViewOrderApi *FmViewOrderApi) DeleteFmViewOrder(c *gin.Context) {
	var fmViewOrder FM.FmViewOrder
	err := c.ShouldBindJSON(&fmViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmViewOrderService.DeleteFmViewOrder(fmViewOrder); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFmViewOrderByIds 批量删除FmViewOrder
// @Tags FmViewOrder
// @Summary 批量删除FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fmViewOrder/deleteFmViewOrderByIds [delete]
func (fmViewOrderApi *FmViewOrderApi) DeleteFmViewOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmViewOrderService.DeleteFmViewOrderByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFmViewOrder 更新FmViewOrder
// @Tags FmViewOrder
// @Summary 更新FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmViewOrder true "更新FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmViewOrder/updateFmViewOrder [put]
func (fmViewOrderApi *FmViewOrderApi) UpdateFmViewOrder(c *gin.Context) {
	var fmViewOrder FM.FmViewOrder
	err := c.ShouldBindJSON(&fmViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(fmViewOrder, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := fmViewOrderService.UpdateFmViewOrder(fmViewOrder); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFmViewOrder 用id查询FmViewOrder
// @Tags FmViewOrder
// @Summary 用id查询FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FM.FmViewOrder true "用id查询FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmViewOrder/findFmViewOrder [get]
func (fmViewOrderApi *FmViewOrderApi) FindFmViewOrder(c *gin.Context) {
	var fmViewOrder FM.FmViewOrder
	err := c.ShouldBindQuery(&fmViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refmViewOrder, err := fmViewOrderService.GetFmViewOrder(fmViewOrder.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refmViewOrder": refmViewOrder}, c)
	}
}

// GetFmViewOrderList 分页获取FmViewOrder列表
// @Tags FmViewOrder
// @Summary 分页获取FmViewOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FMReq.FmViewOrderSearch true "分页获取FmViewOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmViewOrder/getFmViewOrderList [get]
func (fmViewOrderApi *FmViewOrderApi) GetFmViewOrderList(c *gin.Context) {
	var pageInfo FMReq.FmViewOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fmViewOrderService.GetFmViewOrderInfoList(pageInfo); err != nil {
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
