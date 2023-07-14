package product

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	ILevelConfirm interface {
		Create(ctx context.Context, info *v1.CreateLevelConfirmReq) (*v1.CreateLevelConfirmRes, error)
		GetOne(ctx context.Context, info *v1.GetOneLevelConfirmReq) (*v1.GetOneLevelConfirmRes, error)
		GetList(ctx context.Context, info *v1.GetListLevelConfirmReq) (*v1.GetListLevelConfirmRes, error)
		Modify(ctx context.Context, info *v1.ModifyLevelConfirmReq) (*v1.ModifyLevelConfirmRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localLevelConfirm ILevelConfirm
)

func LevelConfirm() ILevelConfirm {
	if localLevelConfirm == nil {
		panic("implement not found for interface ILevelConfirm, forgot register?")
	}
	return localLevelConfirm
}

func RegisterLevelConfirm(i ILevelConfirm) {
	localLevelConfirm = i
}
