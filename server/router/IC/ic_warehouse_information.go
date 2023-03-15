package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IcWarehouseInformationRouter struct {
}

// InitIcWarehouseInformationRouter 初始化 IcWarehouseInformation 路由信息
func (s *IcWarehouseInformationRouter) InitIcWarehouseInformationRouter(Router *gin.RouterGroup) {
	icWarehouseInformationRouter := Router.Group("icWarehouseInformation").Use(middleware.OperationRecord())
	icWarehouseInformationRouterWithoutRecord := Router.Group("icWarehouseInformation")
	var icWarehouseInformationApi = v1.ApiGroupApp.ICApiGroup.IcWarehouseInformationApi
	{
		icWarehouseInformationRouter.POST("createIcWarehouseInformation", icWarehouseInformationApi.CreateIcWarehouseInformation)   // 新建IcWarehouseInformation
		icWarehouseInformationRouter.DELETE("deleteIcWarehouseInformation", icWarehouseInformationApi.DeleteIcWarehouseInformation) // 删除IcWarehouseInformation
		icWarehouseInformationRouter.DELETE("deleteIcWarehouseInformationByIds", icWarehouseInformationApi.DeleteIcWarehouseInformationByIds) // 批量删除IcWarehouseInformation
		icWarehouseInformationRouter.PUT("updateIcWarehouseInformation", icWarehouseInformationApi.UpdateIcWarehouseInformation)    // 更新IcWarehouseInformation
	}
	{
		icWarehouseInformationRouterWithoutRecord.GET("findIcWarehouseInformation", icWarehouseInformationApi.FindIcWarehouseInformation)        // 根据ID获取IcWarehouseInformation
		icWarehouseInformationRouterWithoutRecord.GET("getIcWarehouseInformationList", icWarehouseInformationApi.GetIcWarehouseInformationList)  // 获取IcWarehouseInformation列表
	}
}
