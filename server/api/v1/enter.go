package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/BM"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/CRM"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	HRMApiGroup     HRM.ApiGroup
	CRMApiGroup     CRM.ApiGroup
	BMApiGroup      BM.ApiGroup
	PCApiGroup      PC.ApiGroup
	ICApiGroup      IC.ApiGroup
	FMApiGroup      FM.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
