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

type HrmSalaryRulesApi struct {
}

var hrmSalaryRulesService = service.ServiceGroupApp.HRMServiceGroup.HrmSalaryRulesService


// CreateHrmSalaryRules 创建HrmSalaryRules
// @Tags HrmSalaryRules
// @Summary 创建HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmSalaryRules true "创建HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmSalaryRules/createHrmSalaryRules [post]
func (hrmSalaryRulesApi *HrmSalaryRulesApi) CreateHrmSalaryRules(c *gin.Context) {
	var hrmSalaryRules HRM.HrmSalaryRules
	err := c.ShouldBindJSON(&hrmSalaryRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Mean":{utils.NotEmpty()},
        "Price":{utils.NotEmpty()},
    }
	if err := utils.Verify(hrmSalaryRules, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := hrmSalaryRulesService.CreateHrmSalaryRules(hrmSalaryRules); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHrmSalaryRules 删除HrmSalaryRules
// @Tags HrmSalaryRules
// @Summary 删除HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmSalaryRules true "删除HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmSalaryRules/deleteHrmSalaryRules [delete]
func (hrmSalaryRulesApi *HrmSalaryRulesApi) DeleteHrmSalaryRules(c *gin.Context) {
	var hrmSalaryRules HRM.HrmSalaryRules
	err := c.ShouldBindJSON(&hrmSalaryRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hrmSalaryRulesService.DeleteHrmSalaryRules(hrmSalaryRules); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHrmSalaryRulesByIds 批量删除HrmSalaryRules
// @Tags HrmSalaryRules
// @Summary 批量删除HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hrmSalaryRules/deleteHrmSalaryRulesByIds [delete]
func (hrmSalaryRulesApi *HrmSalaryRulesApi) DeleteHrmSalaryRulesByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hrmSalaryRulesService.DeleteHrmSalaryRulesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHrmSalaryRules 更新HrmSalaryRules
// @Tags HrmSalaryRules
// @Summary 更新HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body HRM.HrmSalaryRules true "更新HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrmSalaryRules/updateHrmSalaryRules [put]
func (hrmSalaryRulesApi *HrmSalaryRulesApi) UpdateHrmSalaryRules(c *gin.Context) {
	var hrmSalaryRules HRM.HrmSalaryRules
	err := c.ShouldBindJSON(&hrmSalaryRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Mean":{utils.NotEmpty()},
          "Price":{utils.NotEmpty()},
      }
    if err := utils.Verify(hrmSalaryRules, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := hrmSalaryRulesService.UpdateHrmSalaryRules(hrmSalaryRules); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHrmSalaryRules 用id查询HrmSalaryRules
// @Tags HrmSalaryRules
// @Summary 用id查询HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query HRM.HrmSalaryRules true "用id查询HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrmSalaryRules/findHrmSalaryRules [get]
func (hrmSalaryRulesApi *HrmSalaryRulesApi) FindHrmSalaryRules(c *gin.Context) {
	var hrmSalaryRules HRM.HrmSalaryRules
	err := c.ShouldBindQuery(&hrmSalaryRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehrmSalaryRules, err := hrmSalaryRulesService.GetHrmSalaryRules(hrmSalaryRules.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehrmSalaryRules": rehrmSalaryRules}, c)
	}
}

// GetHrmSalaryRulesList 分页获取HrmSalaryRules列表
// @Tags HrmSalaryRules
// @Summary 分页获取HrmSalaryRules列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query HRMReq.HrmSalaryRulesSearch true "分页获取HrmSalaryRules列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmSalaryRules/getHrmSalaryRulesList [get]
func (hrmSalaryRulesApi *HrmSalaryRulesApi) GetHrmSalaryRulesList(c *gin.Context) {
	var pageInfo HRMReq.HrmSalaryRulesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hrmSalaryRulesService.GetHrmSalaryRulesInfoList(pageInfo); err != nil {
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
