package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FmViewOrderRouter struct {
}

// InitFmViewOrderRouter 初始化 FmViewOrder 路由信息
func (s *FmViewOrderRouter) InitFmViewOrderRouter(Router *gin.RouterGroup) {
	fmViewOrderRouter := Router.Group("fmViewOrder").Use(middleware.OperationRecord())
	fmViewOrderRouterWithoutRecord := Router.Group("fmViewOrder")
	var fmViewOrderApi = v1.ApiGroupApp.FMApiGroup.FmViewOrderApi
	{
		fmViewOrderRouter.POST("createFmViewOrder", fmViewOrderApi.CreateFmViewOrder)   // 新建FmViewOrder
		fmViewOrderRouter.DELETE("deleteFmViewOrder", fmViewOrderApi.DeleteFmViewOrder) // 删除FmViewOrder
		fmViewOrderRouter.DELETE("deleteFmViewOrderByIds", fmViewOrderApi.DeleteFmViewOrderByIds) // 批量删除FmViewOrder
		fmViewOrderRouter.PUT("updateFmViewOrder", fmViewOrderApi.UpdateFmViewOrder)    // 更新FmViewOrder
	}
	{
		fmViewOrderRouterWithoutRecord.GET("findFmViewOrder", fmViewOrderApi.FindFmViewOrder)        // 根据ID获取FmViewOrder
		fmViewOrderRouterWithoutRecord.GET("getFmViewOrderList", fmViewOrderApi.GetFmViewOrderList)  // 获取FmViewOrder列表
	}
}
