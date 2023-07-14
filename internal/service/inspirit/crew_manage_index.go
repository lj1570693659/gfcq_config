package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	ICrewManageIndex interface {
		Create(ctx context.Context, info *v1.CreateCrewManageIndexReq) (*v1.CreateCrewManageIndexRes, error)
		GetOne(ctx context.Context, info *v1.GetOneCrewManageIndexReq) (*v1.GetOneCrewManageIndexRes, error)
		GetList(ctx context.Context, info *v1.GetListCrewManageIndexReq) (*v1.GetListCrewManageIndexRes, error)
		Modify(ctx context.Context, info *v1.ModifyCrewManageIndexReq) (*v1.ModifyCrewManageIndexRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localCrewManageIndex ICrewManageIndex
)

func CrewManageIndex() ICrewManageIndex {
	if localCrewManageIndex == nil {
		panic("implement not found for interface ICrewManageIndex, forgot register?")
	}
	return localCrewManageIndex
}

func RegisterCrewManageIndex(i ICrewManageIndex) {
	localCrewManageIndex = i
}
