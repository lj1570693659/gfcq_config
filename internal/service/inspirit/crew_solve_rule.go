package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	ICrewSolveRule interface {
		Create(ctx context.Context, info *v1.CreateCrewSolveRuleReq) (*v1.CreateCrewSolveRuleRes, error)
		GetOne(ctx context.Context, info *v1.GetOneCrewSolveRuleReq) (*v1.GetOneCrewSolveRuleRes, error)
		GetList(ctx context.Context, info *v1.GetListCrewSolveRuleReq) (*v1.GetListCrewSolveRuleRes, error)
		GetAll(ctx context.Context, info *v1.GetAllCrewSolveRuleReq) (*v1.GetAllCrewSolveRuleRes, error)
		Modify(ctx context.Context, info *v1.ModifyCrewSolveRuleReq) (*v1.ModifyCrewSolveRuleRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localCrewSolveRule ICrewSolveRule
)

func CrewSolveRule() ICrewSolveRule {
	if localCrewSolveRule == nil {
		panic("implement not found for interface ICrewSolveRule, forgot register?")
	}
	return localCrewSolveRule
}

func RegisterCrewSolveRule(i ICrewSolveRule) {
	localCrewSolveRule = i
}
