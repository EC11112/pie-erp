package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HrmSalaryRulesSearch struct{
    HRM.HrmSalaryRules
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
