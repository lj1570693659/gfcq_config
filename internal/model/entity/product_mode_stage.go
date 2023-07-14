// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductModeStage is the golang structure for table product_mode_stage.
type ProductModeStage struct {
	Id         uint        `json:"id"         description:""`
	Tid        uint        `json:"tid"        description:"项目开发模式ID"`
	Name       string      `json:"name"       description:"项目阶段名称"`
	QuotaRadio float64     `json:"quotaRadio" description:"阶段额度占比"`
	Remark     string      `json:"remark"     description:"预留备注说明信息"`
	CreateTime *gtime.Time `json:"createTime" description:"新增数据时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
