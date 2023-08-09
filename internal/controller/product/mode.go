package product

import (
	"context"
	service "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"google.golang.org/grpc"
)

// ModeController 项目开发模式信息
type ModeController struct {
	v1.UnimplementedModeServer
}

func ModeRegister(s *grpc.Server) {
	v1.RegisterModeServer(s, &ModeController{})
}

// GetList implements GetList
func (s *ModeController) GetList(ctx context.Context, in *v1.GetListModeReq) (*v1.GetListModeRes, error) {
	return service.Mode().GetList(ctx, in)
}

func (s *ModeController) GetAll(ctx context.Context, in *v1.GetAllModeReq) (*v1.GetAllModeRes, error) {
	return service.Mode().GetAll(ctx, in)
}

func (s *ModeController) GetOne(ctx context.Context, in *v1.GetOneModeReq) (*v1.GetOneModeRes, error) {
	return service.Mode().GetOne(ctx, in)
}

func (s *ModeController) Create(ctx context.Context, in *v1.CreateModeReq) (*v1.CreateModeRes, error) {
	return service.Mode().Create(ctx, in)
}

func (s *ModeController) Modify(ctx context.Context, in *v1.ModifyModeReq) (*v1.ModifyModeRes, error) {
	return service.Mode().Modify(ctx, in)
}

func (s *ModeController) Delete(ctx context.Context, in *v1.DeleteModeReq) (*v1.DeleteModeRes, error) {
	isSuccess, msg, err := service.Mode().Delete(ctx, in.GetId())
	return &v1.DeleteModeRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
