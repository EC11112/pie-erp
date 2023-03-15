import service from '@/utils/request'

// @Tags IcPurchasingManagement
// @Summary 创建IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcPurchasingManagement true "创建IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icPurchasingManagement/createIcPurchasingManagement [post]
export const createIcPurchasingManagement = (data) => {
  return service({
    url: '/icPurchasingManagement/createIcPurchasingManagement',
    method: 'post',
    data
  })
}

// @Tags IcPurchasingManagement
// @Summary 删除IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcPurchasingManagement true "删除IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icPurchasingManagement/deleteIcPurchasingManagement [delete]
export const deleteIcPurchasingManagement = (data) => {
  return service({
    url: '/icPurchasingManagement/deleteIcPurchasingManagement',
    method: 'delete',
    data
  })
}

// @Tags IcPurchasingManagement
// @Summary 删除IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icPurchasingManagement/deleteIcPurchasingManagement [delete]
export const deleteIcPurchasingManagementByIds = (data) => {
  return service({
    url: '/icPurchasingManagement/deleteIcPurchasingManagementByIds',
    method: 'delete',
    data
  })
}

// @Tags IcPurchasingManagement
// @Summary 更新IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcPurchasingManagement true "更新IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icPurchasingManagement/updateIcPurchasingManagement [put]
export const updateIcPurchasingManagement = (data) => {
  return service({
    url: '/icPurchasingManagement/updateIcPurchasingManagement',
    method: 'put',
    data
  })
}

// @Tags IcPurchasingManagement
// @Summary 用id查询IcPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IcPurchasingManagement true "用id查询IcPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icPurchasingManagement/findIcPurchasingManagement [get]
export const findIcPurchasingManagement = (params) => {
  return service({
    url: '/icPurchasingManagement/findIcPurchasingManagement',
    method: 'get',
    params
  })
}

// @Tags IcPurchasingManagement
// @Summary 分页获取IcPurchasingManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IcPurchasingManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icPurchasingManagement/getIcPurchasingManagementList [get]
export const getIcPurchasingManagementList = (params) => {
  return service({
    url: '/icPurchasingManagement/getIcPurchasingManagementList',
    method: 'get',
    params
  })
}
