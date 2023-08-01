package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
	"github.com/lj1570693659/gfcq_config/internal/model/do"
	"github.com/lj1570693659/gfcq_config/internal/model/entity"
	productService "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	sLevelAssess struct{}
)

func init() {
	productService.RegisterLevelAssess(&sLevelAssess{})
}

func (s *sLevelAssess) Create(ctx context.Context, in *v1.CreateLevelAssessReq) (*v1.CreateLevelAssessRes, error) {
	res := &v1.CreateLevelAssessRes{
		LevelAssess: &v1.LevelAssessInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.LevelAssessInfo{
		EvaluateDimensions: in.GetEvaluateDimensions(),
		EvaluateCriteria:   in.GetEvaluateCriteria(),
		ScoreCriteria:      in.GetScoreCriteria(),
		EvaluateId:         in.GetEvaluateId(),
		Weight:             in.GetWeight(),
		Remark:             in.Remark,
	}); err != nil {
		return res, err
	}

	data := do.ProductLevelAssess{
		EvaluateDimensions: in.GetEvaluateDimensions(),
		EvaluateCriteria:   in.GetEvaluateCriteria(),
		ScoreCriteria:      in.GetScoreCriteria(),
		EvaluateId:         in.GetEvaluateId(),
		Weight:             in.GetWeight(),
		Remark:             in.Remark,
		CreateTime:         gtime.Now(),
		UpdateTime:         gtime.Now(),
	}

	lastInsertId, err := dao.ProductLevelAssess.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.LevelAssess.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sLevelAssess) GetOne(ctx context.Context, in *v1.GetOneLevelAssessReq) (*v1.GetOneLevelAssessRes, error) {
	var info *v1.LevelAssessInfo
	query := dao.ProductLevelAssess.Ctx(ctx)

	// 评价标准
	if len(in.GetLevelAssess().GetEvaluateCriteria()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelAssess.Columns().EvaluateCriteria), g.Slice{fmt.Sprintf("%s%s", in.GetLevelAssess().GetEvaluateCriteria(), "%")})
	}
	// 主键查询
	if in.GetLevelAssess().GetId() > 0 {
		query = query.Where(dao.ProductLevelAssess.Columns().Id, in.GetLevelAssess().GetId())
	}
	if in.GetLevelAssess().GetEvaluateId() > 0 {
		query = query.Where(dao.ProductLevelAssess.Columns().EvaluateId, in.GetLevelAssess().GetEvaluateId())
	}

	err := query.Scan(&info)

	return &v1.GetOneLevelAssessRes{
		LevelAssess: info,
	}, err
}

func (s *sLevelAssess) GetList(ctx context.Context, in *v1.GetListLevelAssessReq) (*v1.GetListLevelAssessRes, error) {
	res := &v1.GetListLevelAssessRes{}
	resData := make([]*v1.LevelAssessInfo, 0)
	levelEntity := make([]entity.ProductLevelAssess, 0)

	query := dao.ProductLevelAssess.Ctx(ctx)

	// 评价标准
	if len(in.GetLevelAssess().GetEvaluateCriteria()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelAssess.Columns().EvaluateCriteria), g.Slice{fmt.Sprintf("%s%s", in.GetLevelAssess().GetEvaluateCriteria(), "%")})
	}
	// 主键查询
	if in.GetLevelAssess().GetId() > 0 {
		query = query.Where(dao.ProductLevelAssess.Columns().Id, in.GetLevelAssess().GetId())
	}
	if in.GetLevelAssess().GetEvaluateId() > 0 {
		query = query.Where(dao.ProductLevelAssess.Columns().EvaluateId, in.GetLevelAssess().GetEvaluateId())
	}
	// 主键
	if len(in.GetLevelAssess().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelAssess.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetLevelAssess().GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, in.GetPage(), in.GetSize())
	if err != nil {
		return res, err
	}
	err = query.Scan(&levelEntity)

	levelEntityByte, _ := json.Marshal(levelEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Page = in.GetPage()
	res.Size = in.GetSize()
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sLevelAssess) GetListWithoutPage(ctx context.Context, in *v1.GetListWithoutLevelAssessReq) (*v1.GetListWithoutLevelAssessRes, error) {
	res := &v1.GetListWithoutLevelAssessRes{}
	resData := make([]*v1.LevelAssessInfo, 0)
	levelEntity := make([]entity.ProductLevelAssess, 0)

	query := dao.ProductLevelAssess.Ctx(ctx)

	// 评价标准
	if len(in.GetLevelAssess().GetEvaluateCriteria()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelAssess.Columns().EvaluateCriteria), g.Slice{fmt.Sprintf("%s%s", in.GetLevelAssess().GetEvaluateCriteria(), "%")})
	}
	// 主键查询
	if in.GetLevelAssess().GetId() > 0 {
		query = query.Where(dao.ProductLevelAssess.Columns().Id, in.GetLevelAssess().GetId())
	}
	if in.GetLevelAssess().GetEvaluateId() > 0 {
		query = query.Where(dao.ProductLevelAssess.Columns().EvaluateId, in.GetLevelAssess().GetEvaluateId())
	} else if in.GetLevelAssess().GetEvaluateId() == -1 {
		query = query.Where(dao.ProductLevelAssess.Columns().EvaluateId, 0)
	}
	// 主键
	if len(in.GetLevelAssess().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelAssess.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetLevelAssess().GetRemark(), "%")})
	}

	err := query.Scan(&levelEntity)

	levelEntityByte, _ := json.Marshal(levelEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Data = resData
	return res, err
}

