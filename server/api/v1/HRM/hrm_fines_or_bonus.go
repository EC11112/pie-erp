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

type HrmFinesOrBonusApi struct {
}

var hrmFinesOrBonusService = service.ServiceGroupApp.HRMServiceGroup.HrmFinesOrBonusService


// CreateHrmFinesOrBonus 创建HrmFinesOrBonus
// @Tags HrmFinesOrBonus
// @Summary 创建HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmFinesOrBonus true "创建HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmFinesOrBonus/createHrmFinesOrBonus [post]
func (hrmFinesOrBonusApi *HrmFinesOrBonusApi) CreateHrmFinesOrBonus(c *gin.Context) {
	var hrmFinesOrBonus HRM.HrmFinesOrBonus
	err := c.ShouldBindJSON(&hrmFinesOrBonus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Action":{utils.NotEmpty()},
        "User_id":{utils.NotEmpty()},
        "Amount":{utils.NotEmpty()},
    }
	if err := utils.Verify(hrmFinesOrBonus, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := hrmFinesOrBonusService.CreateHrmFinesOrBonus(hrmFinesOrBonus); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHrmFinesOrBonus 删除HrmFinesOrBonus
// @Tags HrmFinesOrBonus
// @Summary 删除HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmFinesOrBonus true "删除HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmFinesOrBonus/deleteHrmFinesOrBonus [delete]
func (hrmFinesOrBonusApi *HrmFinesOrBonusApi) DeleteHrmFinesOrBonus(c *gin.Context) {
	var hrmFinesOrBonus HRM.HrmFinesOrBonus
	err := c.ShouldBindJSON(&hrmFinesOrBonus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hrmFinesOrBonusService.DeleteHrmFinesOrBonus(hrmFinesOrBonus); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHrmFinesOrBonusByIds 批量删除HrmFinesOrBonus
// @Tags HrmFinesOrBonus
// @Summary 批量删除HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hrmFinesOrBonus/deleteHrmFinesOrBonusByIds [delete]
func (hrmFinesOrBonusApi *HrmFinesOrBonusApi) DeleteHrmFinesOrBonusByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hrmFinesOrBonusService.DeleteHrmFinesOrBonusByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHrmFinesOrBonus 更新HrmFinesOrBonus
// @Tags HrmFinesOrBonus
// @Summary 更新HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmFinesOrBonus true "更新HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrmFinesOrBonus/updateHrmFinesOrBonus [put]
func (hrmFinesOrBonusApi *HrmFinesOrBonusApi) UpdateHrmFinesOrBonus(c *gin.Context) {
	var hrmFinesOrBonus HRM.HrmFinesOrBonus
	err := c.ShouldBindJSON(&hrmFinesOrBonus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Action":{utils.NotEmpty()},
          "User_id":{utils.NotEmpty()},
          "Amount":{utils.NotEmpty()},
      }
    if err := utils.Verify(hrmFinesOrBonus, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := hrmFinesOrBonusService.UpdateHrmFinesOrBonus(hrmFinesOrBonus); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHrmFinesOrBonus 用id查询HrmFinesOrBonus
// @Tags HrmFinesOrBonus
// @Summary 用id查询HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query HRM.HrmFinesOrBonus true "用id查询HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrmFinesOrBonus/findHrmFinesOrBonus [get]
func (hrmFinesOrBonusApi *HrmFinesOrBonusApi) FindHrmFinesOrBonus(c *gin.Context) {
	var hrmFinesOrBonus HRM.HrmFinesOrBonus
	err := c.ShouldBindQuery(&hrmFinesOrBonus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehrmFinesOrBonus, err := hrmFinesOrBonusService.GetHrmFinesOrBonus(hrmFinesOrBonus.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehrmFinesOrBonus": rehrmFinesOrBonus}, c)
	}
}

// GetHrmFinesOrBonusList 分页获取HrmFinesOrBonus列表
// @Tags HrmFinesOrBonus
// @Summary 分页获取HrmFinesOrBonus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query HRMReq.HrmFinesOrBonusSearch true "分页获取HrmFinesOrBonus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmFinesOrBonus/getHrmFinesOrBonusList [get]
func (hrmFinesOrBonusApi *HrmFinesOrBonusApi) GetHrmFinesOrBonusList(c *gin.Context) {
	var pageInfo HRMReq.HrmFinesOrBonusSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hrmFinesOrBonusService.GetHrmFinesOrBonusInfoList(pageInfo); err != nil {
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
