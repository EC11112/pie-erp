import service from '@/utils/request'

// @Tags IcInventoryChanges
// @Summary 创建IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcInventoryChanges true "创建IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icInventoryChanges/createIcInventoryChanges [post]
export const createIcInventoryChanges = (data) => {
  return service({
    url: '/icInventoryChanges/createIcInventoryChanges',
    method: 'post',
    data
  })
}

// @Tags IcInventoryChanges
// @Summary 删除IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcInventoryChanges true "删除IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icInventoryChanges/deleteIcInventoryChanges [delete]
export const deleteIcInventoryChanges = (data) => {
  return service({
    url: '/icInventoryChanges/deleteIcInventoryChanges',
    method: 'delete',
    data
  })
}

// @Tags IcInventoryChanges
// @Summary 删除IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icInventoryChanges/deleteIcInventoryChanges [delete]
export const deleteIcInventoryChangesByIds = (data) => {
  return service({
    url: '/icInventoryChanges/deleteIcInventoryChangesByIds',
    method: 'delete',
    data
  })
}

// @Tags IcInventoryChanges
// @Summary 更新IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcInventoryChanges true "更新IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icInventoryChanges/updateIcInventoryChanges [put]
export const updateIcInventoryChanges = (data) => {
  return service({
    url: '/icInventoryChanges/updateIcInventoryChanges',
    method: 'put',
    data
  })
}

// @Tags IcInventoryChanges
// @Summary 用id查询IcInventoryChanges
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IcInventoryChanges true "用id查询IcInventoryChanges"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icInventoryChanges/findIcInventoryChanges [get]
export const findIcInventoryChanges = (params) => {
  return service({
    url: '/icInventoryChanges/findIcInventoryChanges',
    method: 'get',
    params
  })
}

// @Tags IcInventoryChanges
// @Summary 分页获取IcInventoryChanges列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IcInventoryChanges列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icInventoryChanges/getIcInventoryChangesList [get]
export const getIcInventoryChangesList = (params) => {
  return service({
    url: '/icInventoryChanges/getIcInventoryChangesList',
    method: 'get',
    params
  })
}
