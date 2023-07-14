package inspirit

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// ProductStageRadioController 项目阶段应发激励占比配置信息
type ProductStageRadioController struct {
	v1.UnimplementedStageRadioServer
}

// ProductStageRadioRegister 项目等级评估配置信息
func ProductStageRadioRegister(s *grpcx.GrpcServer) {
	v1.RegisterStageRadioServer(s.Server, &ProductStageRadioController{})
}

// GetList implements GetList
func (s *ProductStageRadioController) GetList(ctx context.Context, in *v1.GetListStageRadioReq) (*v1.GetListStageRadioRes, error) {
	res, err := service.StageRadio().GetList(ctx, in)
	return res, err
}

func (s *ProductStageRadioController) GetOne(ctx context.Context, in *v1.GetOneStageRadioReq) (*v1.GetOneStageRadioRes, error) {
	res, err := service.StageRadio().GetOne(ctx, in)
	return res, err
}

func (s *ProductStageRadioController) Create(ctx context.Context, in *v1.CreateStageRadioReq) (*v1.CreateStageRadioRes, error) {
	return service.StageRadio().Create(ctx, in)
}

func (s *ProductStageRadioController) Modify(ctx context.Context, in *v1.ModifyStageRadioReq) (*v1.ModifyStageRadioRes, error) {
	return service.StageRadio().Modify(ctx, in)
}

func (s *ProductStageRadioController) Delete(ctx context.Context, in *v1.DeleteStageRadioReq) (*v1.DeleteStageRadioRes, error) {
	isSuccess, msg, err := service.StageRadio().Delete(ctx, in.GetId())
	return &v1.DeleteStageRadioRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
