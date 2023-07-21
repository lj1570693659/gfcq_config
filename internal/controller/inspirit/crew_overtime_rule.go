package inspirit

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewOvertimeRuleController 项目组成员浮动贡献-加班贡献配置信息
type CrewOvertimeRuleController struct {
	v1.UnimplementedCrewOvertimeRuleServer
}

// CrewOvertimeRuleRegister 项目等级评估配置信息
func CrewOvertimeRuleRegister(s *grpcx.GrpcServer) {
	v1.RegisterCrewOvertimeRuleServer(s.Server, &CrewOvertimeRuleController{})
}

// GetList implements GetList
func (s *CrewOvertimeRuleController) GetList(ctx context.Context, in *v1.GetListCrewOvertimeRuleReq) (*v1.GetListCrewOvertimeRuleRes, error) {
	res, err := service.CrewOvertimeRule().GetList(ctx, in)
	return res, err
}

// GetAll implements GetAll
func (s *CrewOvertimeRuleController) GetAll(ctx context.Context, in *v1.GetAllCrewOvertimeRuleReq) (*v1.GetAllCrewOvertimeRuleRes, error) {
	res, err := service.CrewOvertimeRule().GetAll(ctx, in)
	return res, err
}

func (s *CrewOvertimeRuleController) GetOne(ctx context.Context, in *v1.GetOneCrewOvertimeRuleReq) (*v1.GetOneCrewOvertimeRuleRes, error) {
	res, err := service.CrewOvertimeRule().GetOne(ctx, in)
	return res, err
}

func (s *CrewOvertimeRuleController) Create(ctx context.Context, in *v1.CreateCrewOvertimeRuleReq) (*v1.CreateCrewOvertimeRuleRes, error) {
	return service.CrewOvertimeRule().Create(ctx, in)
}

func (s *CrewOvertimeRuleController) Modify(ctx context.Context, in *v1.ModifyCrewOvertimeRuleReq) (*v1.ModifyCrewOvertimeRuleRes, error) {
	return service.CrewOvertimeRule().Modify(ctx, in)
}

func (s *CrewOvertimeRuleController) Delete(ctx context.Context, in *v1.DeleteCrewOvertimeRuleReq) (*v1.DeleteCrewOvertimeRuleRes, error) {
	isSuccess, msg, err := service.CrewOvertimeRule().Delete(ctx, in.GetId())
	return &v1.DeleteCrewOvertimeRuleRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
