package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FmPurchasingManagementRouter struct {
}

// InitFmPurchasingManagementRouter 初始化 FmPurchasingManagement 路由信息
func (s *FmPurchasingManagementRouter) InitFmPurchasingManagementRouter(Router *gin.RouterGroup) {
	fmPurchasingManagementRouter := Router.Group("fmPurchasingManagement").Use(middleware.OperationRecord())
	fmPurchasingManagementRouterWithoutRecord := Router.Group("fmPurchasingManagement")
	var fmPurchasingManagementApi = v1.ApiGroupApp.FMApiGroup.FmPurchasingManagementApi
	{
		fmPurchasingManagementRouter.POST("createFmPurchasingManagement", fmPurchasingManagementApi.CreateFmPurchasingManagement)   // 新建FmPurchasingManagement
		fmPurchasingManagementRouter.DELETE("deleteFmPurchasingManagement", fmPurchasingManagementApi.DeleteFmPurchasingManagement) // 删除FmPurchasingManagement
		fmPurchasingManagementRouter.DELETE("deleteFmPurchasingManagementByIds", fmPurchasingManagementApi.DeleteFmPurchasingManagementByIds) // 批量删除FmPurchasingManagement
		fmPurchasingManagementRouter.PUT("updateFmPurchasingManagement", fmPurchasingManagementApi.UpdateFmPurchasingManagement)    // 更新FmPurchasingManagement
	}
	{
		fmPurchasingManagementRouterWithoutRecord.GET("findFmPurchasingManagement", fmPurchasingManagementApi.FindFmPurchasingManagement)        // 根据ID获取FmPurchasingManagement
		fmPurchasingManagementRouterWithoutRecord.GET("getFmPurchasingManagementList", fmPurchasingManagementApi.GetFmPurchasingManagementList)  // 获取FmPurchasingManagement列表
	}
}
