package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/BM"
	"github.com/flipped-aurora/gin-vue-admin/server/router/CRM"
	"github.com/flipped-aurora/gin-vue-admin/server/router/FM"
	"github.com/flipped-aurora/gin-vue-admin/server/router/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/router/IC"
	"github.com/flipped-aurora/gin-vue-admin/server/router/PC"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	HRM     HRM.RouterGroup
	CRM     CRM.RouterGroup
	BM      BM.RouterGroup
	PC      PC.RouterGroup
	IC      IC.RouterGroup
	FM      FM.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
