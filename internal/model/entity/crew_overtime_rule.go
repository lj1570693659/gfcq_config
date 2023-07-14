// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewOvertimeRule is the golang structure for table crew_overtime_rule.
type CrewOvertimeRule struct {
	Id         uint        `json:"id"         description:""`
	Redio      float64     `json:"redio"      description:"浮动比例"`
	ScoreMin   float64     `json:"scoreMin"   description:"得分下限"`
	ScoreMax   float64     `json:"scoreMax"   description:"得分上线"`
	ScoreRange uint        `json:"scoreRange" description:"得分区间包含关系（1：左闭右开，2：左开右闭）"`
	Remark     string      `json:"remark"     description:"预留备注说明信息"`
	CreateTime *gtime.Time `json:"createTime" description:"新增数据时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
