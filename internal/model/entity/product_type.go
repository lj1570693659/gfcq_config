// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductType is the golang structure for table product_type.
type ProductType struct {
	Id         uint        `json:"id"         description:""`
	Name       string      `json:"name"       description:"开发模型名称"`
	Remark     string      `json:"remark"     description:"预留备注说明信息"`
	CreateTime *gtime.Time `json:"createTime" description:"新增数据时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
