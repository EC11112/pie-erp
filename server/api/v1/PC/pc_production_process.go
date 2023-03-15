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

type PcProductionProcessApi struct {
}

var pcProductionProcessService = service.ServiceGroupApp.PCServiceGroup.PcProductionProcessService


// CreatePcProductionProcess 创建PcProductionProcess
// @Tags PcProductionProcess
// @Summary 创建PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcProductionProcess true "创建PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProductionProcess/createPcProductionProcess [post]
func (pcProductionProcessApi *PcProductionProcessApi) CreatePcProductionProcess(c *gin.Context) {
	var pcProductionProcess PC.PcProductionProcess
	err := c.ShouldBindJSON(&pcProductionProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Project_id":{utils.NotEmpty()},
        "User_id":{utils.NotEmpty()},
        "Production_cycle":{utils.NotEmpty()},
    }
	if err := utils.Verify(pcProductionProcess, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := pcProductionProcessService.CreatePcProductionProcess(pcProductionProcess); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePcProductionProcess 删除PcProductionProcess
// @Tags PcProductionProcess
// @Summary 删除PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcProductionProcess true "删除PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcProductionProcess/deletePcProductionProcess [delete]
func (pcProductionProcessApi *PcProductionProcessApi) DeletePcProductionProcess(c *gin.Context) {
	var pcProductionProcess PC.PcProductionProcess
	err := c.ShouldBindJSON(&pcProductionProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcProductionProcessService.DeletePcProductionProcess(pcProductionProcess); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePcProductionProcessByIds 批量删除PcProductionProcess
// @Tags PcProductionProcess
// @Summary 批量删除PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pcProductionProcess/deletePcProductionProcessByIds [delete]
func (pcProductionProcessApi *PcProductionProcessApi) DeletePcProductionProcessByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcProductionProcessService.DeletePcProductionProcessByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePcProductionProcess 更新PcProductionProcess
// @Tags PcProductionProcess
// @Summary 更新PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcProductionProcess true "更新PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcProductionProcess/updatePcProductionProcess [put]
func (pcProductionProcessApi *PcProductionProcessApi) UpdatePcProductionProcess(c *gin.Context) {
	var pcProductionProcess PC.PcProductionProcess
	err := c.ShouldBindJSON(&pcProductionProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Project_id":{utils.NotEmpty()},
          "User_id":{utils.NotEmpty()},
          "Production_cycle":{utils.NotEmpty()},
      }
    if err := utils.Verify(pcProductionProcess, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := pcProductionProcessService.UpdatePcProductionProcess(pcProductionProcess); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPcProductionProcess 用id查询PcProductionProcess
// @Tags PcProductionProcess
// @Summary 用id查询PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PC.PcProductionProcess true "用id查询PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcProductionProcess/findPcProductionProcess [get]
func (pcProductionProcessApi *PcProductionProcessApi) FindPcProductionProcess(c *gin.Context) {
	var pcProductionProcess PC.PcProductionProcess
	err := c.ShouldBindQuery(&pcProductionProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repcProductionProcess, err := pcProductionProcessService.GetPcProductionProcess(pcProductionProcess.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repcProductionProcess": repcProductionProcess}, c)
	}
}

// GetPcProductionProcessList 分页获取PcProductionProcess列表
// @Tags PcProductionProcess
// @Summary 分页获取PcProductionProcess列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PCReq.PcProductionProcessSearch true "分页获取PcProductionProcess列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProductionProcess/getPcProductionProcessList [get]
func (pcProductionProcessApi *PcProductionProcessApi) GetPcProductionProcessList(c *gin.Context) {
	var pageInfo PCReq.PcProductionProcessSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := pcProductionProcessService.GetPcProductionProcessInfoList(pageInfo); err != nil {
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
