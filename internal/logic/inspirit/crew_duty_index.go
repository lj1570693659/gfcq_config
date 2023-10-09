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
	"github.com/lj1570693659/gfcq_config/boot"
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
	"github.com/lj1570693659/gfcq_config/internal/model/do"
	"github.com/lj1570693659/gfcq_config/internal/model/entity"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	sCrewDutyIndex struct{}
)

func init() {
	service.RegisterCrewDutyIndex(&sCrewDutyIndex{})
}

func (s *sCrewDutyIndex) Create(ctx context.Context, in *v1.CreateCrewDutyIndexReq) (*v1.CreateCrewDutyIndexRes, error) {
	res := &v1.CreateCrewDutyIndexRes{
		CrewDutyIndex: &v1.CrewDutyIndexInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.CrewDutyIndexInfo{
		ScoreIndex: in.GetScoreIndex(),
		JobLevelId: in.GetJobLevelId(),
		Arith:      in.GetArith(),
		Remark:     in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.CrewDutyIndex{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.CrewDutyIndex.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.CrewDutyIndex.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sCrewDutyIndex) GetOne(ctx context.Context, in *v1.GetOneCrewDutyIndexReq) (*v1.GetOneCrewDutyIndexRes, error) {
	var info *v1.CrewDutyIndexInfo
	query := dao.CrewDutyIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewDutyIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().ScoreIndex, in.GetCrewDutyIndex().GetScoreIndex())
	}
	if in.GetCrewDutyIndex().GetJobLevelId() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().JobLevelId, in.GetCrewDutyIndex().GetJobLevelId())
	}
	if in.GetCrewDutyIndex().GetArith() != v1.ArithEnum_notSureArith {
		query = query.Where(dao.CrewDutyIndex.Columns().Arith, in.GetCrewDutyIndex().GetArith())
	}
	// 主键查询
	if in.GetCrewDutyIndex().GetId() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().Id, in.GetCrewDutyIndex().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneCrewDutyIndexRes{
		CrewDutyIndex: info,
	}, err
}

func (s *sCrewDutyIndex) GetList(ctx context.Context, in *v1.GetListCrewDutyIndexReq) (*v1.GetListCrewDutyIndexRes, error) {
	res := &v1.GetListCrewDutyIndexRes{}
	resData := make([]*v1.CrewDutyIndexInfo, 0)
	budgetEntity := make([]entity.CrewDutyIndex, 0)

	query := dao.CrewDutyIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewDutyIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().ScoreIndex, in.GetCrewDutyIndex().GetScoreIndex())
	}
	if in.GetCrewDutyIndex().GetJobLevelId() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().JobLevelId, in.GetCrewDutyIndex().GetJobLevelId())
	}
	if in.GetCrewDutyIndex().GetArith() != v1.ArithEnum_notSureArith {
		query = query.Where(dao.CrewDutyIndex.Columns().Arith, in.GetCrewDutyIndex().GetArith())
	}
	// 主键查询
	if in.GetCrewDutyIndex().GetId() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().Id, in.GetCrewDutyIndex().GetId())
	}
	// 主键
	if len(in.GetCrewDutyIndex().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewDutyIndex.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewDutyIndex().GetRemark(), "%")})
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

func (s *sCrewDutyIndex) GetAll(ctx context.Context, in *v1.GetAllCrewDutyIndexReq) (*v1.GetAllCrewDutyIndexRes, error) {
	res := &v1.GetAllCrewDutyIndexRes{}
	resData := make([]*v1.CrewDutyIndexInfo, 0)
	budgetEntity := make([]entity.CrewDutyIndex, 0)

	query := dao.CrewDutyIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewDutyIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().ScoreIndex, in.GetCrewDutyIndex().GetScoreIndex())
	}
	if in.GetCrewDutyIndex().GetJobLevelId() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().JobLevelId, in.GetCrewDutyIndex().GetJobLevelId())
	}
	if in.GetCrewDutyIndex().GetArith() != v1.ArithEnum_notSureArith {
		query = query.Where(dao.CrewDutyIndex.Columns().Arith, in.GetCrewDutyIndex().GetArith())
	}
	// 主键查询
	if in.GetCrewDutyIndex().GetId() > 0 {
		query = query.Where(dao.CrewDutyIndex.Columns().Id, in.GetCrewDutyIndex().GetId())
	}
	// 主键
	if len(in.GetCrewDutyIndex().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewDutyIndex.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewDutyIndex().GetRemark(), "%")})
	}

	err := query.OrderAsc(dao.CrewDutyIndex.Columns().ScoreIndex).Scan(&budgetEntity)
	if len(budgetEntity) > 0 {
		for _, v := range budgetEntity {
			resData = append(resData, &v1.CrewDutyIndexInfo{
				Id:         gconv.Int32(v.Id),
				ScoreIndex: gconv.Uint32(v.ScoreIndex),
				JobLevelId: gconv.Uint32(v.JobLevelId),
				Arith:      v1.ArithEnum(v1.ArithEnum_value[v.Arith]),
				Remark:     v.Remark,
				CreateTime: v.CreateTime.String(),
				UpdateTime: v.UpdateTime.String(),
			})
		}
	}
	res.Data = resData
	return res, err
}

func (s *sCrewDutyIndex) Modify(ctx context.Context, in *v1.ModifyCrewDutyIndexReq) (*v1.ModifyCrewDutyIndexRes, error) {
	res := &v1.ModifyCrewDutyIndexRes{CrewDutyIndex: &v1.CrewDutyIndexInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.CrewDutyIndexInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.CrewDutyIndex{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.CrewDutyIndex.Ctx(ctx).Where(dao.CrewDutyIndex.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.CrewDutyIndex = info
	return res, nil
}

func (s *sCrewDutyIndex) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneCrewDutyIndexReq{CrewDutyIndex: &v1.CrewDutyIndexInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.CrewDutyIndex.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.CrewDutyIndex.Ctx(ctx).Where(dao.CrewDutyIndex.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sCrewDutyIndex) checkInputData(ctx context.Context, in *v1.CrewDutyIndexInfo) (*v1.CrewDutyIndexInfo, error) {
	if in.GetScoreIndex() < 0 {
		return in, errors.New("职责指数不能小于0")
	}
	if g.IsEmpty(in.GetJobLevelId()) {
		return in, errors.New("职级不能为空")
	}

	// 职级校验
	jobLevel, err := boot.JobLevelServer.GetOne(ctx, &common.GetOneJobLevelReq{
		Id: gconv.Int32(in.GetJobLevelId()),
	})
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if g.IsNil(jobLevel) || g.IsEmpty(jobLevel) || (!g.IsNil(jobLevel) && g.IsEmpty(jobLevel.JobLevel.Id)) {
		return in, errors.New("当前职级信息不存在，请确认数据是否正确")
	}
	return in, nil
}
