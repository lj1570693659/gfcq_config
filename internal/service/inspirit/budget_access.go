package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	IBudgetAccess interface {
		Create(ctx context.Context, info *v1.CreateBudgetAssessReq) (*v1.CreateBudgetAssessRes, error)
		GetOne(ctx context.Context, info *v1.GetOneBudgetAssessReq) (*v1.GetOneBudgetAssessRes, error)
		GetList(ctx context.Context, info *v1.GetListBudgetAssessReq) (*v1.GetListBudgetAssessRes, error)
		Modify(ctx context.Context, info *v1.ModifyBudgetAssessReq) (*v1.ModifyBudgetAssessRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localBudgetAccess IBudgetAccess
)

func BudgetAccess() IBudgetAccess {
	if localBudgetAccess == nil {
		panic("implement not found for interface IBudgetAccess, forgot register?")
	}
	return localBudgetAccess
}

func RegisterBudgetAccess(i IBudgetAccess) {
	localBudgetAccess = i
}
