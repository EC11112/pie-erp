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

type IcItemInformationApi struct {
}

var icItemInformationService = service.ServiceGroupApp.ICServiceGroup.IcItemInformationService


// CreateIcItemInformation 创建IcItemInformation
// @Tags IcItemInformation
// @Summary 创建IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcItemInformation true "创建IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icItemInformation/createIcItemInformation [post]
func (icItemInformationApi *IcItemInformationApi) CreateIcItemInformation(c *gin.Context) {
	var icItemInformation IC.IcItemInformation
	err := c.ShouldBindJSON(&icItemInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Unit_price":{utils.NotEmpty()},
    }
	if err := utils.Verify(icItemInformation, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := icItemInformationService.CreateIcItemInformation(icItemInformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIcItemInformation 删除IcItemInformation
// @Tags IcItemInformation
// @Summary 删除IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcItemInformation true "删除IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icItemInformation/deleteIcItemInformation [delete]
func (icItemInformationApi *IcItemInformationApi) DeleteIcItemInformation(c *gin.Context) {
	var icItemInformation IC.IcItemInformation
	err := c.ShouldBindJSON(&icItemInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icItemInformationService.DeleteIcItemInformation(icItemInformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIcItemInformationByIds 批量删除IcItemInformation
// @Tags IcItemInformation
// @Summary 批量删除IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /icItemInformation/deleteIcItemInformationByIds [delete]
func (icItemInformationApi *IcItemInformationApi) DeleteIcItemInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := icItemInformationService.DeleteIcItemInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIcItemInformation 更新IcItemInformation
// @Tags IcItemInformation
// @Summary 更新IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body IC.IcItemInformation true "更新IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icItemInformation/updateIcItemInformation [put]
func (icItemInformationApi *IcItemInformationApi) UpdateIcItemInformation(c *gin.Context) {
	var icItemInformation IC.IcItemInformation
	err := c.ShouldBindJSON(&icItemInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Unit_price":{utils.NotEmpty()},
      }
    if err := utils.Verify(icItemInformation, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := icItemInformationService.UpdateIcItemInformation(icItemInformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIcItemInformation 用id查询IcItemInformation
// @Tags IcItemInformation
// @Summary 用id查询IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query IC.IcItemInformation true "用id查询IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icItemInformation/findIcItemInformation [get]
func (icItemInformationApi *IcItemInformationApi) FindIcItemInformation(c *gin.Context) {
	var icItemInformation IC.IcItemInformation
	err := c.ShouldBindQuery(&icItemInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reicItemInformation, err := icItemInformationService.GetIcItemInformation(icItemInformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reicItemInformation": reicItemInformation}, c)
	}
}

// GetIcItemInformationList 分页获取IcItemInformation列表
// @Tags IcItemInformation
// @Summary 分页获取IcItemInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ICReq.IcItemInformationSearch true "分页获取IcItemInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icItemInformation/getIcItemInformationList [get]
func (icItemInformationApi *IcItemInformationApi) GetIcItemInformationList(c *gin.Context) {
	var pageInfo ICReq.IcItemInformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := icItemInformationService.GetIcItemInformationInfoList(pageInfo); err != nil {
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
