import service from '@/utils/request'

// @Tags HrmFinesOrBonus
// @Summary 创建HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmFinesOrBonus true "创建HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmFinesOrBonus/createHrmFinesOrBonus [post]
export const createHrmFinesOrBonus = (data) => {
  return service({
    url: '/hrmFinesOrBonus/createHrmFinesOrBonus',
    method: 'post',
    data
  })
}

// @Tags HrmFinesOrBonus
// @Summary 删除HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmFinesOrBonus true "删除HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmFinesOrBonus/deleteHrmFinesOrBonus [delete]
export const deleteHrmFinesOrBonus = (data) => {
  return service({
    url: '/hrmFinesOrBonus/deleteHrmFinesOrBonus',
    method: 'delete',
    data
  })
}

// @Tags HrmFinesOrBonus
// @Summary 删除HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmFinesOrBonus/deleteHrmFinesOrBonus [delete]
export const deleteHrmFinesOrBonusByIds = (data) => {
  return service({
    url: '/hrmFinesOrBonus/deleteHrmFinesOrBonusByIds',
    method: 'delete',
    data
  })
}

// @Tags HrmFinesOrBonus
// @Summary 更新HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmFinesOrBonus true "更新HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrmFinesOrBonus/updateHrmFinesOrBonus [put]
export const updateHrmFinesOrBonus = (data) => {
  return service({
    url: '/hrmFinesOrBonus/updateHrmFinesOrBonus',
    method: 'put',
    data
  })
}

// @Tags HrmFinesOrBonus
// @Summary 用id查询HrmFinesOrBonus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HrmFinesOrBonus true "用id查询HrmFinesOrBonus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrmFinesOrBonus/findHrmFinesOrBonus [get]
export const findHrmFinesOrBonus = (params) => {
  return service({
    url: '/hrmFinesOrBonus/findHrmFinesOrBonus',
    method: 'get',
    params
  })
}

// @Tags HrmFinesOrBonus
// @Summary 分页获取HrmFinesOrBonus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HrmFinesOrBonus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmFinesOrBonus/getHrmFinesOrBonusList [get]
export const getHrmFinesOrBonusList = (params) => {
  return service({
    url: '/hrmFinesOrBonus/getHrmFinesOrBonusList',
    method: 'get',
    params
  })
}
