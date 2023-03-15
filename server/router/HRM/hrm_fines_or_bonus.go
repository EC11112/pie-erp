package HRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HrmFinesOrBonusRouter struct {
}

// InitHrmFinesOrBonusRouter 初始化 HrmFinesOrBonus 路由信息
func (s *HrmFinesOrBonusRouter) InitHrmFinesOrBonusRouter(Router *gin.RouterGroup) {
	hrmFinesOrBonusRouter := Router.Group("hrmFinesOrBonus").Use(middleware.OperationRecord())
	hrmFinesOrBonusRouterWithoutRecord := Router.Group("hrmFinesOrBonus")
	var hrmFinesOrBonusApi = v1.ApiGroupApp.HRMApiGroup.HrmFinesOrBonusApi
	{
		hrmFinesOrBonusRouter.POST("createHrmFinesOrBonus", hrmFinesOrBonusApi.CreateHrmFinesOrBonus)   // 新建HrmFinesOrBonus
		hrmFinesOrBonusRouter.DELETE("deleteHrmFinesOrBonus", hrmFinesOrBonusApi.DeleteHrmFinesOrBonus) // 删除HrmFinesOrBonus
		hrmFinesOrBonusRouter.DELETE("deleteHrmFinesOrBonusByIds", hrmFinesOrBonusApi.DeleteHrmFinesOrBonusByIds) // 批量删除HrmFinesOrBonus
		hrmFinesOrBonusRouter.PUT("updateHrmFinesOrBonus", hrmFinesOrBonusApi.UpdateHrmFinesOrBonus)    // 更新HrmFinesOrBonus
	}
	{
		hrmFinesOrBonusRouterWithoutRecord.GET("findHrmFinesOrBonus", hrmFinesOrBonusApi.FindHrmFinesOrBonus)        // 根据ID获取HrmFinesOrBonus
		hrmFinesOrBonusRouterWithoutRecord.GET("getHrmFinesOrBonusList", hrmFinesOrBonusApi.GetHrmFinesOrBonusList)  // 获取HrmFinesOrBonus列表
	}
}
