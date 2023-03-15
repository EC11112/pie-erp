package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FmFinancialStatisticsRouter struct {
}

// InitFmFinancialStatisticsRouter 初始化 FmFinancialStatistics 路由信息
func (s *FmFinancialStatisticsRouter) InitFmFinancialStatisticsRouter(Router *gin.RouterGroup) {
	fmFinancialStatisticsRouter := Router.Group("fmFinancialStatistics").Use(middleware.OperationRecord())
	fmFinancialStatisticsRouterWithoutRecord := Router.Group("fmFinancialStatistics")
	var fmFinancialStatisticsApi = v1.ApiGroupApp.FMApiGroup.FmFinancialStatisticsApi
	{
		fmFinancialStatisticsRouter.POST("createFmFinancialStatistics", fmFinancialStatisticsApi.CreateFmFinancialStatistics)   // 新建FmFinancialStatistics
		fmFinancialStatisticsRouter.DELETE("deleteFmFinancialStatistics", fmFinancialStatisticsApi.DeleteFmFinancialStatistics) // 删除FmFinancialStatistics
		fmFinancialStatisticsRouter.DELETE("deleteFmFinancialStatisticsByIds", fmFinancialStatisticsApi.DeleteFmFinancialStatisticsByIds) // 批量删除FmFinancialStatistics
		fmFinancialStatisticsRouter.PUT("updateFmFinancialStatistics", fmFinancialStatisticsApi.UpdateFmFinancialStatistics)    // 更新FmFinancialStatistics
	}
	{
		fmFinancialStatisticsRouterWithoutRecord.GET("findFmFinancialStatistics", fmFinancialStatisticsApi.FindFmFinancialStatistics)        // 根据ID获取FmFinancialStatistics
		fmFinancialStatisticsRouterWithoutRecord.GET("getFmFinancialStatisticsList", fmFinancialStatisticsApi.GetFmFinancialStatisticsList)  // 获取FmFinancialStatistics列表
	}
}
