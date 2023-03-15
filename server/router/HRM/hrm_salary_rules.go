package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HrmSalaryRulesRouter struct {
}

// InitHrmSalaryRulesRouter 初始化 HrmSalaryRules 路由信息
func (s *HrmSalaryRulesRouter) InitHrmSalaryRulesRouter(Router *gin.RouterGroup) {
	hrmSalaryRulesRouter := Router.Group("hrmSalaryRules").Use(middleware.OperationRecord())
	hrmSalaryRulesRouterWithoutRecord := Router.Group("hrmSalaryRules")
	var hrmSalaryRulesApi = v1.ApiGroupApp.HRMApiGroup.HrmSalaryRulesApi
	{
		hrmSalaryRulesRouter.POST("createHrmSalaryRules", hrmSalaryRulesApi.CreateHrmSalaryRules)   // 新建HrmSalaryRules
		hrmSalaryRulesRouter.DELETE("deleteHrmSalaryRules", hrmSalaryRulesApi.DeleteHrmSalaryRules) // 删除HrmSalaryRules
		hrmSalaryRulesRouter.DELETE("deleteHrmSalaryRulesByIds", hrmSalaryRulesApi.DeleteHrmSalaryRulesByIds) // 批量删除HrmSalaryRules
		hrmSalaryRulesRouter.PUT("updateHrmSalaryRules", hrmSalaryRulesApi.UpdateHrmSalaryRules)    // 更新HrmSalaryRules
	}
	{
		hrmSalaryRulesRouterWithoutRecord.GET("findHrmSalaryRules", hrmSalaryRulesApi.FindHrmSalaryRules)        // 根据ID获取HrmSalaryRules
		hrmSalaryRulesRouterWithoutRecord.GET("getHrmSalaryRulesList", hrmSalaryRulesApi.GetHrmSalaryRulesList)  // 获取HrmSalaryRules列表
	}
}
