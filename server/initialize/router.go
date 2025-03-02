package initialize

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath))

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)

		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)

	}
	{

	}
	{

	}
	{

	}
	{
		HRMRouter := router.RouterGroupApp.HRM
		HRMRouter.InitHrmLeaveOrOvertimeRouter(PrivateGroup)
		HRMRouter.InitHrmSalaryRulesRouter(PrivateGroup)
		HRMRouter.InitHrmFinesOrBonusRouter(PrivateGroup)

	}
	{
		CRMRouter := router.RouterGroupApp.CRM
		CRMRouter.InitCrmCustomerInformationRouter(PrivateGroup)
	}
	{
		PCRouter := router.RouterGroupApp.PC
		PCRouter.InitPcProjectManagementRouter(PrivateGroup)
		PCRouter.InitPcProductionProcessRouter(PrivateGroup)
		PCRouter.InitPcInputAndOutputRouter(PrivateGroup)
	}
	{
		ICRouter := router.RouterGroupApp.IC
		ICRouter.InitIcWarehouseInformationRouter(PrivateGroup)
		ICRouter.InitIcItemInformationRouter(PrivateGroup)
		ICRouter.InitIcInventoryChangesRouter(PrivateGroup)
		ICRouter.InitIcViewInventoryRouter(PrivateGroup)

	}
	{

	}
	{
		FMRouter := router.RouterGroupApp.FM
		FMRouter.InitFmViewOrderRouter(PrivateGroup)
		FMRouter.InitFmPurchasingManagementRouter(PrivateGroup)
		FMRouter.InitFmExpensesAndReceiptsRouter(PrivateGroup)
		FMRouter.InitFmFinancialStatisticsRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
