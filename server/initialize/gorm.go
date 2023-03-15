package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/CRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/model/FM"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, HRM.HrmLeaveOrOvertime{}, CRM.CrmCustomerInformation{}, PC.PcProjectManagement{}, PC.PcProductionProcess{}, IC.IcWarehouseInformation{}, IC.IcItemInformation{}, IC.IcInventoryChanges{}, HRM.HrmSalaryRules{}, HRM.HrmFinesOrBonus{}, PC.PcInputAndOutput{}, IC.IcViewInventory{}, FM.FmViewOrder{}, FM.FmPurchasingManagement{}, FM.FmExpensesAndReceipts{}, FM.FmFinancialStatistics{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
