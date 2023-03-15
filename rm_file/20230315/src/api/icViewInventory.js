import service from '@/utils/request'

// @Tags IcViewInventory
// @Summary 创建IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcViewInventory true "创建IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewInventory/createIcViewInventory [post]
export const createIcViewInventory = (data) => {
  return service({
    url: '/icViewInventory/createIcViewInventory',
    method: 'post',
    data
  })
}

// @Tags IcViewInventory
// @Summary 删除IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcViewInventory true "删除IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icViewInventory/deleteIcViewInventory [delete]
export const deleteIcViewInventory = (data) => {
  return service({
    url: '/icViewInventory/deleteIcViewInventory',
    method: 'delete',
    data
  })
}

// @Tags IcViewInventory
// @Summary 删除IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icViewInventory/deleteIcViewInventory [delete]
export const deleteIcViewInventoryByIds = (data) => {
  return service({
    url: '/icViewInventory/deleteIcViewInventoryByIds',
    method: 'delete',
    data
  })
}

// @Tags IcViewInventory
// @Summary 更新IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcViewInventory true "更新IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icViewInventory/updateIcViewInventory [put]
export const updateIcViewInventory = (data) => {
  return service({
    url: '/icViewInventory/updateIcViewInventory',
    method: 'put',
    data
  })
}

// @Tags IcViewInventory
// @Summary 用id查询IcViewInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IcViewInventory true "用id查询IcViewInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icViewInventory/findIcViewInventory [get]
export const findIcViewInventory = (params) => {
  return service({
    url: '/icViewInventory/findIcViewInventory',
    method: 'get',
    params
  })
}

// @Tags IcViewInventory
// @Summary 分页获取IcViewInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IcViewInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icViewInventory/getIcViewInventoryList [get]
export const getIcViewInventoryList = (params) => {
  return service({
    url: '/icViewInventory/getIcViewInventoryList',
    method: 'get',
    params
  })
}
