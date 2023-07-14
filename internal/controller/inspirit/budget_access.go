package inspirit

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// BudgetAccessController 项目得分-预算额度配置信息（激励预算）
type BudgetAccessController struct {
	v1.UnimplementedBudgetAssessServer
}

// BudgetAccessRegister 项目等级评估配置信息
func BudgetAccessRegister(s *grpcx.GrpcServer) {
	v1.RegisterBudgetAssessServer(s.Server, &BudgetAccessController{})
}

// GetList implements GetList
func (s *BudgetAccessController) GetList(ctx context.Context, in *v1.GetListBudgetAssessReq) (*v1.GetListBudgetAssessRes, error) {
	res, err := service.BudgetAccess().GetList(ctx, in)
	return res, err
}

func (s *BudgetAccessController) GetOne(ctx context.Context, in *v1.GetOneBudgetAssessReq) (*v1.GetOneBudgetAssessRes, error) {
	res, err := service.BudgetAccess().GetOne(ctx, in)
	return res, err
}

func (s *BudgetAccessController) Create(ctx context.Context, in *v1.CreateBudgetAssessReq) (*v1.CreateBudgetAssessRes, error) {
	return service.BudgetAccess().Create(ctx, in)
}

func (s *BudgetAccessController) Modify(ctx context.Context, in *v1.ModifyBudgetAssessReq) (*v1.ModifyBudgetAssessRes, error) {
	return service.BudgetAccess().Modify(ctx, in)
}

func (s *BudgetAccessController) Delete(ctx context.Context, in *v1.DeleteBudgetAssessReq) (*v1.DeleteBudgetAssessRes, error) {
	isSuccess, msg, err := service.BudgetAccess().Delete(ctx, in.GetId())
	return &v1.DeleteBudgetAssessRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
