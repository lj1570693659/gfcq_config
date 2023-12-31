package product

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	IType interface {
		GetOne(ctx context.Context, info *v1.GetOneTypeReq) (*v1.GetOneTypeRes, error)
		GetList(ctx context.Context, info *v1.GetListTypeReq) (*v1.GetListTypeRes, error)
		GetAll(ctx context.Context, info *v1.GetAllTypeReq) (*v1.GetAllTypeRes, error)
		Create(ctx context.Context, info *v1.CreateTypeReq) (*v1.CreateTypeRes, error)
		Modify(ctx context.Context, info *v1.ModifyTypeReq) (*v1.ModifyTypeRes, error)
		Delete(ctx context.Context, id int32) (bool, string, error)
	}
)

var (
	localType IType
)

func Type() IType {
	if localType == nil {
		panic("implement not found for interface IType, forgot register?")
	}
	return localType
}

func RegisterType(i IType) {
	localType = i
}
