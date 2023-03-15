import service from '@/utils/request'

// @Tags IcItemInformation
// @Summary 创建IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcItemInformation true "创建IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icItemInformation/createIcItemInformation [post]
export const createIcItemInformation = (data) => {
  return service({
    url: '/icItemInformation/createIcItemInformation',
    method: 'post',
    data
  })
}

// @Tags IcItemInformation
// @Summary 删除IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcItemInformation true "删除IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icItemInformation/deleteIcItemInformation [delete]
export const deleteIcItemInformation = (data) => {
  return service({
    url: '/icItemInformation/deleteIcItemInformation',
    method: 'delete',
    data
  })
}

// @Tags IcItemInformation
// @Summary 删除IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /icItemInformation/deleteIcItemInformation [delete]
export const deleteIcItemInformationByIds = (data) => {
  return service({
    url: '/icItemInformation/deleteIcItemInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags IcItemInformation
// @Summary 更新IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IcItemInformation true "更新IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /icItemInformation/updateIcItemInformation [put]
export const updateIcItemInformation = (data) => {
  return service({
    url: '/icItemInformation/updateIcItemInformation',
    method: 'put',
    data
  })
}

// @Tags IcItemInformation
// @Summary 用id查询IcItemInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IcItemInformation true "用id查询IcItemInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /icItemInformation/findIcItemInformation [get]
export const findIcItemInformation = (params) => {
  return service({
    url: '/icItemInformation/findIcItemInformation',
    method: 'get',
    params
  })
}

// @Tags IcItemInformation
// @Summary 分页获取IcItemInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IcItemInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /icItemInformation/getIcItemInformationList [get]
export const getIcItemInformationList = (params) => {
  return service({
    url: '/icItemInformation/getIcItemInformationList',
    method: 'get',
    params
  })
}
