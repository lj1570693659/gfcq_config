package product

import v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/product"
)

// TypeController 项目类型信息
type TypeController struct {
	v1.UnimplementedTypeServer
}

func TypeRegister(s *grpcx.GrpcServer) {
	v1.RegisterTypeServer(s.Server, &TypeController{})
}

// GetList implements GetList
func (s *TypeController) GetList(ctx context.Context, in *v1.GetListTypeReq) (*v1.GetListTypeRes, error) {
	res, err := service.Type().GetList(ctx, in)
	return res, err
}

// GetAll implements GetAll
func (s *TypeController) GetAll(ctx context.Context, in *v1.GetAllTypeReq) (*v1.GetAllTypeRes, error) {
	res, err := service.Type().GetAll(ctx, in)
	return res, err
}

func (s *TypeController) GetOne(ctx context.Context, in *v1.GetOneTypeReq) (*v1.GetOneTypeRes, error) {
	res, err := service.Type().GetOne(ctx, in)
	return res, err
}
func (s *TypeController) Create(ctx context.Context, in *v1.CreateTypeReq) (*v1.CreateTypeRes, error) {
	res, err := service.Type().Create(ctx, in)
	return res, err
}
func (s *TypeController) Modify(ctx context.Context, in *v1.ModifyTypeReq) (*v1.ModifyTypeRes, error) {
	res, err := service.Type().Modify(ctx, in)
	return res, err
}
func (s *TypeController) Delete(ctx context.Context, in *v1.DeleteTypeReq) (*v1.DeleteTypeRes, error) {
	isSuccess, msg, err := service.Type().Delete(ctx, in.GetId())
	return &v1.DeleteTypeRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
