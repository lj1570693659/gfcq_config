// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductBudgetAccess is the golang structure for table product_budget_access.
type ProductBudgetAccess struct {
	Id          int         `json:"id"          description:""`
	ScoreMin    uint        `json:"scoreMin"    description:"分值下限"`
	ScoreMax    uint        `json:"scoreMax"    description:"分值上限"`
	ScoreRange  uint        `json:"scoreRange"  description:"分数区间包含关系（1：左闭右开，2：左开右闭）"`
	BudgetMin   float64     `json:"budgetMin"   description:"预算额度下限"`
	BudgetMax   float64     `json:"budgetMax"   description:"预算额度上线"`
	BudgetRange uint        `json:"budgetRange" description:"额度区间包含关系（1：左闭右开，2：左开右闭）"`
	Remark      string      `json:"remark"      description:"预留备注说明信息"`
	CreateTime  *gtime.Time `json:"createTime"  description:"新增数据时间"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"最后一次更新数据时间"`
}
