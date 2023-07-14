package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	ICrewOvertimeRule interface {
		Create(ctx context.Context, info *v1.CreateCrewOvertimeRuleReq) (*v1.CreateCrewOvertimeRuleRes, error)
		GetOne(ctx context.Context, info *v1.GetOneCrewOvertimeRuleReq) (*v1.GetOneCrewOvertimeRuleRes, error)
		GetList(ctx context.Context, info *v1.GetListCrewOvertimeRuleReq) (*v1.GetListCrewOvertimeRuleRes, error)
		Modify(ctx context.Context, info *v1.ModifyCrewOvertimeRuleReq) (*v1.ModifyCrewOvertimeRuleRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localCrewOvertimeRule ICrewOvertimeRule
)

func CrewOvertimeRule() ICrewOvertimeRule {
	if localCrewOvertimeRule == nil {
		panic("implement not found for interface ICrewOvertimeRule, forgot register?")
	}
	return localCrewOvertimeRule
}

func RegisterCrewOvertimeRule(i ICrewOvertimeRule) {
	localCrewOvertimeRule = i
}