func (s *sLevelAssess) Modify(ctx context.Context, in *v1.ModifyLevelAssessReq) (*v1.ModifyLevelAssessRes, error) {
	res := &v1.ModifyLevelAssessRes{LevelAssess: &v1.LevelAssessInfo{
		Id:                 in.GetId(),
		EvaluateDimensions: in.GetEvaluateDimensions(),
		EvaluateCriteria:   in.GetEvaluateCriteria(),
		ScoreCriteria:      in.GetScoreCriteria(),
		EvaluateId:         in.GetEvaluateId(),
		Weight:             in.GetWeight(),
		Remark:             in.Remark,
	}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	if _, err := s.checkInputData(ctx, &v1.LevelAssessInfo{
		Id:                 in.GetId(),
		EvaluateDimensions: in.GetEvaluateDimensions(),
		EvaluateCriteria:   in.GetEvaluateCriteria(),
		ScoreCriteria:      in.GetScoreCriteria(),
		EvaluateId:         in.GetEvaluateId(),
		Weight:             in.GetWeight(),
		Remark:             in.Remark,
	}); err != nil {
		return res, err
	}

	data := do.ProductLevelAssess{
		EvaluateDimensions: in.GetEvaluateDimensions(),
		EvaluateCriteria:   in.GetEvaluateCriteria(),
		ScoreCriteria:      in.GetScoreCriteria(),
		EvaluateId:         in.GetEvaluateId(),
		Weight:             in.GetWeight(),
		Remark:             in.Remark,
		UpdateTime:         gtime.Now(),
	}

	_, err := dao.ProductLevelAssess.Ctx(ctx).Where(dao.ProductLevelAssess.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *sLevelAssess) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneLevelAssessReq{LevelAssess: &v1.LevelAssessInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || info == nil {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	// 删除父级时，校验子维度是否为空
	if g.IsEmpty(info.LevelAssess.GetEvaluateId()) {
		zidInfo, err := s.GetOne(ctx, &v1.GetOneLevelAssessReq{LevelAssess: &v1.LevelAssessInfo{EvaluateId: id}})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, "当前数据不存在，请联系相关维护人员", err
		}
		fmt.Println("zidInfo-------------------", zidInfo.GetLevelAssess())
		if !g.IsNil(zidInfo.GetLevelAssess()) && !g.IsEmpty(zidInfo.LevelAssess.Id) {
			return false, "请先移除当前评价维度的下级信息", errors.New(fmt.Sprintf("当前评价维度存在子级信息ID：%d,评价标准:%s", zidInfo.LevelAssess.Id, zidInfo.LevelAssess.EvaluateCriteria))
		}
	}

	_, err = dao.ProductLevelAssess.Ctx(ctx).Where(dao.ProductLevelAssess.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sLevelAssess) GetWeightSum(ctx context.Context, condition g.Map) (float64, error) {
	query := dao.ProductLevelAssess.Ctx(ctx).Where(condition)

	return query.Sum(dao.ProductLevelAssess.Columns().Weight)
}

func (s *sLevelAssess) checkInputData(ctx context.Context, in *v1.LevelAssessInfo) (*v1.LevelAssessInfo, error) {
	if len(in.GetEvaluateDimensions()) == 0 {
		return in, errors.New("评价维度不能为空")
	}

	// 上级评价维度存在时，数据校验
	if in.GetEvaluateId() > 0 {
		if len(in.GetEvaluateCriteria()) == 0 {
			return in, errors.New("评价标准不能为空")
		}
		if len(in.GetScoreCriteria()) == 0 {
			return in, errors.New("评分标准不能为空")
		}
		if g.IsEmpty(in.GetWeight()) {
			return in, errors.New("权重不能为0")
		}

		pidInfo, err := s.GetOne(ctx, &v1.GetOneLevelAssessReq{
			LevelAssess: &v1.LevelAssessInfo{
				Id: in.GetEvaluateId(),
			},
		})
		if (err != nil && err == sql.ErrNoRows) || pidInfo == nil {
			return in, errors.New("选择的上级评价维度不存在，请确认是否增加")
		}
	}

	// 评价权重和不能大于1
	condition := g.Map{}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	weightSum, err := s.GetWeightSum(ctx, condition)
	if err != nil && err == sql.ErrNoRows {
		return in, err
	}

	if library.Decimal(weightSum+gconv.Float64(in.Weight)) > 1 {
		return in, errors.New("当前评价标准权重和超过1，请确认输入信息是否正确")
	}

	return in, nil
}
