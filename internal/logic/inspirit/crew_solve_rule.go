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
	sCrewSolveRule struct{}
)

func init() {
	service.RegisterCrewSolveRule(&sCrewSolveRule{})
}

func (s *sCrewSolveRule) Create(ctx context.Context, in *v1.CreateCrewSolveRuleReq) (*v1.CreateCrewSolveRuleRes, error) {
	res := &v1.CreateCrewSolveRuleRes{
		CrewSolveRule: &v1.CrewSolveRuleInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.CrewSolveRuleInfo{
		Redio:  in.GetRedio(),
		Demand: in.GetDemand(),
		Remark: in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.CrewSolveRule{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.CrewSolveRule.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.CrewSolveRule.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sCrewSolveRule) GetOne(ctx context.Context, in *v1.GetOneCrewSolveRuleReq) (*v1.GetOneCrewSolveRuleRes, error) {
	var info *v1.CrewSolveRuleInfo
	query := dao.CrewSolveRule.Ctx(ctx)

	// 浮动比例
	if in.GetCrewSolveRule().GetRedio() > 0 {
		query = query.Where(dao.CrewSolveRule.Columns().Redio, in.GetCrewSolveRule().GetRedio())
	}

	if value, ok := v1.DemandEnum_value[in.GetCrewSolveRule().GetDemand().String()]; ok {
		query = query.Where(dao.CrewSolveRule.Columns().Demand, value)
	}
	// 主键查询
	if in.GetCrewSolveRule().GetId() > 0 {
		query = query.Where(dao.CrewSolveRule.Columns().Id, in.GetCrewSolveRule().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneCrewSolveRuleRes{
		CrewSolveRule: info,
	}, err
}

func (s *sCrewSolveRule) GetList(ctx context.Context, in *v1.GetListCrewSolveRuleReq) (*v1.GetListCrewSolveRuleRes, error) {
	res := &v1.GetListCrewSolveRuleRes{
		Page: in.GetPage(),
		Size: in.GetSize(),
	}
	resData := make([]*v1.CrewSolveRuleInfo, 0)
	budgetEntity := make([]entity.CrewSolveRule, 0)

	query := dao.CrewSolveRule.Ctx(ctx)

	// 评价标准
	if in.GetCrewSolveRule().GetRedio() > 0 {
		query = query.Where(dao.CrewSolveRule.Columns().Redio, in.GetCrewSolveRule().GetRedio())
	}
	if in.GetCrewSolveRule().GetDemand() > 0 {
		query = query.Where(dao.CrewSolveRule.Columns().Demand, in.GetCrewSolveRule().GetDemand())
	}
	// 主键查询
	if in.GetCrewSolveRule().GetId() > 0 {
		query = query.Where(dao.CrewSolveRule.Columns().Id, in.GetCrewSolveRule().GetId())
	}
	// 备注
	if len(in.GetCrewSolveRule().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewSolveRule.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewSolveRule().GetRemark(), "%")})
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

func (s *sCrewSolveRule) Modify(ctx context.Context, in *v1.ModifyCrewSolveRuleReq) (*v1.ModifyCrewSolveRuleRes, error) {
	res := &v1.ModifyCrewSolveRuleRes{CrewSolveRule: &v1.CrewSolveRuleInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.CrewSolveRuleInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.CrewSolveRule{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.CrewSolveRule.Ctx(ctx).Where(dao.CrewSolveRule.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.CrewSolveRule = info
	return res, nil
}

func (s *sCrewSolveRule) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneCrewSolveRuleReq{CrewSolveRule: &v1.CrewSolveRuleInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.CrewSolveRule.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.CrewSolveRule.Ctx(ctx).Where(dao.CrewSolveRule.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sCrewSolveRule) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.CrewSolveRule, error) {
	info := &entity.CrewSolveRule{}
	err := dao.CrewSolveRule.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sCrewSolveRule) checkInputData(ctx context.Context, in *v1.CrewSolveRuleInfo) (*v1.CrewSolveRuleInfo, error) {
	if in.GetRedio() < 0 {
		return in, errors.New("浮动比例不能小于0")
	}
	if in.GetRedio() > 1 {
		return in, errors.New("浮动比例不能大于1")
	}

	if in.GetDemand() != v1.DemandEnum_highlight && in.GetDemand() != v1.DemandEnum_middle && in.GetDemand() != v1.DemandEnum_less {
		return in, errors.New("贡献维度下限不能超过上限")
	}

	return in, nil
}
