package IC

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IcItemInformationRouter struct {
}

// InitIcItemInformationRouter 初始化 IcItemInformation 路由信息
func (s *IcItemInformationRouter) InitIcItemInformationRouter(Router *gin.RouterGroup) {
	icItemInformationRouter := Router.Group("icItemInformation").Use(middleware.OperationRecord())
	icItemInformationRouterWithoutRecord := Router.Group("icItemInformation")
	var icItemInformationApi = v1.ApiGroupApp.ICApiGroup.IcItemInformationApi
	{
		icItemInformationRouter.POST("createIcItemInformation", icItemInformationApi.CreateIcItemInformation)   // 新建IcItemInformation
		icItemInformationRouter.DELETE("deleteIcItemInformation", icItemInformationApi.DeleteIcItemInformation) // 删除IcItemInformation
		icItemInformationRouter.DELETE("deleteIcItemInformationByIds", icItemInformationApi.DeleteIcItemInformationByIds) // 批量删除IcItemInformation
		icItemInformationRouter.PUT("updateIcItemInformation", icItemInformationApi.UpdateIcItemInformation)    // 更新IcItemInformation
	}
	{
		icItemInformationRouterWithoutRecord.GET("findIcItemInformation", icItemInformationApi.FindIcItemInformation)        // 根据ID获取IcItemInformation
		icItemInformationRouterWithoutRecord.GET("getIcItemInformationList", icItemInformationApi.GetIcItemInformationList)  // 获取IcItemInformation列表
	}
}
