import service from '@/utils/request'

// @Tags CrmCustomerInformation
// @Summary 创建CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomerInformation true "创建CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerInformation/createCrmCustomerInformation [post]
export const createCrmCustomerInformation = (data) => {
  return service({
    url: '/crmCustomerInformation/createCrmCustomerInformation',
    method: 'post',
    data
  })
}

// @Tags CrmCustomerInformation
// @Summary 删除CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomerInformation true "删除CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomerInformation/deleteCrmCustomerInformation [delete]
export const deleteCrmCustomerInformation = (data) => {
  return service({
    url: '/crmCustomerInformation/deleteCrmCustomerInformation',
    method: 'delete',
    data
  })
}

// @Tags CrmCustomerInformation
// @Summary 删除CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomerInformation/deleteCrmCustomerInformation [delete]
export const deleteCrmCustomerInformationByIds = (data) => {
  return service({
    url: '/crmCustomerInformation/deleteCrmCustomerInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags CrmCustomerInformation
// @Summary 更新CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomerInformation true "更新CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCustomerInformation/updateCrmCustomerInformation [put]
export const updateCrmCustomerInformation = (data) => {
  return service({
    url: '/crmCustomerInformation/updateCrmCustomerInformation',
    method: 'put',
    data
  })
}

// @Tags CrmCustomerInformation
// @Summary 用id查询CrmCustomerInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmCustomerInformation true "用id查询CrmCustomerInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomerInformation/findCrmCustomerInformation [get]
export const findCrmCustomerInformation = (params) => {
  return service({
    url: '/crmCustomerInformation/findCrmCustomerInformation',
    method: 'get',
    params
  })
}

// @Tags CrmCustomerInformation
// @Summary 分页获取CrmCustomerInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CrmCustomerInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerInformation/getCrmCustomerInformationList [get]
export const getCrmCustomerInformationList = (params) => {
  return service({
    url: '/crmCustomerInformation/getCrmCustomerInformationList',
    method: 'get',
    params
  })
}
