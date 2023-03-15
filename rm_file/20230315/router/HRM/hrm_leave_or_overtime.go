package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HrmLeaveOrOvertimeRouter struct {
}

// InitHrmLeaveOrOvertimeRouter 初始化 HrmLeaveOrOvertime 路由信息
func (s *HrmLeaveOrOvertimeRouter) InitHrmLeaveOrOvertimeRouter(Router *gin.RouterGroup) {
	hrmLeaveOrOvertimeRouter := Router.Group("hrmLeaveOrOvertime").Use(middleware.OperationRecord())
	hrmLeaveOrOvertimeRouterWithoutRecord := Router.Group("hrmLeaveOrOvertime")
	var hrmLeaveOrOvertimeApi = v1.ApiGroupApp.HRMApiGroup.HrmLeaveOrOvertimeApi
	{
		hrmLeaveOrOvertimeRouter.POST("createHrmLeaveOrOvertime", hrmLeaveOrOvertimeApi.CreateHrmLeaveOrOvertime)   // 新建HrmLeaveOrOvertime
		hrmLeaveOrOvertimeRouter.DELETE("deleteHrmLeaveOrOvertime", hrmLeaveOrOvertimeApi.DeleteHrmLeaveOrOvertime) // 删除HrmLeaveOrOvertime
		hrmLeaveOrOvertimeRouter.DELETE("deleteHrmLeaveOrOvertimeByIds", hrmLeaveOrOvertimeApi.DeleteHrmLeaveOrOvertimeByIds) // 批量删除HrmLeaveOrOvertime
		hrmLeaveOrOvertimeRouter.PUT("updateHrmLeaveOrOvertime", hrmLeaveOrOvertimeApi.UpdateHrmLeaveOrOvertime)    // 更新HrmLeaveOrOvertime
	}
	{
		hrmLeaveOrOvertimeRouterWithoutRecord.GET("findHrmLeaveOrOvertime", hrmLeaveOrOvertimeApi.FindHrmLeaveOrOvertime)        // 根据ID获取HrmLeaveOrOvertime
		hrmLeaveOrOvertimeRouterWithoutRecord.GET("getHrmLeaveOrOvertimeList", hrmLeaveOrOvertimeApi.GetHrmLeaveOrOvertimeList)  // 获取HrmLeaveOrOvertime列表
	}
}
