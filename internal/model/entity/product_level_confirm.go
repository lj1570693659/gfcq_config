// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductLevelConfirm is the golang structure for table product_level_confirm.
type ProductLevelConfirm struct {
	Id            int         `json:"id"            description:""`
	Name          string      `json:"name"          description:"项目优先级"`
	ScoreMin      float64     `json:"scoreMin"      description:"得分下限"`
	ScoreMax      float64     `json:"scoreMax"      description:"得分上线"`
	ScoreRange    uint        `json:"scoreRange"    description:"得分区间包含关系（1：左闭右开，2：左开右闭）"`
	IsNeedPm      uint        `json:"isNeedPm"      description:"是否委派PM(1:是 2：否)"`
	PmDemand      string      `json:"pmDemand"      description:"pm要求"`
	ProductDemand string      `json:"productDemand" description:"项目工作相关要求"`
	MonitDemand   string      `json:"monitDemand"   description:"监控要求"`
	IsNeedPml     uint        `json:"isNeedPml"     description:"是否需要项目负责人(1:是 2：否)"`
	Remark        string      `json:"remark"        description:"预留备注说明信息"`
	CreateTime    *gtime.Time `json:"createTime"    description:"新增数据时间"`
	UpdateTime    *gtime.Time `json:"updateTime"    description:"最新更新数据"`
}
