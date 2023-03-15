package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    HRMReq "github.com/flipped-aurora/gin-vue-admin/server/model/HRM/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HrmLeaveOrOvertimeApi struct {
}

var hrmLeaveOrOvertimeService = service.ServiceGroupApp.HRMServiceGroup.HrmLeaveOrOvertimeService


// CreateHrmLeaveOrOvertime 创建HrmLeaveOrOvertime
// @Tags HrmLeaveOrOvertime
// @Summary 创建HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmLeaveOrOvertime true "创建HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmLeaveOrOvertime/createHrmLeaveOrOvertime [post]
func (hrmLeaveOrOvertimeApi *HrmLeaveOrOvertimeApi) CreateHrmLeaveOrOvertime(c *gin.Context) {
	var hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime
	err := c.ShouldBindJSON(&hrmLeaveOrOvertime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "User_id":{utils.NotEmpty()},
        "Is_leave":{utils.NotEmpty()},
        "Start_date":{utils.NotEmpty()},
        "End_date":{utils.NotEmpty()},
        "Is_approved":{utils.NotEmpty()},
    }
	if err := utils.Verify(hrmLeaveOrOvertime, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := hrmLeaveOrOvertimeService.CreateHrmLeaveOrOvertime(hrmLeaveOrOvertime); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHrmLeaveOrOvertime 删除HrmLeaveOrOvertime
// @Tags HrmLeaveOrOvertime
// @Summary 删除HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmLeaveOrOvertime true "删除HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmLeaveOrOvertime/deleteHrmLeaveOrOvertime [delete]
func (hrmLeaveOrOvertimeApi *HrmLeaveOrOvertimeApi) DeleteHrmLeaveOrOvertime(c *gin.Context) {
	var hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime
	err := c.ShouldBindJSON(&hrmLeaveOrOvertime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hrmLeaveOrOvertimeService.DeleteHrmLeaveOrOvertime(hrmLeaveOrOvertime); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHrmLeaveOrOvertimeByIds 批量删除HrmLeaveOrOvertime
// @Tags HrmLeaveOrOvertime
// @Summary 批量删除HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hrmLeaveOrOvertime/deleteHrmLeaveOrOvertimeByIds [delete]
func (hrmLeaveOrOvertimeApi *HrmLeaveOrOvertimeApi) DeleteHrmLeaveOrOvertimeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hrmLeaveOrOvertimeService.DeleteHrmLeaveOrOvertimeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHrmLeaveOrOvertime 更新HrmLeaveOrOvertime
// @Tags HrmLeaveOrOvertime
// @Summary 更新HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmLeaveOrOvertime true "更新HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrmLeaveOrOvertime/updateHrmLeaveOrOvertime [put]
func (hrmLeaveOrOvertimeApi *HrmLeaveOrOvertimeApi) UpdateHrmLeaveOrOvertime(c *gin.Context) {
	var hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime
	err := c.ShouldBindJSON(&hrmLeaveOrOvertime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "User_id":{utils.NotEmpty()},
          "Is_leave":{utils.NotEmpty()},
          "Start_date":{utils.NotEmpty()},
          "End_date":{utils.NotEmpty()},
          "Is_approved":{utils.NotEmpty()},
      }
    if err := utils.Verify(hrmLeaveOrOvertime, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := hrmLeaveOrOvertimeService.UpdateHrmLeaveOrOvertime(hrmLeaveOrOvertime); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHrmLeaveOrOvertime 用id查询HrmLeaveOrOvertime
// @Tags HrmLeaveOrOvertime
// @Summary 用id查询HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query HRM.HrmLeaveOrOvertime true "用id查询HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrmLeaveOrOvertime/findHrmLeaveOrOvertime [get]
func (hrmLeaveOrOvertimeApi *HrmLeaveOrOvertimeApi) FindHrmLeaveOrOvertime(c *gin.Context) {
	var hrmLeaveOrOvertime HRM.HrmLeaveOrOvertime
	err := c.ShouldBindQuery(&hrmLeaveOrOvertime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehrmLeaveOrOvertime, err := hrmLeaveOrOvertimeService.GetHrmLeaveOrOvertime(hrmLeaveOrOvertime.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehrmLeaveOrOvertime": rehrmLeaveOrOvertime}, c)
	}
}

// GetHrmLeaveOrOvertimeList 分页获取HrmLeaveOrOvertime列表
// @Tags HrmLeaveOrOvertime
// @Summary 分页获取HrmLeaveOrOvertime列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query HRMReq.HrmLeaveOrOvertimeSearch true "分页获取HrmLeaveOrOvertime列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmLeaveOrOvertime/getHrmLeaveOrOvertimeList [get]
func (hrmLeaveOrOvertimeApi *HrmLeaveOrOvertimeApi) GetHrmLeaveOrOvertimeList(c *gin.Context) {
	var pageInfo HRMReq.HrmLeaveOrOvertimeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hrmLeaveOrOvertimeService.GetHrmLeaveOrOvertimeInfoList(pageInfo); err != nil {
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
