package product

import (
	"context"
	"github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"google.golang.org/grpc"
)

// LevelAssessController 项目等级评估配置信息
type LevelAssessController struct {
	v1.UnimplementedLevelAssessServer
}

// LevelAssessRegister 项目等级评估配置信息
func LevelAssessRegister(s *grpc.Server) {
	v1.RegisterLevelAssessServer(s, &LevelAssessController{})
}

// GetList implements GetList
func (s *LevelAssessController) GetList(ctx context.Context, in *v1.GetListLevelAssessReq) (*v1.GetListLevelAssessRes, error) {
	res, err := product.LevelAssess().GetList(ctx, in)
	return res, err
}

// GetListWithoutPage implements GetListWithoutPage
func (*LevelAssessController) GetListWithoutPage(ctx context.Context, in *v1.GetListWithoutLevelAssessReq) (*v1.GetListWithoutLevelAssessRes, error) {
	res, err := product.LevelAssess().GetListWithoutPage(ctx, in)
	return res, err
}

func (s *LevelAssessController) GetOne(ctx context.Context, in *v1.GetOneLevelAssessReq) (*v1.GetOneLevelAssessRes, error) {
	res, err := product.LevelAssess().GetOne(ctx, in)
	return res, err
}

func (s *LevelAssessController) Create(ctx context.Context, in *v1.CreateLevelAssessReq) (*v1.CreateLevelAssessRes, error) {
	return product.LevelAssess().Create(ctx, in)
}

func (s *LevelAssessController) Modify(ctx context.Context, in *v1.ModifyLevelAssessReq) (*v1.ModifyLevelAssessRes, error) {
	return product.LevelAssess().Modify(ctx, in)
}

func (s *LevelAssessController) Delete(ctx context.Context, in *v1.DeleteLevelAssessReq) (*v1.DeleteLevelAssessRes, error) {
	isSuccess, msg, err := product.LevelAssess().Delete(ctx, in.GetId())
	return &v1.DeleteLevelAssessRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
