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

type PcInputAndOutputApi struct {
}

var pcInputAndOutputService = service.ServiceGroupApp.PCServiceGroup.PcInputAndOutputService


// CreatePcInputAndOutput 创建PcInputAndOutput
// @Tags PcInputAndOutput
// @Summary 创建PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcInputAndOutput true "创建PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcInputAndOutput/createPcInputAndOutput [post]
func (pcInputAndOutputApi *PcInputAndOutputApi) CreatePcInputAndOutput(c *gin.Context) {
	var pcInputAndOutput PC.PcInputAndOutput
	err := c.ShouldBindJSON(&pcInputAndOutput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Process_id":{utils.NotEmpty()},
        "Input_or_output":{utils.NotEmpty()},
        "Item_id":{utils.NotEmpty()},
        "Item_number":{utils.NotEmpty()},
    }
	if err := utils.Verify(pcInputAndOutput, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := pcInputAndOutputService.CreatePcInputAndOutput(pcInputAndOutput); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePcInputAndOutput 删除PcInputAndOutput
// @Tags PcInputAndOutput
// @Summary 删除PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcInputAndOutput true "删除PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcInputAndOutput/deletePcInputAndOutput [delete]
func (pcInputAndOutputApi *PcInputAndOutputApi) DeletePcInputAndOutput(c *gin.Context) {
	var pcInputAndOutput PC.PcInputAndOutput
	err := c.ShouldBindJSON(&pcInputAndOutput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcInputAndOutputService.DeletePcInputAndOutput(pcInputAndOutput); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePcInputAndOutputByIds 批量删除PcInputAndOutput
// @Tags PcInputAndOutput
// @Summary 批量删除PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pcInputAndOutput/deletePcInputAndOutputByIds [delete]
func (pcInputAndOutputApi *PcInputAndOutputApi) DeletePcInputAndOutputByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcInputAndOutputService.DeletePcInputAndOutputByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePcInputAndOutput 更新PcInputAndOutput
// @Tags PcInputAndOutput
// @Summary 更新PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PC.PcInputAndOutput true "更新PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcInputAndOutput/updatePcInputAndOutput [put]
func (pcInputAndOutputApi *PcInputAndOutputApi) UpdatePcInputAndOutput(c *gin.Context) {
	var pcInputAndOutput PC.PcInputAndOutput
	err := c.ShouldBindJSON(&pcInputAndOutput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Process_id":{utils.NotEmpty()},
          "Input_or_output":{utils.NotEmpty()},
          "Item_id":{utils.NotEmpty()},
          "Item_number":{utils.NotEmpty()},
      }
    if err := utils.Verify(pcInputAndOutput, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := pcInputAndOutputService.UpdatePcInputAndOutput(pcInputAndOutput); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPcInputAndOutput 用id查询PcInputAndOutput
// @Tags PcInputAndOutput
// @Summary 用id查询PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PC.PcInputAndOutput true "用id查询PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcInputAndOutput/findPcInputAndOutput [get]
func (pcInputAndOutputApi *PcInputAndOutputApi) FindPcInputAndOutput(c *gin.Context) {
	var pcInputAndOutput PC.PcInputAndOutput
	err := c.ShouldBindQuery(&pcInputAndOutput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repcInputAndOutput, err := pcInputAndOutputService.GetPcInputAndOutput(pcInputAndOutput.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repcInputAndOutput": repcInputAndOutput}, c)
	}
}

// GetPcInputAndOutputList 分页获取PcInputAndOutput列表
// @Tags PcInputAndOutput
// @Summary 分页获取PcInputAndOutput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PCReq.PcInputAndOutputSearch true "分页获取PcInputAndOutput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcInputAndOutput/getPcInputAndOutputList [get]
func (pcInputAndOutputApi *PcInputAndOutputApi) GetPcInputAndOutputList(c *gin.Context) {
	var pageInfo PCReq.PcInputAndOutputSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := pcInputAndOutputService.GetPcInputAndOutputInfoList(pageInfo); err != nil {
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
