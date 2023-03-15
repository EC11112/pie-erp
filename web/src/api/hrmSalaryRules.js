import service from '@/utils/request'

// @Tags HrmSalaryRules
// @Summary 创建HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmSalaryRules true "创建HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmSalaryRules/createHrmSalaryRules [post]
export const createHrmSalaryRules = (data) => {
  return service({
    url: '/hrmSalaryRules/createHrmSalaryRules',
    method: 'post',
    data
  })
}

// @Tags HrmSalaryRules
// @Summary 删除HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmSalaryRules true "删除HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmSalaryRules/deleteHrmSalaryRules [delete]
export const deleteHrmSalaryRules = (data) => {
  return service({
    url: '/hrmSalaryRules/deleteHrmSalaryRules',
    method: 'delete',
    data
  })
}

// @Tags HrmSalaryRules
// @Summary 删除HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hrmSalaryRules/deleteHrmSalaryRules [delete]
export const deleteHrmSalaryRulesByIds = (data) => {
  return service({
    url: '/hrmSalaryRules/deleteHrmSalaryRulesByIds',
    method: 'delete',
    data
  })
}

// @Tags HrmSalaryRules
// @Summary 更新HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HrmSalaryRules true "更新HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hrmSalaryRules/updateHrmSalaryRules [put]
export const updateHrmSalaryRules = (data) => {
  return service({
    url: '/hrmSalaryRules/updateHrmSalaryRules',
    method: 'put',
    data
  })
}

// @Tags HrmSalaryRules
// @Summary 用id查询HrmSalaryRules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HrmSalaryRules true "用id查询HrmSalaryRules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hrmSalaryRules/findHrmSalaryRules [get]
export const findHrmSalaryRules = (params) => {
  return service({
    url: '/hrmSalaryRules/findHrmSalaryRules',
    method: 'get',
    params
  })
}

// @Tags HrmSalaryRules
// @Summary 分页获取HrmSalaryRules列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HrmSalaryRules列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hrmSalaryRules/getHrmSalaryRulesList [get]
export const getHrmSalaryRulesList = (params) => {
  return service({
    url: '/hrmSalaryRules/getHrmSalaryRulesList',
    method: 'get',
    params
  })
}
