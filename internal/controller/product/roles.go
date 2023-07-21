package product

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// RolesController 项目角色配置信息
type RolesController struct {
	v1.UnimplementedRolesServer
}

func RolesRegister(s *grpcx.GrpcServer) {
	v1.RegisterRolesServer(s.Server, &RolesController{})
}

// GetList implements GetList
func (s *RolesController) GetList(ctx context.Context, in *v1.GetListRolesReq) (*v1.GetListRolesRes, error) {
	return service.Roles().GetList(ctx, in)
}

// GetAll implements GetAll
func (s *RolesController) GetAll(ctx context.Context, in *v1.GetAllRolesReq) (*v1.GetAllRolesRes, error) {
	return service.Roles().GetAll(ctx, in)
}

func (s *RolesController) GetOne(ctx context.Context, in *v1.GetOneRolesReq) (*v1.GetOneRolesRes, error) {
	return service.Roles().GetOne(ctx, in)
}

func (s *RolesController) Create(ctx context.Context, in *v1.CreateRolesReq) (*v1.CreateRolesRes, error) {
	return service.Roles().Create(ctx, in)
}

func (s *RolesController) Modify(ctx context.Context, in *v1.ModifyRolesReq) (*v1.ModifyRolesRes, error) {
	return service.Roles().Modify(ctx, in)
}

func (s *RolesController) Delete(ctx context.Context, in *v1.DeleteRolesReq) (*v1.DeleteRolesRes, error) {
	isSuccess, msg, err := service.Roles().Delete(ctx, in.GetId())
	return &v1.DeleteRolesRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
