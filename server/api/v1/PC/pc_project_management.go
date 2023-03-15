package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/PC"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    PCReq "github.com/flipped-aurora/gin-vue-admin/server/model/PC/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PcProjectManagementApi struct {
}

var pcProjectManagementService = service.ServiceGroupApp.PCServiceGroup.PcProjectManagementService


// CreatePcProjectManagement 创建PcProjectManagement
// @Tags PcProjectManagement
// @Summary 创建PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcProjectManagement true "创建PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProjectManagement/createPcProjectManagement [post]
func (pcProjectManagementApi *PcProjectManagementApi) CreatePcProjectManagement(c *gin.Context) {
	var pcProjectManagement PC.PcProjectManagement
	err := c.ShouldBindJSON(&pcProjectManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "User_id":{utils.NotEmpty()},
    }
	if err := utils.Verify(pcProjectManagement, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := pcProjectManagementService.CreatePcProjectManagement(pcProjectManagement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePcProjectManagement 删除PcProjectManagement
// @Tags PcProjectManagement
// @Summary 删除PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcProjectManagement true "删除PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcProjectManagement/deletePcProjectManagement [delete]
func (pcProjectManagementApi *PcProjectManagementApi) DeletePcProjectManagement(c *gin.Context) {
	var pcProjectManagement PC.PcProjectManagement
	err := c.ShouldBindJSON(&pcProjectManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcProjectManagementService.DeletePcProjectManagement(pcProjectManagement); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePcProjectManagementByIds 批量删除PcProjectManagement
// @Tags PcProjectManagement
// @Summary 批量删除PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pcProjectManagement/deletePcProjectManagementByIds [delete]
func (pcProjectManagementApi *PcProjectManagementApi) DeletePcProjectManagementByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcProjectManagementService.DeletePcProjectManagementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePcProjectManagement 更新PcProjectManagement
// @Tags PcProjectManagement
// @Summary 更新PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcProjectManagement true "更新PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcProjectManagement/updatePcProjectManagement [put]
func (pcProjectManagementApi *PcProjectManagementApi) UpdatePcProjectManagement(c *gin.Context) {
	var pcProjectManagement PC.PcProjectManagement
	err := c.ShouldBindJSON(&pcProjectManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "User_id":{utils.NotEmpty()},
      }
    if err := utils.Verify(pcProjectManagement, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := pcProjectManagementService.UpdatePcProjectManagement(pcProjectManagement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPcProjectManagement 用id查询PcProjectManagement
// @Tags PcProjectManagement
// @Summary 用id查询PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PC.PcProjectManagement true "用id查询PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcProjectManagement/findPcProjectManagement [get]
func (pcProjectManagementApi *PcProjectManagementApi) FindPcProjectManagement(c *gin.Context) {
	var pcProjectManagement PC.PcProjectManagement
	err := c.ShouldBindQuery(&pcProjectManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repcProjectManagement, err := pcProjectManagementService.GetPcProjectManagement(pcProjectManagement.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repcProjectManagement": repcProjectManagement}, c)
	}
}

// GetPcProjectManagementList 分页获取PcProjectManagement列表
// @Tags PcProjectManagement
// @Summary 分页获取PcProjectManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PCReq.PcProjectManagementSearch true "分页获取PcProjectManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProjectManagement/getPcProjectManagementList [get]
func (pcProjectManagementApi *PcProjectManagementApi) GetPcProjectManagementList(c *gin.Context) {
	var pageInfo PCReq.PcProjectManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := pcProjectManagementService.GetPcProjectManagementInfoList(pageInfo); err != nil {
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
