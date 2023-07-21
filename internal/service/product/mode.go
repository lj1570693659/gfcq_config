package product

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	IMode interface {
		Create(ctx context.Context, info *v1.CreateModeReq) (*v1.CreateModeRes, error)
		GetOne(ctx context.Context, info *v1.GetOneModeReq) (*v1.GetOneModeRes, error)
		GetList(ctx context.Context, info *v1.GetListModeReq) (*v1.GetListModeRes, error)
		GetAll(ctx context.Context, info *v1.GetAllModeReq) (*v1.GetAllModeRes, error)
		Modify(ctx context.Context, info *v1.ModifyModeReq) (*v1.ModifyModeRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localMode IMode
)

func Mode() IMode {
	if localMode == nil {
		panic("implement not found for interface IMode, forgot register?")
	}
	return localMode
}

func RegisterMode(i IMode) {
	localMode = i
}
