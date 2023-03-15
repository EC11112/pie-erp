package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PcProductionProcessRouter struct {
}

// InitPcProductionProcessRouter 初始化 PcProductionProcess 路由信息
func (s *PcProductionProcessRouter) InitPcProductionProcessRouter(Router *gin.RouterGroup) {
	pcProductionProcessRouter := Router.Group("pcProductionProcess").Use(middleware.OperationRecord())
	pcProductionProcessRouterWithoutRecord := Router.Group("pcProductionProcess")
	var pcProductionProcessApi = v1.ApiGroupApp.PCApiGroup.PcProductionProcessApi
	{
		pcProductionProcessRouter.POST("createPcProductionProcess", pcProductionProcessApi.CreatePcProductionProcess)   // 新建PcProductionProcess
		pcProductionProcessRouter.DELETE("deletePcProductionProcess", pcProductionProcessApi.DeletePcProductionProcess) // 删除PcProductionProcess
		pcProductionProcessRouter.DELETE("deletePcProductionProcessByIds", pcProductionProcessApi.DeletePcProductionProcessByIds) // 批量删除PcProductionProcess
		pcProductionProcessRouter.PUT("updatePcProductionProcess", pcProductionProcessApi.UpdatePcProductionProcess)    // 更新PcProductionProcess
	}
	{
		pcProductionProcessRouterWithoutRecord.GET("findPcProductionProcess", pcProductionProcessApi.FindPcProductionProcess)        // 根据ID获取PcProductionProcess
		pcProductionProcessRouterWithoutRecord.GET("getPcProductionProcessList", pcProductionProcessApi.GetPcProductionProcessList)  // 获取PcProductionProcess列表
	}
}
