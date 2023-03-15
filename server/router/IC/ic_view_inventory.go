package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IcViewInventoryRouter struct {
}

// InitIcViewInventoryRouter 初始化 IcViewInventory 路由信息
func (s *IcViewInventoryRouter) InitIcViewInventoryRouter(Router *gin.RouterGroup) {
	icViewInventoryRouter := Router.Group("icViewInventory").Use(middleware.OperationRecord())
	icViewInventoryRouterWithoutRecord := Router.Group("icViewInventory")
	var icViewInventoryApi = v1.ApiGroupApp.ICApiGroup.IcViewInventoryApi
	{
		icViewInventoryRouter.POST("createIcViewInventory", icViewInventoryApi.CreateIcViewInventory)   // 新建IcViewInventory
		icViewInventoryRouter.DELETE("deleteIcViewInventory", icViewInventoryApi.DeleteIcViewInventory) // 删除IcViewInventory
		icViewInventoryRouter.DELETE("deleteIcViewInventoryByIds", icViewInventoryApi.DeleteIcViewInventoryByIds) // 批量删除IcViewInventory
		icViewInventoryRouter.PUT("updateIcViewInventory", icViewInventoryApi.UpdateIcViewInventory)    // 更新IcViewInventory
	}
	{
		icViewInventoryRouterWithoutRecord.GET("findIcViewInventory", icViewInventoryApi.FindIcViewInventory)        // 根据ID获取IcViewInventory
		icViewInventoryRouterWithoutRecord.GET("getIcViewInventoryList", icViewInventoryApi.GetIcViewInventoryList)  // 获取IcViewInventory列表
	}
}
