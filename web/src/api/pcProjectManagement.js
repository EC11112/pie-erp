import service from '@/utils/request'

// @Tags PcProjectManagement
// @Summary 创建PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcProjectManagement true "创建PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProjectManagement/createPcProjectManagement [post]
export const createPcProjectManagement = (data) => {
  return service({
    url: '/pcProjectManagement/createPcProjectManagement',
    method: 'post',
    data
  })
}

// @Tags PcProjectManagement
// @Summary 删除PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcProjectManagement true "删除PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcProjectManagement/deletePcProjectManagement [delete]
export const deletePcProjectManagement = (data) => {
  return service({
    url: '/pcProjectManagement/deletePcProjectManagement',
    method: 'delete',
    data
  })
}

// @Tags PcProjectManagement
// @Summary 删除PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pcProjectManagement/deletePcProjectManagement [delete]
export const deletePcProjectManagementByIds = (data) => {
  return service({
    url: '/pcProjectManagement/deletePcProjectManagementByIds',
    method: 'delete',
    data
  })
}

// @Tags PcProjectManagement
// @Summary 更新PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcProjectManagement true "更新PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pcProjectManagement/updatePcProjectManagement [put]
export const updatePcProjectManagement = (data) => {
  return service({
    url: '/pcProjectManagement/updatePcProjectManagement',
    method: 'put',
    data
  })
}

// @Tags PcProjectManagement
// @Summary 用id查询PcProjectManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PcProjectManagement true "用id查询PcProjectManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pcProjectManagement/findPcProjectManagement [get]
export const findPcProjectManagement = (params) => {
  return service({
    url: '/pcProjectManagement/findPcProjectManagement',
    method: 'get',
    params
  })
}

// @Tags PcProjectManagement
// @Summary 分页获取PcProjectManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PcProjectManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pcProjectManagement/getPcProjectManagementList [get]
export const getPcProjectManagementList = (params) => {
  return service({
    url: '/pcProjectManagement/getPcProjectManagementList',
    method: 'get',
    params
  })
}
