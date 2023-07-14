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
	service "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	sModeStage struct{}
)

func init() {
	service.RegisterModeStage(&sModeStage{})
}

func (s *sModeStage) Create(ctx context.Context, in *v1.CreateModeStageReq) (*v1.CreateModeStageRes, error) {
	res := &v1.CreateModeStageRes{
		ModeStage: &v1.ModeStageInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.ModeStageInfo{
		Name:       in.GetName(),
		Tid:        in.GetTid(),
		QuotaRadio: in.GetQuotaRadio(),
	}); err != nil {
		return res, err
	}

	data := do.ProductModeStage{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.ProductModeStage.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.ModeStage.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sModeStage) GetOne(ctx context.Context, in *v1.GetOneModeStageReq) (*v1.GetOneModeStageRes, error) {
	var info *v1.ModeStageInfo
	query := dao.ProductModeStage.Ctx(ctx)

	// 优先级
	if len(in.GetModeStage().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductModeStage.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetModeStage().GetName(), "%")})
	}
	// 主键查询
	if in.GetModeStage().GetId() > 0 {
		query = query.Where(dao.ProductModeStage.Columns().Id, in.GetModeStage().GetId())
	}
	if in.GetModeStage().GetTid() > 0 {
		query = query.Where(dao.ProductModeStage.Columns().Tid, in.GetModeStage().GetTid())
	}

	err := query.Scan(&info)

	return &v1.GetOneModeStageRes{
		ModeStage: info,
	}, err
}

func (s *sModeStage) GetList(ctx context.Context, in *v1.GetListModeStageReq) (*v1.GetListModeStageRes, error) {
	res := &v1.GetListModeStageRes{
		Page:      in.GetPage(),
		Size:      in.GetSize(),
		TotalSize: 0,
		Data:      make([]*v1.ModeStageInfo, 0),
	}
	resData := make([]*v1.ModeStageInfo, 0)
	stageEntity := make([]entity.ProductModeStage, 0)

	query := dao.ProductModeStage.Ctx(ctx)

	// 优先级
	if len(in.GetModeStage().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductModeStage.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetModeStage().GetName(), "%")})
	}
	// 主键查询
	if in.GetModeStage().GetId() > 0 {
		query = query.Where(dao.ProductModeStage.Columns().Id, in.GetModeStage().GetId())
	}
	if in.GetModeStage().GetTid() > 0 {
		query = query.Where(dao.ProductModeStage.Columns().Tid, in.GetModeStage().GetTid())
	}

	// 主键
	if len(in.GetModeStage().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductModeStage.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetModeStage().GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, in.GetPage(), in.GetSize())
	if err != nil {
		return res, err
	}
	err = query.Scan(&stageEntity)

	levelEntityByte, _ := json.Marshal(stageEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sModeStage) Modify(ctx context.Context, in *v1.ModifyModeStageReq) (*v1.ModifyModeStageRes, error) {
	res := &v1.ModifyModeStageRes{ModeStage: &v1.ModeStageInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.ModeStageInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductModeStage{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductModeStage.Ctx(ctx).Where(dao.ProductModeStage.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.ModeStage = info
	return res, nil
}

func (s *sModeStage) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneModeStageReq{ModeStage: &v1.ModeStageInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.ModeStage.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.ProductModeStage.Ctx(ctx).Where(dao.ProductModeStage.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sModeStage) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.ProductModeStage, error) {
	info := &entity.ProductModeStage{}
	err := dao.ProductModeStage.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sModeStage) GetRadioSumByCondition(ctx context.Context, condition g.Map) (float64, error) {
	return dao.ProductModeStage.Ctx(ctx).Where(condition).Sum(dao.ProductModeStage.Columns().QuotaRadio)
}

func (s *sModeStage) checkInputData(ctx context.Context, in *v1.ModeStageInfo) (*v1.ModeStageInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("研发模式不能为空")
	}

	if in.GetTid() == 0 {
		return in, errors.New("对应研发模式不能为空")
	}

	if in.GetQuotaRadio() < 0 || in.GetQuotaRadio() > 1 {
		return in, errors.New("请输入值区间位于[0 ~ 1]的额度占比")
	}

	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductModeStage.Columns().Tid): in.GetTid(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}

	// 同一研发模式下，阶段额度占比和不能大于1
	sum, err := s.GetRadioSumByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if library.Decimal(sum+gconv.Float64(in.GetQuotaRadio())) > 1 {
		return in, errors.New("阶段额度占比和不能超过1，请确认输入信息是否正确")
	}

	// 同一研发模式下阶段名称不能重复
	condition[fmt.Sprintf("%s = ?", dao.ProductModeStage.Columns().Name)] = in.GetName()
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前研发模式下阶段已存在，请确认输入信息是否正确")
	}

	return in, nil
}
