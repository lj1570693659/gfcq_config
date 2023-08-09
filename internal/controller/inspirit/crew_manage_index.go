package inspirit

import (
	"context"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"google.golang.org/grpc"
)

// CrewManageIndexController 团队成员管理指数配置信息
type CrewManageIndexController struct {
	v1.UnimplementedCrewManageIndexServer
}

// CrewManageIndexRegister 项目等级评估配置信息
func CrewManageIndexRegister(s *grpc.Server) {
	v1.RegisterCrewManageIndexServer(s, &CrewManageIndexController{})
}

// GetList implements GetList
func (s *CrewManageIndexController) GetList(ctx context.Context, in *v1.GetListCrewManageIndexReq) (*v1.GetListCrewManageIndexRes, error) {
	res, err := service.CrewManageIndex().GetList(ctx, in)
	return res, err
}

// GetAll implements GetAll
func (s *CrewManageIndexController) GetAll(ctx context.Context, in *v1.GetAllCrewManageIndexReq) (*v1.GetAllCrewManageIndexRes, error) {
	res, err := service.CrewManageIndex().GetAll(ctx, in)
	return res, err
}

func (s *CrewManageIndexController) GetOne(ctx context.Context, in *v1.GetOneCrewManageIndexReq) (*v1.GetOneCrewManageIndexRes, error) {
	res, err := service.CrewManageIndex().GetOne(ctx, in)
	return res, err
}

func (s *CrewManageIndexController) Create(ctx context.Context, in *v1.CreateCrewManageIndexReq) (*v1.CreateCrewManageIndexRes, error) {
	return service.CrewManageIndex().Create(ctx, in)
}

func (s *CrewManageIndexController) Modify(ctx context.Context, in *v1.ModifyCrewManageIndexReq) (*v1.ModifyCrewManageIndexRes, error) {
	return service.CrewManageIndex().Modify(ctx, in)
}

func (s *CrewManageIndexController) Delete(ctx context.Context, in *v1.DeleteCrewManageIndexReq) (*v1.DeleteCrewManageIndexRes, error) {
	isSuccess, msg, err := service.CrewManageIndex().Delete(ctx, in.GetId())
	return &v1.DeleteCrewManageIndexRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
