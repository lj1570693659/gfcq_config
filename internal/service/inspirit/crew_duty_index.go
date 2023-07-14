package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	ICrewDutyIndex interface {
		Create(ctx context.Context, info *v1.CreateCrewDutyIndexReq) (*v1.CreateCrewDutyIndexRes, error)
		GetOne(ctx context.Context, info *v1.GetOneCrewDutyIndexReq) (*v1.GetOneCrewDutyIndexRes, error)
		GetList(ctx context.Context, info *v1.GetListCrewDutyIndexReq) (*v1.GetListCrewDutyIndexRes, error)
		Modify(ctx context.Context, info *v1.ModifyCrewDutyIndexReq) (*v1.ModifyCrewDutyIndexRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localCrewDutyIndex ICrewDutyIndex
)

func CrewDutyIndex() ICrewDutyIndex {
	if localCrewDutyIndex == nil {
		panic("implement not found for interface ICrewDutyIndex, forgot register?")
	}
	return localCrewDutyIndex
}

func RegisterCrewDutyIndex(i ICrewDutyIndex) {
	localCrewDutyIndex = i
}
