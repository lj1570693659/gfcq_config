package inspirit

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lj1570693659/gfcq_config/internal/consts"
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
	"github.com/lj1570693659/gfcq_config/internal/model/do"
	"github.com/lj1570693659/gfcq_config/internal/model/entity"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	sBudgetAccess struct{}
)

func init() {
	service.RegisterBudgetAccess(&sBudgetAccess{})
}

func (s *sBudgetAccess) Create(ctx context.Context, in *v1.CreateBudgetAssessReq) (*v1.CreateBudgetAssessRes, error) {
	res := &v1.CreateBudgetAssessRes{
		BudgetAssess: &v1.BudgetAssessInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.BudgetAssessInfo{
		ScoreMin:    in.GetScoreMin(),
		ScoreMax:    in.GetScoreMax(),
		ScoreRange:  in.GetScoreRange(),
		BudgetMin:   in.GetBudgetMin(),
		BudgetMax:   in.GetBudgetMax(),
		BudgetRange: in.GetBudgetRange(),
		Remark:      in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.ProductBudgetAccess{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()

	lastInsertId, err := dao.ProductBudgetAccess.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.BudgetAssess.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sBudgetAccess) GetOne(ctx context.Context, in *v1.GetOneBudgetAssessReq) (*v1.GetOneBudgetAssessRes, error) {
	var info *v1.BudgetAssessInfo
	query := dao.ProductBudgetAccess.Ctx(ctx)

	// 评价标准
	if in.GetBudgetAssess().GetScoreMin() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().ScoreMin, in.GetBudgetAssess().GetScoreMin())
	}
	if in.GetBudgetAssess().GetScoreMax() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().ScoreMax, in.GetBudgetAssess().GetScoreMax())
	}
	if in.GetBudgetAssess().GetBudgetMin() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().BudgetMin, in.GetBudgetAssess().GetBudgetMin())
	}
	if in.GetBudgetAssess().GetBudgetMax() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().BudgetMax, in.GetBudgetAssess().GetBudgetMax())
	}
	// 主键查询
	if in.GetBudgetAssess().GetId() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().Id, in.GetBudgetAssess().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneBudgetAssessRes{
		BudgetAssess: info,
	}, err
}

func (s *sBudgetAccess) GetList(ctx context.Context, in *v1.GetListBudgetAssessReq) (*v1.GetListBudgetAssessRes, error) {
	res := &v1.GetListBudgetAssessRes{}
	resData := make([]*v1.BudgetAssessInfo, 0)
	budgetEntity := make([]entity.ProductBudgetAccess, 0)

	query := dao.ProductBudgetAccess.Ctx(ctx)

	// 评价标准
	if in.GetBudgetAssess().GetScoreMin() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().ScoreMin, in.GetBudgetAssess().GetScoreMin())
	}
	if in.GetBudgetAssess().GetScoreMax() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().ScoreMax, in.GetBudgetAssess().GetScoreMax())
	}
	if in.GetBudgetAssess().GetBudgetMin() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().BudgetMin, in.GetBudgetAssess().GetBudgetMin())
	}
	if in.GetBudgetAssess().GetBudgetMax() > 0 {
		query = query.Where(dao.ProductBudgetAccess.Columns().BudgetMax, in.GetBudgetAssess().GetBudgetMax())
	}
	// 主键
	if len(in.GetBudgetAssess().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductBudgetAccess.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetBudgetAssess().GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, in.GetPage(), in.GetSize())
	if err != nil {
		return res, err
	}
	err = query.Scan(&budgetEntity)

	levelEntityByte, _ := json.Marshal(budgetEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Page = in.GetPage()
	res.Size = in.GetSize()
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sBudgetAccess) Modify(ctx context.Context, in *v1.ModifyBudgetAssessReq) (*v1.ModifyBudgetAssessRes, error) {
	res := &v1.ModifyBudgetAssessRes{BudgetAssess: &v1.BudgetAssessInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.BudgetAssessInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductBudgetAccess{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductBudgetAccess.Ctx(ctx).Where(dao.ProductBudgetAccess.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.BudgetAssess = info
	return res, nil
}

func (s *sBudgetAccess) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneBudgetAssessReq{BudgetAssess: &v1.BudgetAssessInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.BudgetAssess.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.ProductBudgetAccess.Ctx(ctx).Where(dao.ProductBudgetAccess.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sBudgetAccess) checkInputData(ctx context.Context, in *v1.BudgetAssessInfo) (*v1.BudgetAssessInfo, error) {
	if in.GetScoreMin() < 0 {
		return in, errors.New("评分下限不能小于0")
	}
	if in.GetScoreMax() < 0 || in.GetScoreMax() > 100 {
		return in, errors.New("评分上限不能小于0，且不能大于100")
	}
	if in.GetBudgetMin() < 0 || in.GetBudgetMax() < 0 {
		return in, errors.New("激励预算不能小于0")
	}
	if in.GetScoreRange() != consts.ScoreRangeMin && in.GetScoreRange() != consts.ScoreRangeMax {
		return in, errors.New("评分区间包含关系错误")
	}
	if in.GetBudgetRange() != consts.ScoreRangeMin && in.GetBudgetRange() != consts.ScoreRangeMax {
		return in, errors.New("激励预算区间包含关系错误")
	}

	return in, nil
}
