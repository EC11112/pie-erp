import service from '@/utils/request'

// @Tags FmFinancialStatistics
// @Summary 创建FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmFinancialStatistics true "创建FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmFinancialStatistics/createFmFinancialStatistics [post]
export const createFmFinancialStatistics = (data) => {
  return service({
    url: '/fmFinancialStatistics/createFmFinancialStatistics',
    method: 'post',
    data
  })
}

// @Tags FmFinancialStatistics
// @Summary 删除FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmFinancialStatistics true "删除FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmFinancialStatistics/deleteFmFinancialStatistics [delete]
export const deleteFmFinancialStatistics = (data) => {
  return service({
    url: '/fmFinancialStatistics/deleteFmFinancialStatistics',
    method: 'delete',
    data
  })
}

// @Tags FmFinancialStatistics
// @Summary 删除FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fmFinancialStatistics/deleteFmFinancialStatistics [delete]
export const deleteFmFinancialStatisticsByIds = (data) => {
  return service({
    url: '/fmFinancialStatistics/deleteFmFinancialStatisticsByIds',
    method: 'delete',
    data
  })
}

// @Tags FmFinancialStatistics
// @Summary 更新FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FmFinancialStatistics true "更新FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fmFinancialStatistics/updateFmFinancialStatistics [put]
export const updateFmFinancialStatistics = (data) => {
  return service({
    url: '/fmFinancialStatistics/updateFmFinancialStatistics',
    method: 'put',
    data
  })
}

// @Tags FmFinancialStatistics
// @Summary 用id查询FmFinancialStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FmFinancialStatistics true "用id查询FmFinancialStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fmFinancialStatistics/findFmFinancialStatistics [get]
export const findFmFinancialStatistics = (params) => {
  return service({
    url: '/fmFinancialStatistics/findFmFinancialStatistics',
    method: 'get',
    params
  })
}

// @Tags FmFinancialStatistics
// @Summary 分页获取FmFinancialStatistics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FmFinancialStatistics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fmFinancialStatistics/getFmFinancialStatisticsList [get]
export const getFmFinancialStatisticsList = (params) => {
  return service({
    url: '/fmFinancialStatistics/getFmFinancialStatisticsList',
    method: 'get',
    params
  })
}
