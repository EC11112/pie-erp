package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/HRM"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HrmLeaveOrOvertimeSearch struct{
    HRM.HrmLeaveOrOvertime
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartStart_date  *time.Time  `json:"startStart_date" form:"startStart_date"`
    EndStart_date  *time.Time  `json:"endStart_date" form:"endStart_date"`
    StartEnd_date  *time.Time  `json:"startEnd_date" form:"startEnd_date"`
    EndEnd_date  *time.Time  `json:"endEnd_date" form:"endEnd_date"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
