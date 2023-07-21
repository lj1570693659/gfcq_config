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
	sCrewHoursIndex struct{}
)

func init() {
	service.RegisterCrewHoursIndex(&sCrewHoursIndex{})
}

func (s *sCrewHoursIndex) Create(ctx context.Context, in *v1.CreateCrewHoursIndexReq) (*v1.CreateCrewHoursIndexRes, error) {
	res := &v1.CreateCrewHoursIndexRes{
		CrewHoursIndex: &v1.CrewHoursIndexInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.CrewHoursIndexInfo{
		ScoreIndex: in.GetScoreIndex(),
		ScoreMin:   in.GetScoreMin(),
		ScoreMax:   in.GetScoreMax(),
		ScoreRange: in.GetScoreRange(),
		Remark:     in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.CrewHoursIndex{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.CrewHoursIndex.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.CrewHoursIndex.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sCrewHoursIndex) GetOne(ctx context.Context, in *v1.GetOneCrewHoursIndexReq) (*v1.GetOneCrewHoursIndexRes, error) {
	var info *v1.CrewHoursIndexInfo
	query := dao.CrewHoursIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewHoursIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreIndex, in.GetCrewHoursIndex().GetScoreIndex())
	}
	if in.GetCrewHoursIndex().GetScoreMin() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreMin, in.GetCrewHoursIndex().GetScoreMin())
	}
	if in.GetCrewHoursIndex().GetScoreMax() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreMax, in.GetCrewHoursIndex().GetScoreMax())
	}
	// 主键查询
	if in.GetCrewHoursIndex().GetId() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().Id, in.GetCrewHoursIndex().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneCrewHoursIndexRes{
		CrewHoursIndex: info,
	}, err
}

func (s *sCrewHoursIndex) GetList(ctx context.Context, in *v1.GetListCrewHoursIndexReq) (*v1.GetListCrewHoursIndexRes, error) {
	res := &v1.GetListCrewHoursIndexRes{
		Page: in.GetPage(),
		Size: in.GetSize(),
	}
	resData := make([]*v1.CrewHoursIndexInfo, 0)
	budgetEntity := make([]entity.CrewHoursIndex, 0)

	query := dao.CrewHoursIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewHoursIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreIndex, in.GetCrewHoursIndex().GetScoreIndex())
	}
	if in.GetCrewHoursIndex().GetScoreMin() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreMin, in.GetCrewHoursIndex().GetScoreMin())
	}
	if in.GetCrewHoursIndex().GetScoreMax() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreMax, in.GetCrewHoursIndex().GetScoreMax())
	}
	// 主键查询
	if in.GetCrewHoursIndex().GetId() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().Id, in.GetCrewHoursIndex().GetId())
	}
	// 备注
	if len(in.GetCrewHoursIndex().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewHoursIndex.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewHoursIndex().GetRemark(), "%")})
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

func (s *sCrewHoursIndex) GetAll(ctx context.Context, in *v1.GetAllCrewHoursIndexReq) (*v1.GetAllCrewHoursIndexRes, error) {
	res := &v1.GetAllCrewHoursIndexRes{}
	resData := make([]*v1.CrewHoursIndexInfo, 0)
	budgetEntity := make([]entity.CrewHoursIndex, 0)

	query := dao.CrewHoursIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewHoursIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreIndex, in.GetCrewHoursIndex().GetScoreIndex())
	}
	if in.GetCrewHoursIndex().GetScoreMin() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreMin, in.GetCrewHoursIndex().GetScoreMin())
	}
	if in.GetCrewHoursIndex().GetScoreMax() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().ScoreMax, in.GetCrewHoursIndex().GetScoreMax())
	}
	// 主键查询
	if in.GetCrewHoursIndex().GetId() > 0 {
		query = query.Where(dao.CrewHoursIndex.Columns().Id, in.GetCrewHoursIndex().GetId())
	}
	// 备注
	if len(in.GetCrewHoursIndex().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewHoursIndex.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewHoursIndex().GetRemark(), "%")})
	}

	err := query.Scan(&budgetEntity)

	levelEntityByte, _ := json.Marshal(budgetEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Data = resData
	return res, err
}

func (s *sCrewHoursIndex) Modify(ctx context.Context, in *v1.ModifyCrewHoursIndexReq) (*v1.ModifyCrewHoursIndexRes, error) {
	res := &v1.ModifyCrewHoursIndexRes{CrewHoursIndex: &v1.CrewHoursIndexInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.CrewHoursIndexInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.CrewHoursIndex{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.CrewHoursIndex.Ctx(ctx).Where(dao.CrewHoursIndex.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.CrewHoursIndex = info
	return res, nil
}

func (s *sCrewHoursIndex) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneCrewHoursIndexReq{CrewHoursIndex: &v1.CrewHoursIndexInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.CrewHoursIndex.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.CrewHoursIndex.Ctx(ctx).Where(dao.CrewHoursIndex.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sCrewHoursIndex) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.CrewHoursIndex, error) {
	info := &entity.CrewHoursIndex{}
	err := dao.CrewHoursIndex.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sCrewHoursIndex) checkInputData(ctx context.Context, in *v1.CrewHoursIndexInfo) (*v1.CrewHoursIndexInfo, error) {
	if in.GetScoreIndex() < 0 {
		return in, errors.New("工时指数不能小于0")
	}
	min := library.Decimal(gconv.Float64(in.GetScoreMin()))
	max := library.Decimal(gconv.Float64(in.GetScoreMax()))
	if min < 0 {
		return in, errors.New("工时比例下限不能小于0")
	}
	if max > 1 {
		return in, errors.New("工时比例上限不能大于1")
	}
	if min > max {
		return in, errors.New("工时比例下限不能超过上限")
	}

	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.CrewHoursIndex.Columns().ScoreIndex): in.GetScoreIndex(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前工时指数已存在，请确认输入信息是否正确")
	}

	return in, nil
}
