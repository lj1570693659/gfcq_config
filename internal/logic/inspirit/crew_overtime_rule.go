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
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
	"github.com/lj1570693659/gfcq_config/internal/model/do"
	"github.com/lj1570693659/gfcq_config/internal/model/entity"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	sCrewOvertimeRule struct{}
)

func init() {
	service.RegisterCrewOvertimeRule(&sCrewOvertimeRule{})
}

func (s *sCrewOvertimeRule) Create(ctx context.Context, in *v1.CreateCrewOvertimeRuleReq) (*v1.CreateCrewOvertimeRuleRes, error) {
	res := &v1.CreateCrewOvertimeRuleRes{
		CrewOvertimeRule: &v1.CrewOvertimeRuleInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.CrewOvertimeRuleInfo{
		Redio:      in.GetRedio(),
		ScoreMin:   in.GetScoreMin(),
		ScoreMax:   in.GetScoreMax(),
		ScoreRange: in.GetScoreRange(),
		Remark:     in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.CrewOvertimeRule{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.CrewOvertimeRule.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.CrewOvertimeRule.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sCrewOvertimeRule) GetOne(ctx context.Context, in *v1.GetOneCrewOvertimeRuleReq) (*v1.GetOneCrewOvertimeRuleRes, error) {
	var info *v1.CrewOvertimeRuleInfo
	query := dao.CrewOvertimeRule.Ctx(ctx)

	// 浮动比例
	if in.GetCrewOvertimeRule().GetRedio() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().Redio, in.GetCrewOvertimeRule().GetRedio())
	}
	if in.GetCrewOvertimeRule().GetScoreMin() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().ScoreMin, in.GetCrewOvertimeRule().GetScoreMin())
	}
	if in.GetCrewOvertimeRule().GetScoreMax() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().ScoreMax, in.GetCrewOvertimeRule().GetScoreMax())
	}
	// 主键查询
	if in.GetCrewOvertimeRule().GetId() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().Id, in.GetCrewOvertimeRule().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneCrewOvertimeRuleRes{
		CrewOvertimeRule: info,
	}, err
}

func (s *sCrewOvertimeRule) GetList(ctx context.Context, in *v1.GetListCrewOvertimeRuleReq) (*v1.GetListCrewOvertimeRuleRes, error) {
	res := &v1.GetListCrewOvertimeRuleRes{
		Page: in.GetPage(),
		Size: in.GetSize(),
	}
	resData := make([]*v1.CrewOvertimeRuleInfo, 0)
	budgetEntity := make([]entity.CrewOvertimeRule, 0)

	query := dao.CrewOvertimeRule.Ctx(ctx)

	// 评价标准
	if in.GetCrewOvertimeRule().GetRedio() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().Redio, in.GetCrewOvertimeRule().GetRedio())
	}
	if in.GetCrewOvertimeRule().GetScoreMin() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().ScoreMin, in.GetCrewOvertimeRule().GetScoreMin())
	}
	if in.GetCrewOvertimeRule().GetScoreMax() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().ScoreMax, in.GetCrewOvertimeRule().GetScoreMax())
	}
	// 主键查询
	if in.GetCrewOvertimeRule().GetId() > 0 {
		query = query.Where(dao.CrewOvertimeRule.Columns().Id, in.GetCrewOvertimeRule().GetId())
	}
	// 备注
	if len(in.GetCrewOvertimeRule().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewOvertimeRule.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewOvertimeRule().GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, in.GetPage(), in.GetSize())
	if err != nil {
		return res, err
	}
	err = query.Scan(&budgetEntity)

	levelEntityByte, _ := json.Marshal(budgetEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sCrewOvertimeRule) Modify(ctx context.Context, in *v1.ModifyCrewOvertimeRuleReq) (*v1.ModifyCrewOvertimeRuleRes, error) {
	res := &v1.ModifyCrewOvertimeRuleRes{CrewOvertimeRule: &v1.CrewOvertimeRuleInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.CrewOvertimeRuleInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.CrewOvertimeRule{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.CrewOvertimeRule.Ctx(ctx).Where(dao.CrewOvertimeRule.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.CrewOvertimeRule = info
	return res, nil
}

func (s *sCrewOvertimeRule) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneCrewOvertimeRuleReq{CrewOvertimeRule: &v1.CrewOvertimeRuleInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.CrewOvertimeRule.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.CrewOvertimeRule.Ctx(ctx).Where(dao.CrewOvertimeRule.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sCrewOvertimeRule) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.CrewOvertimeRule, error) {
	info := &entity.CrewOvertimeRule{}
	err := dao.CrewOvertimeRule.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sCrewOvertimeRule) checkInputData(ctx context.Context, in *v1.CrewOvertimeRuleInfo) (*v1.CrewOvertimeRuleInfo, error) {
	if in.GetRedio() < 0 {
		return in, errors.New("浮动比例不能小于0")
	}
	if in.GetRedio() > 1 {
		return in, errors.New("浮动比例不能大于1")
	}
	min := library.Decimal(gconv.Float64(in.GetScoreMin()))
	max := library.Decimal(gconv.Float64(in.GetScoreMax()))
	if min < 0 {
		return in, errors.New("贡献维度下限不能小于0")
	}
	if max > 1 {
		return in, errors.New("贡献维度上限不能大于1")
	}
	if min > max {
		return in, errors.New("贡献维度下限不能超过上限")
	}

	return in, nil
}
