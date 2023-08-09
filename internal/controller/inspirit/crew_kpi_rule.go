package inspirit

import (
	"context"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"google.golang.org/grpc"
)

// CrewKpiRuleController 团队成员绩效等级配置信息
type CrewKpiRuleController struct {
	v1.UnimplementedCrewKpiRuleServer
}

// CrewKpiRuleRegister 项目等级评估配置信息
func CrewKpiRuleRegister(s *grpc.Server) {
	v1.RegisterCrewKpiRuleServer(s, &CrewKpiRuleController{})
}

// GetList implements GetList
func (s *CrewKpiRuleController) GetList(ctx context.Context, in *v1.GetListCrewKpiRuleReq) (*v1.GetListCrewKpiRuleRes, error) {
	res, err := service.CrewKpiRule().GetList(ctx, in)
	return res, err
}

// GetAll implements GetAll
func (s *CrewKpiRuleController) GetAll(ctx context.Context, in *v1.GetAllCrewKpiRuleReq) (*v1.GetAllCrewKpiRuleRes, error) {
	res, err := service.CrewKpiRule().GetAll(ctx, in)
	return res, err
}

func (s *CrewKpiRuleController) GetOne(ctx context.Context, in *v1.GetOneCrewKpiRuleReq) (*v1.GetOneCrewKpiRuleRes, error) {
	res, err := service.CrewKpiRule().GetOne(ctx, in)
	return res, err
}

func (s *CrewKpiRuleController) Create(ctx context.Context, in *v1.CreateCrewKpiRuleReq) (*v1.CreateCrewKpiRuleRes, error) {
	return service.CrewKpiRule().Create(ctx, in)
}

func (s *CrewKpiRuleController) Modify(ctx context.Context, in *v1.ModifyCrewKpiRuleReq) (*v1.ModifyCrewKpiRuleRes, error) {
	return service.CrewKpiRule().Modify(ctx, in)
}

func (s *CrewKpiRuleController) Delete(ctx context.Context, in *v1.DeleteCrewKpiRuleReq) (*v1.DeleteCrewKpiRuleRes, error) {
	isSuccess, msg, err := service.CrewKpiRule().Delete(ctx, in.GetId())
	return &v1.DeleteCrewKpiRuleRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
