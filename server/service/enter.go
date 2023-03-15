package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/BM"
	"github.com/flipped-aurora/gin-vue-admin/server/service/CRM"
	"github.com/flipped-aurora/gin-vue-admin/server/service/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/service/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/service/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/service/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	HRMServiceGroup     HRM.ServiceGroup
	CRMServiceGroup     CRM.ServiceGroup
	BMServiceGroup      BM.ServiceGroup
	PCServiceGroup      PC.ServiceGroup
	ICServiceGroup      IC.ServiceGroup
	FMServiceGroup      FM.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
