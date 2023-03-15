import service from '@/utils/request'

// @Tags PcInputAndOutput
// @Summary 创建PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcInputAndOutput true "创建PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcInputAndOutput/createPcInputAndOutput [post]
export const createPcInputAndOutput = (data) => {
  return service({
    url: '/pcInputAndOutput/createPcInputAndOutput',
    method: 'post',
    data
  })
}

// @Tags PcInputAndOutput
// @Summary 删除PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcInputAndOutput true "删除PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcInputAndOutput/deletePcInputAndOutput [delete]
export const deletePcInputAndOutput = (data) => {
  return service({
    url: '/pcInputAndOutput/deletePcInputAndOutput',
    method: 'delete',
    data
  })
}

// @Tags PcInputAndOutput
// @Summary 删除PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcInputAndOutput/deletePcInputAndOutput [delete]
export const deletePcInputAndOutputByIds = (data) => {
  return service({
    url: '/pcInputAndOutput/deletePcInputAndOutputByIds',
    method: 'delete',
    data
  })
}

// @Tags PcInputAndOutput
// @Summary 更新PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcInputAndOutput true "更新PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcInputAndOutput/updatePcInputAndOutput [put]
export const updatePcInputAndOutput = (data) => {
  return service({
    url: '/pcInputAndOutput/updatePcInputAndOutput',
    method: 'put',
    data
  })
}

// @Tags PcInputAndOutput
// @Summary 用id查询PcInputAndOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PcInputAndOutput true "用id查询PcInputAndOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcInputAndOutput/findPcInputAndOutput [get]
export const findPcInputAndOutput = (params) => {
  return service({
    url: '/pcInputAndOutput/findPcInputAndOutput',
    method: 'get',
    params
  })
}

// @Tags PcInputAndOutput
// @Summary 分页获取PcInputAndOutput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PcInputAndOutput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcInputAndOutput/getPcInputAndOutputList [get]
export const getPcInputAndOutputList = (params) => {
  return service({
    url: '/pcInputAndOutput/getPcInputAndOutputList',
    method: 'get',
    params
  })
}
