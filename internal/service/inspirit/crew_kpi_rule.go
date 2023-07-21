package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	ICrewKpiRule interface {
		Create(ctx context.Context, info *v1.CreateCrewKpiRuleReq) (*v1.CreateCrewKpiRuleRes, error)
		GetOne(ctx context.Context, info *v1.GetOneCrewKpiRuleReq) (*v1.GetOneCrewKpiRuleRes, error)
		GetList(ctx context.Context, info *v1.GetListCrewKpiRuleReq) (*v1.GetListCrewKpiRuleRes, error)
		//GetAll(ctx context.Context, info *v1.GetAllCrewKpiRuleReq) (*v1.GetAllCrewKpiRuleRes, error)
		Modify(ctx context.Context, info *v1.ModifyCrewKpiRuleReq) (*v1.ModifyCrewKpiRuleRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localCrewKpiRule ICrewKpiRule
)

func CrewKpiRule() ICrewKpiRule {
	if localCrewKpiRule == nil {
		panic("implement not found for interface ICrewKpiRule, forgot register?")
	}
	return localCrewKpiRule
}

func RegisterCrewKpiRule(i ICrewKpiRule) {
	localCrewKpiRule = i
}
