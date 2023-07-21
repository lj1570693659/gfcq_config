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
	sMode struct{}
)

func init() {
	service.RegisterMode(&sMode{})
}

func (s *sMode) Create(ctx context.Context, in *v1.CreateModeReq) (*v1.CreateModeRes, error) {
	res := &v1.CreateModeRes{
		Mode: &v1.ModeInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.ModeInfo{
		Name:   in.GetName(),
		Factor: in.GetFactor(),
		Remark: in.Remark,
	}); err != nil {
		return res, err
	}

	data := do.ProductMode{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.ProductMode.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.Mode.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sMode) GetOne(ctx context.Context, in *v1.GetOneModeReq) (*v1.GetOneModeRes, error) {
	var info *v1.ModeInfo
	query := dao.ProductMode.Ctx(ctx)

	// 优先级
	if len(in.GetMode().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductMode.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetMode().GetName(), "%")})
	}
	// 主键查询
	if in.GetMode().GetId() > 0 {
		query = query.Where(dao.ProductMode.Columns().Id, in.GetMode().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneModeRes{
		Mode: info,
	}, err
}

func (s *sMode) GetList(ctx context.Context, in *v1.GetListModeReq) (*v1.GetListModeRes, error) {
	res := &v1.GetListModeRes{
		Page:      in.GetPage(),
		Size:      in.GetSize(),
		TotalSize: 0,
		Data:      make([]*v1.ModeInfo, 0),
	}
	resData := make([]*v1.ModeInfo, 0)
	confirmEntity := make([]entity.ProductMode, 0)

	query := dao.ProductMode.Ctx(ctx)

	// 优先级
	if len(in.GetMode().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductMode.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetMode().GetName(), "%")})
	}
	// 主键查询
	if in.GetMode().GetId() > 0 {
		query = query.Where(dao.ProductMode.Columns().Id, in.GetMode().GetId())
	}

	// 主键
	if len(in.GetMode().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductMode.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetMode().GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, in.GetPage(), in.GetSize())
	if err != nil {
		return res, err
	}
	err = query.Scan(&confirmEntity)

	levelEntityByte, _ := json.Marshal(confirmEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sMode) GetAll(ctx context.Context, in *v1.GetAllModeReq) (*v1.GetAllModeRes, error) {
	resData := make([]*v1.ModeInfo, 0)
	confirmEntity := make([]entity.ProductMode, 0)

	query := dao.ProductMode.Ctx(ctx)

	// 优先级
	if len(in.GetMode().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductMode.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetMode().GetName(), "%")})
	}
	// 主键查询
	if in.GetMode().GetId() > 0 {
		query = query.Where(dao.ProductMode.Columns().Id, in.GetMode().GetId())
	}

	// 主键
	if len(in.GetMode().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductMode.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetMode().GetRemark(), "%")})
	}

	err := query.Scan(&confirmEntity)

	levelEntityByte, _ := json.Marshal(confirmEntity)
	json.Unmarshal(levelEntityByte, &resData)

	return &v1.GetAllModeRes{
		Data: resData,
	}, err
}

func (s *sMode) Modify(ctx context.Context, in *v1.ModifyModeReq) (*v1.ModifyModeRes, error) {
	res := &v1.ModifyModeRes{Mode: &v1.ModeInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.ModeInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductMode{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductMode.Ctx(ctx).Where(dao.ProductMode.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.Mode = info
	return res, nil
}

func (s *sMode) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneModeReq{Mode: &v1.ModeInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || info == nil {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.ProductMode.Ctx(ctx).Where(dao.ProductMode.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sMode) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.ProductMode, error) {
	info := &entity.ProductMode{}
	err := dao.ProductMode.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sMode) checkInputData(ctx context.Context, in *v1.ModeInfo) (*v1.ModeInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("研发模式不能为空")
	}

	if in.GetFactor() < 0 || in.GetFactor() > 1 {
		return in, errors.New("请输入值区间位于[0 ~ 1]的开发系数")
	}

	// 评价权重和不能大于1
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductMode.Columns().Name): in.GetName(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前研发模式已存在，请确认输入信息是否正确")
	}

	return in, nil
}
