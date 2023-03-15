import service from '@/utils/request'

// @Tags FmViewOrder
// @Summary 创建FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmViewOrder true "创建FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmViewOrder/createFmViewOrder [post]
export const createFmViewOrder = (data) => {
  return service({
    url: '/fmViewOrder/createFmViewOrder',
    method: 'post',
    data
  })
}

// @Tags FmViewOrder
// @Summary 删除FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmViewOrder true "删除FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmViewOrder/deleteFmViewOrder [delete]
export const deleteFmViewOrder = (data) => {
  return service({
    url: '/fmViewOrder/deleteFmViewOrder',
    method: 'delete',
    data
  })
}

// @Tags FmViewOrder
// @Summary 删除FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmViewOrder/deleteFmViewOrder [delete]
export const deleteFmViewOrderByIds = (data) => {
  return service({
    url: '/fmViewOrder/deleteFmViewOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags FmViewOrder
// @Summary 更新FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmViewOrder true "更新FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmViewOrder/updateFmViewOrder [put]
export const updateFmViewOrder = (data) => {
  return service({
    url: '/fmViewOrder/updateFmViewOrder',
    method: 'put',
    data
  })
}

// @Tags FmViewOrder
// @Summary 用id查询FmViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FmViewOrder true "用id查询FmViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmViewOrder/findFmViewOrder [get]
export const findFmViewOrder = (params) => {
  return service({
    url: '/fmViewOrder/findFmViewOrder',
    method: 'get',
    params
  })
}

// @Tags FmViewOrder
// @Summary 分页获取FmViewOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FmViewOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmViewOrder/getFmViewOrderList [get]
export const getFmViewOrderList = (params) => {
  return service({
    url: '/fmViewOrder/getFmViewOrderList',
    method: 'get',
    params
  })
}
