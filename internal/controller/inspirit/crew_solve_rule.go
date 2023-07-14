package inspirit

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewSolveRuleController 团队成员浮动贡献-解决问题贡献配置信息
type CrewSolveRuleController struct {
	v1.UnimplementedCrewSolveRuleServer
}

// CrewSolveRuleRegister 项目等级评估配置信息
func CrewSolveRuleRegister(s *grpcx.GrpcServer) {
	v1.RegisterCrewSolveRuleServer(s.Server, &CrewSolveRuleController{})
}

// GetList implements GetList
func (s *CrewSolveRuleController) GetList(ctx context.Context, in *v1.GetListCrewSolveRuleReq) (*v1.GetListCrewSolveRuleRes, error) {
	res, err := service.CrewSolveRule().GetList(ctx, in)
	return res, err
}

func (s *CrewSolveRuleController) GetOne(ctx context.Context, in *v1.GetOneCrewSolveRuleReq) (*v1.GetOneCrewSolveRuleRes, error) {
	res, err := service.CrewSolveRule().GetOne(ctx, in)
	return res, err
}

func (s *CrewSolveRuleController) Create(ctx context.Context, in *v1.CreateCrewSolveRuleReq) (*v1.CreateCrewSolveRuleRes, error) {
	return service.CrewSolveRule().Create(ctx, in)
}

func (s *CrewSolveRuleController) Modify(ctx context.Context, in *v1.ModifyCrewSolveRuleReq) (*v1.ModifyCrewSolveRuleRes, error) {
	return service.CrewSolveRule().Modify(ctx, in)
}

func (s *CrewSolveRuleController) Delete(ctx context.Context, in *v1.DeleteCrewSolveRuleReq) (*v1.DeleteCrewSolveRuleRes, error) {
	isSuccess, msg, err := service.CrewSolveRule().Delete(ctx, in.GetId())
	return &v1.DeleteCrewSolveRuleRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
