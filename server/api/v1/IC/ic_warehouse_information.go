package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/IC"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ICReq "github.com/flipped-aurora/gin-vue-admin/server/model/IC/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type IcWarehouseInformationApi struct {
}

var icWarehouseInformationService = service.ServiceGroupApp.ICServiceGroup.IcWarehouseInformationService


// CreateIcWarehouseInformation 创建IcWarehouseInformation
// @Tags IcWarehouseInformation
// @Summary 创建IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcWarehouseInformation true "创建IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icWarehouseInformation/createIcWarehouseInformation [post]
func (icWarehouseInformationApi *IcWarehouseInformationApi) CreateIcWarehouseInformation(c *gin.Context) {
	var icWarehouseInformation IC.IcWarehouseInformation
	err := c.ShouldBindJSON(&icWarehouseInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
    }
	if err := utils.Verify(icWarehouseInformation, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := icWarehouseInformationService.CreateIcWarehouseInformation(icWarehouseInformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIcWarehouseInformation 删除IcWarehouseInformation
// @Tags IcWarehouseInformation
// @Summary 删除IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcWarehouseInformation true "删除IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icWarehouseInformation/deleteIcWarehouseInformation [delete]
func (icWarehouseInformationApi *IcWarehouseInformationApi) DeleteIcWarehouseInformation(c *gin.Context) {
	var icWarehouseInformation IC.IcWarehouseInformation
	err := c.ShouldBindJSON(&icWarehouseInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icWarehouseInformationService.DeleteIcWarehouseInformation(icWarehouseInformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIcWarehouseInformationByIds 批量删除IcWarehouseInformation
// @Tags IcWarehouseInformation
// @Summary 批量删除IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /icWarehouseInformation/deleteIcWarehouseInformationByIds [delete]
func (icWarehouseInformationApi *IcWarehouseInformationApi) DeleteIcWarehouseInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icWarehouseInformationService.DeleteIcWarehouseInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIcWarehouseInformation 更新IcWarehouseInformation
// @Tags IcWarehouseInformation
// @Summary 更新IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcWarehouseInformation true "更新IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icWarehouseInformation/updateIcWarehouseInformation [put]
func (icWarehouseInformationApi *IcWarehouseInformationApi) UpdateIcWarehouseInformation(c *gin.Context) {
	var icWarehouseInformation IC.IcWarehouseInformation
	err := c.ShouldBindJSON(&icWarehouseInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
      }
    if err := utils.Verify(icWarehouseInformation, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := icWarehouseInformationService.UpdateIcWarehouseInformation(icWarehouseInformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIcWarehouseInformation 用id查询IcWarehouseInformation
// @Tags IcWarehouseInformation
// @Summary 用id查询IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query IC.IcWarehouseInformation true "用id查询IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icWarehouseInformation/findIcWarehouseInformation [get]
func (icWarehouseInformationApi *IcWarehouseInformationApi) FindIcWarehouseInformation(c *gin.Context) {
	var icWarehouseInformation IC.IcWarehouseInformation
	err := c.ShouldBindQuery(&icWarehouseInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reicWarehouseInformation, err := icWarehouseInformationService.GetIcWarehouseInformation(icWarehouseInformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reicWarehouseInformation": reicWarehouseInformation}, c)
	}
}

// GetIcWarehouseInformationList 分页获取IcWarehouseInformation列表
// @Tags IcWarehouseInformation
// @Summary 分页获取IcWarehouseInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ICReq.IcWarehouseInformationSearch true "分页获取IcWarehouseInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icWarehouseInformation/getIcWarehouseInformationList [get]
func (icWarehouseInformationApi *IcWarehouseInformationApi) GetIcWarehouseInformationList(c *gin.Context) {
	var pageInfo ICReq.IcWarehouseInformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := icWarehouseInformationService.GetIcWarehouseInformationInfoList(pageInfo); err != nil {
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
