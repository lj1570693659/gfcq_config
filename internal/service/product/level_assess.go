package product

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	ILevelAssess interface {
		Create(ctx context.Context, info *v1.CreateLevelAssessReq) (*v1.CreateLevelAssessRes, error)
		GetOne(ctx context.Context, info *v1.GetOneLevelAssessReq) (*v1.GetOneLevelAssessRes, error)
		GetList(ctx context.Context, info *v1.GetListLevelAssessReq) (*v1.GetListLevelAssessRes, error)
		Modify(ctx context.Context, info *v1.ModifyLevelAssessReq) (*v1.ModifyLevelAssessRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localLevelAssess ILevelAssess
)

func LevelAssess() ILevelAssess {
	if localLevelAssess == nil {
		panic("implement not found for interface ILevelAssess, forgot register?")
	}
	return localLevelAssess
}

func RegisterLevelAssess(i ILevelAssess) {
	localLevelAssess = i
}
