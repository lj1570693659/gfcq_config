package inspirit

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	IStageRadio interface {
		Create(ctx context.Context, info *v1.CreateStageRadioReq) (*v1.CreateStageRadioRes, error)
		GetOne(ctx context.Context, info *v1.GetOneStageRadioReq) (*v1.GetOneStageRadioRes, error)
		GetList(ctx context.Context, info *v1.GetListStageRadioReq) (*v1.GetListStageRadioRes, error)
		GetAll(ctx context.Context, info *v1.GetAllStageRadioReq) (*v1.GetAllStageRadioRes, error)
		GetQuotaRadioByScore(ctx context.Context, info *v1.GetQuotaRadioByScoreReq) (*v1.GetQuotaRadioByScoreRes, error)
		Modify(ctx context.Context, info *v1.ModifyStageRadioReq) (*v1.ModifyStageRadioRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localStageRadio IStageRadio
)

func StageRadio() IStageRadio {
	if localStageRadio == nil {
		panic("implement not found for interface IStageRadio, forgot register?")
	}
	return localStageRadio
}

func RegisterStageRadio(i IStageRadio) {
	localStageRadio = i
}
