import service from '@/utils/request'

// @Tags IcWarehouseInformation
// @Summary 创建IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcWarehouseInformation true "创建IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icWarehouseInformation/createIcWarehouseInformation [post]
export const createIcWarehouseInformation = (data) => {
  return service({
    url: '/icWarehouseInformation/createIcWarehouseInformation',
    method: 'post',
    data
  })
}

// @Tags IcWarehouseInformation
// @Summary 删除IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcWarehouseInformation true "删除IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icWarehouseInformation/deleteIcWarehouseInformation [delete]
export const deleteIcWarehouseInformation = (data) => {
  return service({
    url: '/icWarehouseInformation/deleteIcWarehouseInformation',
    method: 'delete',
    data
  })
}

// @Tags IcWarehouseInformation
// @Summary 删除IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icWarehouseInformation/deleteIcWarehouseInformation [delete]
export const deleteIcWarehouseInformationByIds = (data) => {
  return service({
    url: '/icWarehouseInformation/deleteIcWarehouseInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags IcWarehouseInformation
// @Summary 更新IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcWarehouseInformation true "更新IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icWarehouseInformation/updateIcWarehouseInformation [put]
export const updateIcWarehouseInformation = (data) => {
  return service({
    url: '/icWarehouseInformation/updateIcWarehouseInformation',
    method: 'put',
    data
  })
}

// @Tags IcWarehouseInformation
// @Summary 用id查询IcWarehouseInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IcWarehouseInformation true "用id查询IcWarehouseInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icWarehouseInformation/findIcWarehouseInformation [get]
export const findIcWarehouseInformation = (params) => {
  return service({
    url: '/icWarehouseInformation/findIcWarehouseInformation',
    method: 'get',
    params
  })
}

// @Tags IcWarehouseInformation
// @Summary 分页获取IcWarehouseInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IcWarehouseInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icWarehouseInformation/getIcWarehouseInformationList [get]
export const getIcWarehouseInformationList = (params) => {
  return service({
    url: '/icWarehouseInformation/getIcWarehouseInformationList',
    method: 'get',
    params
  })
}
