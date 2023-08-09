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
	sCrewKpiRule struct{}
)

func init() {
	service.RegisterCrewKpiRule(&sCrewKpiRule{})
}

func (s *sCrewKpiRule) Create(ctx context.Context, in *v1.CreateCrewKpiRuleReq) (*v1.CreateCrewKpiRuleRes, error) {
	res := &v1.CreateCrewKpiRuleRes{
		CrewKpiRule: &v1.CrewKpiRuleInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.CrewKpiRuleInfo{
		ScoreMin:   in.GetScoreMin(),
		ScoreMax:   in.GetScoreMax(),
		ScoreRange: in.GetScoreRange(),
		Redio:      in.GetRedio(),
		LevelName:  in.GetLevelName(),
		Remark:     in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.CrewKpiRule{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.CrewKpiRule.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.CrewKpiRule.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sCrewKpiRule) GetOne(ctx context.Context, in *v1.GetOneCrewKpiRuleReq) (*v1.GetOneCrewKpiRuleRes, error) {
	var info *v1.CrewKpiRuleInfo
	query := dao.CrewKpiRule.Ctx(ctx)

	// 浮动比例
	if in.GetCrewKpiRule().GetRedio() > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().Redio, in.GetCrewKpiRule().GetRedio())
	}

	// -1 全部
	if len(in.GetCrewKpiRule().GetLevelName()) > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().LevelName, in.GetCrewKpiRule().GetLevelName())
	}

	// 主键查询
	if in.GetCrewKpiRule().GetId() > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().Id, in.GetCrewKpiRule().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneCrewKpiRuleRes{
		CrewKpiRule: info,
	}, err
}

func (s *sCrewKpiRule) GetList(ctx context.Context, in *v1.GetListCrewKpiRuleReq) (*v1.GetListCrewKpiRuleRes, error) {
	res := &v1.GetListCrewKpiRuleRes{
		Page: in.GetPage(),
		Size: in.GetSize(),
	}
	resData := make([]*v1.CrewKpiRuleInfo, 0)
	budgetEntity := make([]entity.CrewKpiRule, 0)

	query := dao.CrewKpiRule.Ctx(ctx)

	// 评价标准
	if in.GetCrewKpiRule().GetRedio() > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().Redio, in.GetCrewKpiRule().GetRedio())
	}
	if len(in.GetCrewKpiRule().GetLevelName()) > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().LevelName, in.GetCrewKpiRule().GetLevelName())
	}
	// 主键查询
	if in.GetCrewKpiRule().GetId() > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().Id, in.GetCrewKpiRule().GetId())
	}
	// 备注
	if len(in.GetCrewKpiRule().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewKpiRule.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewKpiRule().GetRemark(), "%")})
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

func (s *sCrewKpiRule) GetAll(ctx context.Context, in *v1.GetAllCrewKpiRuleReq) (*v1.GetAllCrewKpiRuleRes, error) {
	res := &v1.GetAllCrewKpiRuleRes{}
	resData := make([]*v1.CrewKpiRuleInfo, 0)
	budgetEntity := make([]entity.CrewKpiRule, 0)

	query := dao.CrewKpiRule.Ctx(ctx)

	// 评价标准
	if in.GetCrewKpiRule().GetRedio() > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().Redio, in.GetCrewKpiRule().GetRedio())
	}
	if len(in.GetCrewKpiRule().GetLevelName()) > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().LevelName, in.GetCrewKpiRule().GetLevelName())
	}
	// 主键查询
	if in.GetCrewKpiRule().GetId() > 0 {
		query = query.Where(dao.CrewKpiRule.Columns().Id, in.GetCrewKpiRule().GetId())
	}
	// 备注
	if len(in.GetCrewKpiRule().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewKpiRule.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewKpiRule().GetRemark(), "%")})
	}

	err := query.Scan(&budgetEntity)

	levelEntityByte, _ := json.Marshal(budgetEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Data = resData
	return res, err
}

func (s *sCrewKpiRule) Modify(ctx context.Context, in *v1.ModifyCrewKpiRuleReq) (*v1.ModifyCrewKpiRuleRes, error) {
	res := &v1.ModifyCrewKpiRuleRes{CrewKpiRule: &v1.CrewKpiRuleInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.CrewKpiRuleInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.CrewKpiRule{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.CrewKpiRule.Ctx(ctx).Where(dao.CrewKpiRule.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.CrewKpiRule = info
	return res, nil
}

func (s *sCrewKpiRule) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneCrewKpiRuleReq{CrewKpiRule: &v1.CrewKpiRuleInfo{Id: id}})
	if (err != nil && err.Error() == sql.ErrNoRows.Error()) || (!g.IsNil(info) && g.IsNil(info.CrewKpiRule)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.CrewKpiRule.Ctx(ctx).Where(dao.CrewKpiRule.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sCrewKpiRule) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.CrewKpiRule, error) {
	info := &entity.CrewKpiRule{}
	err := dao.CrewKpiRule.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sCrewKpiRule) checkInputData(ctx context.Context, in *v1.CrewKpiRuleInfo) (*v1.CrewKpiRuleInfo, error) {
	if in.GetRedio() < 0 {
		return in, errors.New("绩效比例不能小于0")
	}
	if in.GetScoreMin() < 0 {
		return in, errors.New("绩效得分下限不能小于0")
	}
	if in.GetScoreMax() < 1 {
		return in, errors.New("绩效得分上限不能小于1")
	}
	if in.GetScoreMin() > in.GetScoreMax() {
		return in, errors.New("绩效得分下限不能超过上限")
	}
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.CrewKpiRule.Columns().LevelName): in.GetLevelName(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前绩效等级已存在，请确认输入信息是否正确")
	}

	return in, nil
}
