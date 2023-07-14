// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewDutyIndex is the golang structure for table crew_duty_index.
type CrewDutyIndex struct {
	Id         uint        `json:"id"         description:""`
	ScoreIndex uint        `json:"scoreIndex" description:"职责指数"`
	JobLevelId uint        `json:"jobLevelId" description:"职级ID"`
	Arith      string      `json:"arith"      description:"运算方式"`
	Remark     string      `json:"remark"     description:"预留备注说明信息"`
	CreateTime *gtime.Time `json:"createTime" description:"新增数据时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
