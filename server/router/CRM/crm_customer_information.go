package CRM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmCustomerInformationRouter struct {
}

// InitCrmCustomerInformationRouter 初始化 CrmCustomerInformation 路由信息
func (s *CrmCustomerInformationRouter) InitCrmCustomerInformationRouter(Router *gin.RouterGroup) {
	crmCustomerInformationRouter := Router.Group("crmCustomerInformation").Use(middleware.OperationRecord())
	crmCustomerInformationRouterWithoutRecord := Router.Group("crmCustomerInformation")
	var crmCustomerInformationApi = v1.ApiGroupApp.CRMApiGroup.CrmCustomerInformationApi
	{
		crmCustomerInformationRouter.POST("createCrmCustomerInformation", crmCustomerInformationApi.CreateCrmCustomerInformation)   // 新建CrmCustomerInformation
		crmCustomerInformationRouter.DELETE("deleteCrmCustomerInformation", crmCustomerInformationApi.DeleteCrmCustomerInformation) // 删除CrmCustomerInformation
		crmCustomerInformationRouter.DELETE("deleteCrmCustomerInformationByIds", crmCustomerInformationApi.DeleteCrmCustomerInformationByIds) // 批量删除CrmCustomerInformation
		crmCustomerInformationRouter.PUT("updateCrmCustomerInformation", crmCustomerInformationApi.UpdateCrmCustomerInformation)    // 更新CrmCustomerInformation
	}
	{
		crmCustomerInformationRouterWithoutRecord.GET("findCrmCustomerInformation", crmCustomerInformationApi.FindCrmCustomerInformation)        // 根据ID获取CrmCustomerInformation
		crmCustomerInformationRouterWithoutRecord.GET("getCrmCustomerInformationList", crmCustomerInformationApi.GetCrmCustomerInformationList)  // 获取CrmCustomerInformation列表
	}
}
