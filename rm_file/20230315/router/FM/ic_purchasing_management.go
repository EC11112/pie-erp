package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IcPurchasingManagementRouter struct {
}

// InitIcPurchasingManagementRouter 初始化 IcPurchasingManagement 路由信息
func (s *IcPurchasingManagementRouter) InitIcPurchasingManagementRouter(Router *gin.RouterGroup) {
	icPurchasingManagementRouter := Router.Group("icPurchasingManagement").Use(middleware.OperationRecord())
	icPurchasingManagementRouterWithoutRecord := Router.Group("icPurchasingManagement")
	var icPurchasingManagementApi = v1.ApiGroupApp.FMApiGroup.IcPurchasingManagementApi
	{
		icPurchasingManagementRouter.POST("createIcPurchasingManagement", icPurchasingManagementApi.CreateIcPurchasingManagement)   // 新建IcPurchasingManagement
		icPurchasingManagementRouter.DELETE("deleteIcPurchasingManagement", icPurchasingManagementApi.DeleteIcPurchasingManagement) // 删除IcPurchasingManagement
		icPurchasingManagementRouter.DELETE("deleteIcPurchasingManagementByIds", icPurchasingManagementApi.DeleteIcPurchasingManagementByIds) // 批量删除IcPurchasingManagement
		icPurchasingManagementRouter.PUT("updateIcPurchasingManagement", icPurchasingManagementApi.UpdateIcPurchasingManagement)    // 更新IcPurchasingManagement
	}
	{
		icPurchasingManagementRouterWithoutRecord.GET("findIcPurchasingManagement", icPurchasingManagementApi.FindIcPurchasingManagement)        // 根据ID获取IcPurchasingManagement
		icPurchasingManagementRouterWithoutRecord.GET("getIcPurchasingManagementList", icPurchasingManagementApi.GetIcPurchasingManagementList)  // 获取IcPurchasingManagement列表
	}
}
