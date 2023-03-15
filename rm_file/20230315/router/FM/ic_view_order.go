package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IcViewOrderRouter struct {
}

// InitIcViewOrderRouter 初始化 IcViewOrder 路由信息
func (s *IcViewOrderRouter) InitIcViewOrderRouter(Router *gin.RouterGroup) {
	icViewOrderRouter := Router.Group("icViewOrder").Use(middleware.OperationRecord())
	icViewOrderRouterWithoutRecord := Router.Group("icViewOrder")
	var icViewOrderApi = v1.ApiGroupApp.FMApiGroup.IcViewOrderApi
	{
		icViewOrderRouter.POST("createIcViewOrder", icViewOrderApi.CreateIcViewOrder)   // 新建IcViewOrder
		icViewOrderRouter.DELETE("deleteIcViewOrder", icViewOrderApi.DeleteIcViewOrder) // 删除IcViewOrder
		icViewOrderRouter.DELETE("deleteIcViewOrderByIds", icViewOrderApi.DeleteIcViewOrderByIds) // 批量删除IcViewOrder
		icViewOrderRouter.PUT("updateIcViewOrder", icViewOrderApi.UpdateIcViewOrder)    // 更新IcViewOrder
	}
	{
		icViewOrderRouterWithoutRecord.GET("findIcViewOrder", icViewOrderApi.FindIcViewOrder)        // 根据ID获取IcViewOrder
		icViewOrderRouterWithoutRecord.GET("getIcViewOrderList", icViewOrderApi.GetIcViewOrderList)  // 获取IcViewOrder列表
	}
}
