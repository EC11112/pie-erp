package FM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FmExpensesAndReceiptsRouter struct {
}

// InitFmExpensesAndReceiptsRouter 初始化 FmExpensesAndReceipts 路由信息
func (s *FmExpensesAndReceiptsRouter) InitFmExpensesAndReceiptsRouter(Router *gin.RouterGroup) {
	fmExpensesAndReceiptsRouter := Router.Group("fmExpensesAndReceipts").Use(middleware.OperationRecord())
	fmExpensesAndReceiptsRouterWithoutRecord := Router.Group("fmExpensesAndReceipts")
	var fmExpensesAndReceiptsApi = v1.ApiGroupApp.FMApiGroup.FmExpensesAndReceiptsApi
	{
		fmExpensesAndReceiptsRouter.POST("createFmExpensesAndReceipts", fmExpensesAndReceiptsApi.CreateFmExpensesAndReceipts)   // 新建FmExpensesAndReceipts
		fmExpensesAndReceiptsRouter.DELETE("deleteFmExpensesAndReceipts", fmExpensesAndReceiptsApi.DeleteFmExpensesAndReceipts) // 删除FmExpensesAndReceipts
		fmExpensesAndReceiptsRouter.DELETE("deleteFmExpensesAndReceiptsByIds", fmExpensesAndReceiptsApi.DeleteFmExpensesAndReceiptsByIds) // 批量删除FmExpensesAndReceipts
		fmExpensesAndReceiptsRouter.PUT("updateFmExpensesAndReceipts", fmExpensesAndReceiptsApi.UpdateFmExpensesAndReceipts)    // 更新FmExpensesAndReceipts
	}
	{
		fmExpensesAndReceiptsRouterWithoutRecord.GET("findFmExpensesAndReceipts", fmExpensesAndReceiptsApi.FindFmExpensesAndReceipts)        // 根据ID获取FmExpensesAndReceipts
		fmExpensesAndReceiptsRouterWithoutRecord.GET("getFmExpensesAndReceiptsList", fmExpensesAndReceiptsApi.GetFmExpensesAndReceiptsList)  // 获取FmExpensesAndReceipts列表
	}
}
