package product

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	IRoles interface {
		Create(ctx context.Context, info *v1.CreateRolesReq) (*v1.CreateRolesRes, error)
		GetOne(ctx context.Context, info *v1.GetOneRolesReq) (*v1.GetOneRolesRes, error)
		GetList(ctx context.Context, info *v1.GetListRolesReq) (*v1.GetListRolesRes, error)
		GetAll(ctx context.Context, info *v1.GetAllRolesReq) (*v1.GetAllRolesRes, error)
		Modify(ctx context.Context, info *v1.ModifyRolesReq) (*v1.ModifyRolesRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localRoles IRoles
)

func Roles() IRoles {
	if localRoles == nil {
		panic("implement not found for interface IRoles, forgot register?")
	}
	return localRoles
}

func RegisterRoles(i IRoles) {
	localRoles = i
}
