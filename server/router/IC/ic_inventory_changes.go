package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IcInventoryChangesRouter struct {
}

// InitIcInventoryChangesRouter 初始化 IcInventoryChanges 路由信息
func (s *IcInventoryChangesRouter) InitIcInventoryChangesRouter(Router *gin.RouterGroup) {
	icInventoryChangesRouter := Router.Group("icInventoryChanges").Use(middleware.OperationRecord())
	icInventoryChangesRouterWithoutRecord := Router.Group("icInventoryChanges")
	var icInventoryChangesApi = v1.ApiGroupApp.ICApiGroup.IcInventoryChangesApi
	{
		icInventoryChangesRouter.POST("createIcInventoryChanges", icInventoryChangesApi.CreateIcInventoryChanges)   // 新建IcInventoryChanges
		icInventoryChangesRouter.DELETE("deleteIcInventoryChanges", icInventoryChangesApi.DeleteIcInventoryChanges) // 删除IcInventoryChanges
		icInventoryChangesRouter.DELETE("deleteIcInventoryChangesByIds", icInventoryChangesApi.DeleteIcInventoryChangesByIds) // 批量删除IcInventoryChanges
		icInventoryChangesRouter.PUT("updateIcInventoryChanges", icInventoryChangesApi.UpdateIcInventoryChanges)    // 更新IcInventoryChanges
	}
	{
		icInventoryChangesRouterWithoutRecord.GET("findIcInventoryChanges", icInventoryChangesApi.FindIcInventoryChanges)        // 根据ID获取IcInventoryChanges
		icInventoryChangesRouterWithoutRecord.GET("getIcInventoryChangesList", icInventoryChangesApi.GetIcInventoryChangesList)  // 获取IcInventoryChanges列表
	}
}
