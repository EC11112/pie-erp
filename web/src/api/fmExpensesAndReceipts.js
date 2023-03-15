import service from '@/utils/request'

// @Tags FmExpensesAndReceipts
// @Summary 创建FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmExpensesAndReceipts true "创建FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmExpensesAndReceipts/createFmExpensesAndReceipts [post]
export const createFmExpensesAndReceipts = (data) => {
  return service({
    url: '/fmExpensesAndReceipts/createFmExpensesAndReceipts',
    method: 'post',
    data
  })
}

// @Tags FmExpensesAndReceipts
// @Summary 删除FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmExpensesAndReceipts true "删除FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmExpensesAndReceipts/deleteFmExpensesAndReceipts [delete]
export const deleteFmExpensesAndReceipts = (data) => {
  return service({
    url: '/fmExpensesAndReceipts/deleteFmExpensesAndReceipts',
    method: 'delete',
    data
  })
}

// @Tags FmExpensesAndReceipts
// @Summary 删除FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmExpensesAndReceipts/deleteFmExpensesAndReceipts [delete]
export const deleteFmExpensesAndReceiptsByIds = (data) => {
  return service({
    url: '/fmExpensesAndReceipts/deleteFmExpensesAndReceiptsByIds',
    method: 'delete',
    data
  })
}

// @Tags FmExpensesAndReceipts
// @Summary 更新FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmExpensesAndReceipts true "更新FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmExpensesAndReceipts/updateFmExpensesAndReceipts [put]
export const updateFmExpensesAndReceipts = (data) => {
  return service({
    url: '/fmExpensesAndReceipts/updateFmExpensesAndReceipts',
    method: 'put',
    data
  })
}

// @Tags FmExpensesAndReceipts
// @Summary 用id查询FmExpensesAndReceipts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FmExpensesAndReceipts true "用id查询FmExpensesAndReceipts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmExpensesAndReceipts/findFmExpensesAndReceipts [get]
export const findFmExpensesAndReceipts = (params) => {
  return service({
    url: '/fmExpensesAndReceipts/findFmExpensesAndReceipts',
    method: 'get',
    params
  })
}

// @Tags FmExpensesAndReceipts
// @Summary 分页获取FmExpensesAndReceipts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FmExpensesAndReceipts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmExpensesAndReceipts/getFmExpensesAndReceiptsList [get]
export const getFmExpensesAndReceiptsList = (params) => {
  return service({
    url: '/fmExpensesAndReceipts/getFmExpensesAndReceiptsList',
    method: 'get',
    params
  })
}
