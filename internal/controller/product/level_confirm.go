package product

import (
	"context"
	service "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"google.golang.org/grpc"
)

// LevelConfirmController 项目优先级确认配置信息
type LevelConfirmController struct {
	v1.UnimplementedLevelConfirmServer
}

func LevelConfirmRegister(s *grpc.Server) {
	v1.RegisterLevelConfirmServer(s, &LevelConfirmController{})
}

// GetList implements GetList
func (s *LevelConfirmController) GetList(ctx context.Context, in *v1.GetListLevelConfirmReq) (*v1.GetListLevelConfirmRes, error) {
	return service.LevelConfirm().GetList(ctx, in)
}

func (s *LevelConfirmController) GetOne(ctx context.Context, in *v1.GetOneLevelConfirmReq) (*v1.GetOneLevelConfirmRes, error) {
	return service.LevelConfirm().GetOne(ctx, in)
}

func (s *LevelConfirmController) Create(ctx context.Context, in *v1.CreateLevelConfirmReq) (*v1.CreateLevelConfirmRes, error) {
	return service.LevelConfirm().Create(ctx, in)
}

func (s *LevelConfirmController) Modify(ctx context.Context, in *v1.ModifyLevelConfirmReq) (*v1.ModifyLevelConfirmRes, error) {
	return service.LevelConfirm().Modify(ctx, in)
}

func (s *LevelConfirmController) Delete(ctx context.Context, in *v1.DeleteLevelConfirmReq) (*v1.DeleteLevelConfirmRes, error) {
	isSuccess, msg, err := service.LevelConfirm().Delete(ctx, in.GetId())
	return &v1.DeleteLevelConfirmRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
