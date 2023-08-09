// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewKpiRule is the golang structure for table crew_kpi_rule.
type CrewKpiRule struct {
	Id         uint        `json:"id"         description:""`
	Redio      float64     `json:"redio"      description:"比例"`
	ScoreMin   float64     `json:"scoreMin"   description:"得分下限"`
	ScoreMax   float64     `json:"scoreMax"   description:"得分上线"`
	ScoreRange uint        `json:"scoreRange" description:"得分区间包含关系（1：左闭右开，2：左开右闭）"`
	LevelName  string      `json:"levelName"  description:"等级名称"`
	Remark     string      `json:"remark"     description:"预留备注说明信息"`
	CreateTime *gtime.Time `json:"createTime" description:"新增数据时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
