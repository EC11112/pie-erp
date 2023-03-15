import service from '@/utils/request'

// @Tags HrmLeaveOrOvertime
// @Summary 创建HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmLeaveOrOvertime true "创建HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmLeaveOrOvertime/createHrmLeaveOrOvertime [post]
export const createHrmLeaveOrOvertime = (data) => {
  return service({
    url: '/hrmLeaveOrOvertime/createHrmLeaveOrOvertime',
    method: 'post',
    data
  })
}

// @Tags HrmLeaveOrOvertime
// @Summary 删除HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmLeaveOrOvertime true "删除HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmLeaveOrOvertime/deleteHrmLeaveOrOvertime [delete]
export const deleteHrmLeaveOrOvertime = (data) => {
  return service({
    url: '/hrmLeaveOrOvertime/deleteHrmLeaveOrOvertime',
    method: 'delete',
    data
  })
}

// @Tags HrmLeaveOrOvertime
// @Summary 删除HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmLeaveOrOvertime/deleteHrmLeaveOrOvertime [delete]
export const deleteHrmLeaveOrOvertimeByIds = (data) => {
  return service({
    url: '/hrmLeaveOrOvertime/deleteHrmLeaveOrOvertimeByIds',
    method: 'delete',
    data
  })
}

// @Tags HrmLeaveOrOvertime
// @Summary 更新HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmLeaveOrOvertime true "更新HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrmLeaveOrOvertime/updateHrmLeaveOrOvertime [put]
export const updateHrmLeaveOrOvertime = (data) => {
  return service({
    url: '/hrmLeaveOrOvertime/updateHrmLeaveOrOvertime',
    method: 'put',
    data
  })
}

// @Tags HrmLeaveOrOvertime
// @Summary 用id查询HrmLeaveOrOvertime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HrmLeaveOrOvertime true "用id查询HrmLeaveOrOvertime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrmLeaveOrOvertime/findHrmLeaveOrOvertime [get]
export const findHrmLeaveOrOvertime = (params) => {
  return service({
    url: '/hrmLeaveOrOvertime/findHrmLeaveOrOvertime',
    method: 'get',
    params
  })
}

// @Tags HrmLeaveOrOvertime
// @Summary 分页获取HrmLeaveOrOvertime列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HrmLeaveOrOvertime列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmLeaveOrOvertime/getHrmLeaveOrOvertimeList [get]
export const getHrmLeaveOrOvertimeList = (params) => {
  return service({
    url: '/hrmLeaveOrOvertime/getHrmLeaveOrOvertimeList',
    method: 'get',
    params
  })
}
