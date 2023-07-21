package product

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	IModeStage interface {
		Create(ctx context.Context, info *v1.CreateModeStageReq) (*v1.CreateModeStageRes, error)
		GetOne(ctx context.Context, info *v1.GetOneModeStageReq) (*v1.GetOneModeStageRes, error)
		GetList(ctx context.Context, info *v1.GetListModeStageReq) (*v1.GetListModeStageRes, error)
		GetAll(ctx context.Context, info *v1.GetAllModeStageReq) (*v1.GetAllModeStageRes, error)
		Modify(ctx context.Context, info *v1.ModifyModeStageReq) (*v1.ModifyModeStageRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localModeStage IModeStage
)

func ModeStage() IModeStage {
	if localModeStage == nil {
		panic("implement not found for interface IModeStage, forgot register?")
	}
	return localModeStage
}

func RegisterModeStage(i IModeStage) {
	localModeStage = i
}
