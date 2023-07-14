// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewManageIndex is the golang structure for table crew_manage_index.
type CrewManageIndex struct {
	Id            uint        `json:"id"            description:""`
	ScoreIndex    uint        `json:"scoreIndex"    description:"管理指数"`
	ProductRoleId uint        `json:"productRoleId" description:"项目角色ID"`
	Remark        string      `json:"remark"        description:"预留备注说明信息"`
	CreateTime    *gtime.Time `json:"createTime"    description:"新增数据时间"`
	UpdateTime    *gtime.Time `json:"updateTime"    description:"最后一次更新数据时间"`
}
