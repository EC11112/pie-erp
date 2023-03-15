package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PcProjectManagementRouter struct {
}

// InitPcProjectManagementRouter 初始化 PcProjectManagement 路由信息
func (s *PcProjectManagementRouter) InitPcProjectManagementRouter(Router *gin.RouterGroup) {
	pcProjectManagementRouter := Router.Group("pcProjectManagement").Use(middleware.OperationRecord())
	pcProjectManagementRouterWithoutRecord := Router.Group("pcProjectManagement")
	var pcProjectManagementApi = v1.ApiGroupApp.PCApiGroup.PcProjectManagementApi
	{
		pcProjectManagementRouter.POST("createPcProjectManagement", pcProjectManagementApi.CreatePcProjectManagement)   // 新建PcProjectManagement
		pcProjectManagementRouter.DELETE("deletePcProjectManagement", pcProjectManagementApi.DeletePcProjectManagement) // 删除PcProjectManagement
		pcProjectManagementRouter.DELETE("deletePcProjectManagementByIds", pcProjectManagementApi.DeletePcProjectManagementByIds) // 批量删除PcProjectManagement
		pcProjectManagementRouter.PUT("updatePcProjectManagement", pcProjectManagementApi.UpdatePcProjectManagement)    // 更新PcProjectManagement
	}
	{
		pcProjectManagementRouterWithoutRecord.GET("findPcProjectManagement", pcProjectManagementApi.FindPcProjectManagement)        // 根据ID获取PcProjectManagement
		pcProjectManagementRouterWithoutRecord.GET("getPcProjectManagementList", pcProjectManagementApi.GetPcProjectManagementList)  // 获取PcProjectManagement列表
	}
}
