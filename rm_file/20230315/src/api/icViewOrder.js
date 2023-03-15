import service from '@/utils/request'

// @Tags IcViewOrder
// @Summary 创建IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcViewOrder true "创建IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewOrder/createIcViewOrder [post]
export const createIcViewOrder = (data) => {
  return service({
    url: '/icViewOrder/createIcViewOrder',
    method: 'post',
    data
  })
}

// @Tags IcViewOrder
// @Summary 删除IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcViewOrder true "删除IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icViewOrder/deleteIcViewOrder [delete]
export const deleteIcViewOrder = (data) => {
  return service({
    url: '/icViewOrder/deleteIcViewOrder',
    method: 'delete',
    data
  })
}

// @Tags IcViewOrder
// @Summary 删除IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icViewOrder/deleteIcViewOrder [delete]
export const deleteIcViewOrderByIds = (data) => {
  return service({
    url: '/icViewOrder/deleteIcViewOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags IcViewOrder
// @Summary 更新IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcViewOrder true "更新IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icViewOrder/updateIcViewOrder [put]
export const updateIcViewOrder = (data) => {
  return service({
    url: '/icViewOrder/updateIcViewOrder',
    method: 'put',
    data
  })
}

// @Tags IcViewOrder
// @Summary 用id查询IcViewOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IcViewOrder true "用id查询IcViewOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icViewOrder/findIcViewOrder [get]
export const findIcViewOrder = (params) => {
  return service({
    url: '/icViewOrder/findIcViewOrder',
    method: 'get',
    params
  })
}

// @Tags IcViewOrder
// @Summary 分页获取IcViewOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IcViewOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewOrder/getIcViewOrderList [get]
export const getIcViewOrderList = (params) => {
  return service({
    url: '/icViewOrder/getIcViewOrderList',
    method: 'get',
    params
  })
}
