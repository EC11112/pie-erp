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

type IcViewOrderApi struct {
}

var icViewOrderService = service.ServiceGroupApp.FMServiceGroup.IcViewOrderService


// CreateIcViewOrder 创建IcViewOrder
// @Tags IcViewOrder
// @Summary 创建IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.IcViewOrder true "创建IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewOrder/createIcViewOrder [post]
func (icViewOrderApi *IcViewOrderApi) CreateIcViewOrder(c *gin.Context) {
	var icViewOrder FM.IcViewOrder
	err := c.ShouldBindJSON(&icViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(icViewOrder, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := icViewOrderService.CreateIcViewOrder(icViewOrder); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIcViewOrder 删除IcViewOrder
// @Tags IcViewOrder
// @Summary 删除IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.IcViewOrder true "删除IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icViewOrder/deleteIcViewOrder [delete]
func (icViewOrderApi *IcViewOrderApi) DeleteIcViewOrder(c *gin.Context) {
	var icViewOrder FM.IcViewOrder
	err := c.ShouldBindJSON(&icViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icViewOrderService.DeleteIcViewOrder(icViewOrder); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIcViewOrderByIds 批量删除IcViewOrder
// @Tags IcViewOrder
// @Summary 批量删除IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /icViewOrder/deleteIcViewOrderByIds [delete]
func (icViewOrderApi *IcViewOrderApi) DeleteIcViewOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icViewOrderService.DeleteIcViewOrderByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIcViewOrder 更新IcViewOrder
// @Tags IcViewOrder
// @Summary 更新IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.IcViewOrder true "更新IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icViewOrder/updateIcViewOrder [put]
func (icViewOrderApi *IcViewOrderApi) UpdateIcViewOrder(c *gin.Context) {
	var icViewOrder FM.IcViewOrder
	err := c.ShouldBindJSON(&icViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(icViewOrder, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := icViewOrderService.UpdateIcViewOrder(icViewOrder); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIcViewOrder 用id查询IcViewOrder
// @Tags IcViewOrder
// @Summary 用id查询IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FM.IcViewOrder true "用id查询IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icViewOrder/findIcViewOrder [get]
func (icViewOrderApi *IcViewOrderApi) FindIcViewOrder(c *gin.Context) {
	var icViewOrder FM.IcViewOrder
	err := c.ShouldBindQuery(&icViewOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reicViewOrder, err := icViewOrderService.GetIcViewOrder(icViewOrder.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reicViewOrder": reicViewOrder}, c)
	}
}

// GetIcViewOrderList 分页获取IcViewOrder列表
// @Tags IcViewOrder
// @Summary 分页获取IcViewOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FMReq.IcViewOrderSearch true "分页获取IcViewOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewOrder/getIcViewOrderList [get]
func (icViewOrderApi *IcViewOrderApi) GetIcViewOrderList(c *gin.Context) {
	var pageInfo FMReq.IcViewOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := icViewOrderService.GetIcViewOrderInfoList(pageInfo); err != nil {
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
