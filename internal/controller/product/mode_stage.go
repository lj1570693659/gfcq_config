package product

import (
	"context"
	service "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"google.golang.org/grpc"
)

// ModeStageController 项目开发模式信息
type ModeStageController struct {
	v1.UnimplementedModeStageServer
}

func ModeStageRegister(s *grpc.Server) {
	v1.RegisterModeStageServer(s, &ModeStageController{})
}

// GetList implements GetList
func (s *ModeStageController) GetList(ctx context.Context, in *v1.GetListModeStageReq) (*v1.GetListModeStageRes, error) {
	return service.ModeStage().GetList(ctx, in)
}

// GetAll implements GetAll
func (s *ModeStageController) GetAll(ctx context.Context, in *v1.GetAllModeStageReq) (*v1.GetAllModeStageRes, error) {
	return service.ModeStage().GetAll(ctx, in)
}

func (s *ModeStageController) GetOne(ctx context.Context, in *v1.GetOneModeStageReq) (*v1.GetOneModeStageRes, error) {
	return service.ModeStage().GetOne(ctx, in)
}

func (s *ModeStageController) Create(ctx context.Context, in *v1.CreateModeStageReq) (*v1.CreateModeStageRes, error) {
	return service.ModeStage().Create(ctx, in)
}

func (s *ModeStageController) Modify(ctx context.Context, in *v1.ModifyModeStageReq) (*v1.ModifyModeStageRes, error) {
	return service.ModeStage().Modify(ctx, in)
}

func (s *ModeStageController) Delete(ctx context.Context, in *v1.DeleteModeStageReq) (*v1.DeleteModeStageRes, error) {
	isSuccess, msg, err := service.ModeStage().Delete(ctx, in.GetId())
	return &v1.DeleteModeStageRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
