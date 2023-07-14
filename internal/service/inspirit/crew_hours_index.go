package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	ICrewHoursIndex interface {
		Create(ctx context.Context, info *v1.CreateCrewHoursIndexReq) (*v1.CreateCrewHoursIndexRes, error)
		GetOne(ctx context.Context, info *v1.GetOneCrewHoursIndexReq) (*v1.GetOneCrewHoursIndexRes, error)
		GetList(ctx context.Context, info *v1.GetListCrewHoursIndexReq) (*v1.GetListCrewHoursIndexRes, error)
		Modify(ctx context.Context, info *v1.ModifyCrewHoursIndexReq) (*v1.ModifyCrewHoursIndexRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localCrewHoursIndex ICrewHoursIndex
)

func CrewHoursIndex() ICrewHoursIndex {
	if localCrewHoursIndex == nil {
		panic("implement not found for interface ICrewHoursIndex, forgot register?")
	}
	return localCrewHoursIndex
}

func RegisterCrewHoursIndex(i ICrewHoursIndex) {
	localCrewHoursIndex = i
}
