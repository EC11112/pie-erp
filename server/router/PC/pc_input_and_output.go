package PC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PcInputAndOutputRouter struct {
}

// InitPcInputAndOutputRouter 初始化 PcInputAndOutput 路由信息
func (s *PcInputAndOutputRouter) InitPcInputAndOutputRouter(Router *gin.RouterGroup) {
	pcInputAndOutputRouter := Router.Group("pcInputAndOutput").Use(middleware.OperationRecord())
	pcInputAndOutputRouterWithoutRecord := Router.Group("pcInputAndOutput")
	var pcInputAndOutputApi = v1.ApiGroupApp.PCApiGroup.PcInputAndOutputApi
	{
		pcInputAndOutputRouter.POST("createPcInputAndOutput", pcInputAndOutputApi.CreatePcInputAndOutput)   // 新建PcInputAndOutput
		pcInputAndOutputRouter.DELETE("deletePcInputAndOutput", pcInputAndOutputApi.DeletePcInputAndOutput) // 删除PcInputAndOutput
		pcInputAndOutputRouter.DELETE("deletePcInputAndOutputByIds", pcInputAndOutputApi.DeletePcInputAndOutputByIds) // 批量删除PcInputAndOutput
		pcInputAndOutputRouter.PUT("updatePcInputAndOutput", pcInputAndOutputApi.UpdatePcInputAndOutput)    // 更新PcInputAndOutput
	}
	{
		pcInputAndOutputRouterWithoutRecord.GET("findPcInputAndOutput", pcInputAndOutputApi.FindPcInputAndOutput)        // 根据ID获取PcInputAndOutput
		pcInputAndOutputRouterWithoutRecord.GET("getPcInputAndOutputList", pcInputAndOutputApi.GetPcInputAndOutputList)  // 获取PcInputAndOutput列表
	}
}
