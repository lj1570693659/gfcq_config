// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductLevelAssess is the golang structure for table product_level_assess.
type ProductLevelAssess struct {
	Id                 int         `json:"id"                 description:""`
	EvaluateDimensions string      `json:"evaluateDimensions" description:"评价维度"`
	EvaluateCriteria   string      `json:"evaluateCriteria"   description:"评价标准"`
	ScoreCriteria      string      `json:"scoreCriteria"      description:"评分标准"`
	EvaluateId         uint        `json:"evaluateId"         description:"上级评价维度"`
	Weight             float64     `json:"weight"             description:"权重"`
	Remark             string      `json:"remark"             description:"预留备注说明信息"`
	CreateTime         *gtime.Time `json:"createTime"         description:"新增数据时间"`
	UpdateTime         *gtime.Time `json:"updateTime"         description:"最新更新数据"`
}
