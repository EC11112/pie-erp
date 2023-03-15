package CRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/CRM"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    CRMReq "github.com/flipped-aurora/gin-vue-admin/server/model/CRM/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CrmCustomerInformationApi struct {
}

var crmCustomerInformationService = service.ServiceGroupApp.CRMServiceGroup.CrmCustomerInformationService


// CreateCrmCustomerInformation 创建CrmCustomerInformation
// @Tags CrmCustomerInformation
// @Summary 创建CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body CRM.CrmCustomerInformation true "创建CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerInformation/createCrmCustomerInformation [post]
func (crmCustomerInformationApi *CrmCustomerInformationApi) CreateCrmCustomerInformation(c *gin.Context) {
	var crmCustomerInformation CRM.CrmCustomerInformation
	err := c.ShouldBindJSON(&crmCustomerInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Organizational_context":{utils.NotEmpty()},
        "User_id":{utils.NotEmpty()},
        "Phone_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(crmCustomerInformation, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := crmCustomerInformationService.CreateCrmCustomerInformation(crmCustomerInformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmCustomerInformation 删除CrmCustomerInformation
// @Tags CrmCustomerInformation
// @Summary 删除CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body CRM.CrmCustomerInformation true "删除CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomerInformation/deleteCrmCustomerInformation [delete]
func (crmCustomerInformationApi *CrmCustomerInformationApi) DeleteCrmCustomerInformation(c *gin.Context) {
	var crmCustomerInformation CRM.CrmCustomerInformation
	err := c.ShouldBindJSON(&crmCustomerInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := crmCustomerInformationService.DeleteCrmCustomerInformation(crmCustomerInformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmCustomerInformationByIds 批量删除CrmCustomerInformation
// @Tags CrmCustomerInformation
// @Summary 批量删除CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmCustomerInformation/deleteCrmCustomerInformationByIds [delete]
func (crmCustomerInformationApi *CrmCustomerInformationApi) DeleteCrmCustomerInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := crmCustomerInformationService.DeleteCrmCustomerInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmCustomerInformation 更新CrmCustomerInformation
// @Tags CrmCustomerInformation
// @Summary 更新CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body CRM.CrmCustomerInformation true "更新CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCustomerInformation/updateCrmCustomerInformation [put]
func (crmCustomerInformationApi *CrmCustomerInformationApi) UpdateCrmCustomerInformation(c *gin.Context) {
	var crmCustomerInformation CRM.CrmCustomerInformation
	err := c.ShouldBindJSON(&crmCustomerInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Organizational_context":{utils.NotEmpty()},
          "User_id":{utils.NotEmpty()},
          "Phone_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(crmCustomerInformation, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := crmCustomerInformationService.UpdateCrmCustomerInformation(crmCustomerInformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmCustomerInformation 用id查询CrmCustomerInformation
// @Tags CrmCustomerInformation
// @Summary 用id查询CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CRM.CrmCustomerInformation true "用id查询CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomerInformation/findCrmCustomerInformation [get]
func (crmCustomerInformationApi *CrmCustomerInformationApi) FindCrmCustomerInformation(c *gin.Context) {
	var crmCustomerInformation CRM.CrmCustomerInformation
	err := c.ShouldBindQuery(&crmCustomerInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recrmCustomerInformation, err := crmCustomerInformationService.GetCrmCustomerInformation(crmCustomerInformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCustomerInformation": recrmCustomerInformation}, c)
	}
}

// GetCrmCustomerInformationList 分页获取CrmCustomerInformation列表
// @Tags CrmCustomerInformation
// @Summary 分页获取CrmCustomerInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CRMReq.CrmCustomerInformationSearch true "分页获取CrmCustomerInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerInformation/getCrmCustomerInformationList [get]
func (crmCustomerInformationApi *CrmCustomerInformationApi) GetCrmCustomerInformationList(c *gin.Context) {
	var pageInfo CRMReq.CrmCustomerInformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmCustomerInformationService.GetCrmCustomerInformationInfoList(pageInfo); err != nil {
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
