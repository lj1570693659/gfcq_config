// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductRoles is the golang structure for table product_roles.
type ProductRoles struct {
	Id         int         `json:"id"         description:""`
	Name       string      `json:"name"       description:""`
	Pid        uint        `json:"pid"        description:"上级角色"`
	Explain    string      `json:"explain"    description:"角色与职责说明"`
	IsSpecial  uint        `json:"isSpecial"  description:"1: 需要特殊处理 2：不需要特殊处理"`
	Remark     string      `json:"remark"     description:"预留备注信息"`
	CreateTime *gtime.Time `json:"createTime" description:"新增数据时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
