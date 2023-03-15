import service from '@/utils/request'

// @Tags PcProductionProcess
// @Summary 创建PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcProductionProcess true "创建PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProductionProcess/createPcProductionProcess [post]
export const createPcProductionProcess = (data) => {
  return service({
    url: '/pcProductionProcess/createPcProductionProcess',
    method: 'post',
    data
  })
}

// @Tags PcProductionProcess
// @Summary 删除PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcProductionProcess true "删除PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcProductionProcess/deletePcProductionProcess [delete]
export const deletePcProductionProcess = (data) => {
  return service({
    url: '/pcProductionProcess/deletePcProductionProcess',
    method: 'delete',
    data
  })
}

// @Tags PcProductionProcess
// @Summary 删除PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcProductionProcess/deletePcProductionProcess [delete]
export const deletePcProductionProcessByIds = (data) => {
  return service({
    url: '/pcProductionProcess/deletePcProductionProcessByIds',
    method: 'delete',
    data
  })
}

// @Tags PcProductionProcess
// @Summary 更新PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcProductionProcess true "更新PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcProductionProcess/updatePcProductionProcess [put]
export const updatePcProductionProcess = (data) => {
  return service({
    url: '/pcProductionProcess/updatePcProductionProcess',
    method: 'put',
    data
  })
}

// @Tags PcProductionProcess
// @Summary 用id查询PcProductionProcess
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PcProductionProcess true "用id查询PcProductionProcess"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcProductionProcess/findPcProductionProcess [get]
export const findPcProductionProcess = (params) => {
  return service({
    url: '/pcProductionProcess/findPcProductionProcess',
    method: 'get',
    params
  })
}

// @Tags PcProductionProcess
// @Summary 分页获取PcProductionProcess列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PcProductionProcess列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProductionProcess/getPcProductionProcessList [get]
export const getPcProductionProcessList = (params) => {
  return service({
    url: '/pcProductionProcess/getPcProductionProcessList',
    method: 'get',
    params
  })
}
