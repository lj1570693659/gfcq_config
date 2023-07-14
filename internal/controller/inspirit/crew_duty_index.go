package inspirit

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewDutyIndexController 项目得分-预算额度配置信息（激励预算）
type CrewDutyIndexController struct {
	v1.UnimplementedCrewDutyIndexServer
}

// CrewDutyIndexRegister 项目等级评估配置信息
func CrewDutyIndexRegister(s *grpcx.GrpcServer) {
	v1.RegisterCrewDutyIndexServer(s.Server, &CrewDutyIndexController{})
}

// GetList implements GetList
func (s *CrewDutyIndexController) GetList(ctx context.Context, in *v1.GetListCrewDutyIndexReq) (*v1.GetListCrewDutyIndexRes, error) {
	res, err := service.CrewDutyIndex().GetList(ctx, in)
	return res, err
}

func (s *CrewDutyIndexController) GetOne(ctx context.Context, in *v1.GetOneCrewDutyIndexReq) (*v1.GetOneCrewDutyIndexRes, error) {
	res, err := service.CrewDutyIndex().GetOne(ctx, in)
	return res, err
}

func (s *CrewDutyIndexController) Create(ctx context.Context, in *v1.CreateCrewDutyIndexReq) (*v1.CreateCrewDutyIndexRes, error) {
	return service.CrewDutyIndex().Create(ctx, in)
}

func (s *CrewDutyIndexController) Modify(ctx context.Context, in *v1.ModifyCrewDutyIndexReq) (*v1.ModifyCrewDutyIndexRes, error) {
	return service.CrewDutyIndex().Modify(ctx, in)
}

func (s *CrewDutyIndexController) Delete(ctx context.Context, in *v1.DeleteCrewDutyIndexReq) (*v1.DeleteCrewDutyIndexRes, error) {
	isSuccess, msg, err := service.CrewDutyIndex().Delete(ctx, in.GetId())
	return &v1.DeleteCrewDutyIndexRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
