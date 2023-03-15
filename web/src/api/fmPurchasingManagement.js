import service from '@/utils/request'

// @Tags FmPurchasingManagement
// @Summary 创建FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmPurchasingManagement true "创建FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmPurchasingManagement/createFmPurchasingManagement [post]
export const createFmPurchasingManagement = (data) => {
  return service({
    url: '/fmPurchasingManagement/createFmPurchasingManagement',
    method: 'post',
    data
  })
}

// @Tags FmPurchasingManagement
// @Summary 删除FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmPurchasingManagement true "删除FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmPurchasingManagement/deleteFmPurchasingManagement [delete]
export const deleteFmPurchasingManagement = (data) => {
  return service({
    url: '/fmPurchasingManagement/deleteFmPurchasingManagement',
    method: 'delete',
    data
  })
}

// @Tags FmPurchasingManagement
// @Summary 删除FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmPurchasingManagement/deleteFmPurchasingManagement [delete]
export const deleteFmPurchasingManagementByIds = (data) => {
  return service({
    url: '/fmPurchasingManagement/deleteFmPurchasingManagementByIds',
    method: 'delete',
    data
  })
}

// @Tags FmPurchasingManagement
// @Summary 更新FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmPurchasingManagement true "更新FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmPurchasingManagement/updateFmPurchasingManagement [put]
export const updateFmPurchasingManagement = (data) => {
  return service({
    url: '/fmPurchasingManagement/updateFmPurchasingManagement',
    method: 'put',
    data
  })
}

// @Tags FmPurchasingManagement
// @Summary 用id查询FmPurchasingManagement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FmPurchasingManagement true "用id查询FmPurchasingManagement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmPurchasingManagement/findFmPurchasingManagement [get]
export const findFmPurchasingManagement = (params) => {
  return service({
    url: '/fmPurchasingManagement/findFmPurchasingManagement',
    method: 'get',
    params
  })
}

// @Tags FmPurchasingManagement
// @Summary 分页获取FmPurchasingManagement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FmPurchasingManagement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmPurchasingManagement/getFmPurchasingManagementList [get]
export const getFmPurchasingManagementList = (params) => {
  return service({
    url: '/fmPurchasingManagement/getFmPurchasingManagementList',
    method: 'get',
    params
  })
}
