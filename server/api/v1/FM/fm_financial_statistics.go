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

type FmFinancialStatisticsApi struct {
}

var fmFinancialStatisticsService = service.ServiceGroupApp.FMServiceGroup.FmFinancialStatisticsService


// CreateFmFinancialStatistics 创建FmFinancialStatistics
// @Tags FmFinancialStatistics
// @Summary 创建FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmFinancialStatistics true "创建FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmFinancialStatistics/createFmFinancialStatistics [post]
func (fmFinancialStatisticsApi *FmFinancialStatisticsApi) CreateFmFinancialStatistics(c *gin.Context) {
	var fmFinancialStatistics FM.FmFinancialStatistics
	err := c.ShouldBindJSON(&fmFinancialStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Mean":{utils.NotEmpty()},
        "Price":{utils.NotEmpty()},
    }
	if err := utils.Verify(fmFinancialStatistics, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := fmFinancialStatisticsService.CreateFmFinancialStatistics(fmFinancialStatistics); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFmFinancialStatistics 删除FmFinancialStatistics
// @Tags FmFinancialStatistics
// @Summary 删除FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmFinancialStatistics true "删除FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmFinancialStatistics/deleteFmFinancialStatistics [delete]
func (fmFinancialStatisticsApi *FmFinancialStatisticsApi) DeleteFmFinancialStatistics(c *gin.Context) {
	var fmFinancialStatistics FM.FmFinancialStatistics
	err := c.ShouldBindJSON(&fmFinancialStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmFinancialStatisticsService.DeleteFmFinancialStatistics(fmFinancialStatistics); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFmFinancialStatisticsByIds 批量删除FmFinancialStatistics
// @Tags FmFinancialStatistics
// @Summary 批量删除FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fmFinancialStatistics/deleteFmFinancialStatisticsByIds [delete]
func (fmFinancialStatisticsApi *FmFinancialStatisticsApi) DeleteFmFinancialStatisticsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fmFinancialStatisticsService.DeleteFmFinancialStatisticsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFmFinancialStatistics 更新FmFinancialStatistics
// @Tags FmFinancialStatistics
// @Summary 更新FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body FM.FmFinancialStatistics true "更新FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmFinancialStatistics/updateFmFinancialStatistics [put]
func (fmFinancialStatisticsApi *FmFinancialStatisticsApi) UpdateFmFinancialStatistics(c *gin.Context) {
	var fmFinancialStatistics FM.FmFinancialStatistics
	err := c.ShouldBindJSON(&fmFinancialStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Mean":{utils.NotEmpty()},
          "Price":{utils.NotEmpty()},
      }
    if err := utils.Verify(fmFinancialStatistics, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := fmFinancialStatisticsService.UpdateFmFinancialStatistics(fmFinancialStatistics); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFmFinancialStatistics 用id查询FmFinancialStatistics
// @Tags FmFinancialStatistics
// @Summary 用id查询FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FM.FmFinancialStatistics true "用id查询FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmFinancialStatistics/findFmFinancialStatistics [get]
func (fmFinancialStatisticsApi *FmFinancialStatisticsApi) FindFmFinancialStatistics(c *gin.Context) {
	var fmFinancialStatistics FM.FmFinancialStatistics
	err := c.ShouldBindQuery(&fmFinancialStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refmFinancialStatistics, err := fmFinancialStatisticsService.GetFmFinancialStatistics(fmFinancialStatistics.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refmFinancialStatistics": refmFinancialStatistics}, c)
	}
}

// GetFmFinancialStatisticsList 分页获取FmFinancialStatistics列表
// @Tags FmFinancialStatistics
// @Summary 分页获取FmFinancialStatistics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query FMReq.FmFinancialStatisticsSearch true "分页获取FmFinancialStatistics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmFinancialStatistics/getFmFinancialStatisticsList [get]
func (fmFinancialStatisticsApi *FmFinancialStatisticsApi) GetFmFinancialStatisticsList(c *gin.Context) {
	var pageInfo FMReq.FmFinancialStatisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fmFinancialStatisticsService.GetFmFinancialStatisticsInfoList(pageInfo); err != nil {
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
