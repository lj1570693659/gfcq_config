package inspirit

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewHoursIndexController 团队成员工时指数配置信息
type CrewHoursIndexController struct {
	v1.UnimplementedCrewHoursIndexServer
}

// CrewHoursIndexRegister 项目等级评估配置信息
func CrewHoursIndexRegister(s *grpcx.GrpcServer) {
	v1.RegisterCrewHoursIndexServer(s.Server, &CrewHoursIndexController{})
}

// GetList implements GetList
func (s *CrewHoursIndexController) GetList(ctx context.Context, in *v1.GetListCrewHoursIndexReq) (*v1.GetListCrewHoursIndexRes, error) {
	res, err := service.CrewHoursIndex().GetList(ctx, in)
	return res, err
}

func (s *CrewHoursIndexController) GetAll(ctx context.Context, in *v1.GetAllCrewHoursIndexReq) (*v1.GetAllCrewHoursIndexRes, error) {
	res, err := service.CrewHoursIndex().GetAll(ctx, in)
	return res, err
}

func (s *CrewHoursIndexController) GetOne(ctx context.Context, in *v1.GetOneCrewHoursIndexReq) (*v1.GetOneCrewHoursIndexRes, error) {
	res, err := service.CrewHoursIndex().GetOne(ctx, in)
	return res, err
}

func (s *CrewHoursIndexController) Create(ctx context.Context, in *v1.CreateCrewHoursIndexReq) (*v1.CreateCrewHoursIndexRes, error) {
	return service.CrewHoursIndex().Create(ctx, in)
}

func (s *CrewHoursIndexController) Modify(ctx context.Context, in *v1.ModifyCrewHoursIndexReq) (*v1.ModifyCrewHoursIndexRes, error) {
	return service.CrewHoursIndex().Modify(ctx, in)
}

func (s *CrewHoursIndexController) Delete(ctx context.Context, in *v1.DeleteCrewHoursIndexReq) (*v1.DeleteCrewHoursIndexRes, error) {
	isSuccess, msg, err := service.CrewHoursIndex().Delete(ctx, in.GetId())
	return &v1.DeleteCrewHoursIndexRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
