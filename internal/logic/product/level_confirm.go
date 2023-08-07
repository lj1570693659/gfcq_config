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
	"github.com/lj1570693659/gfcq_config/internal/consts"
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
	"github.com/lj1570693659/gfcq_config/internal/model/do"
	"github.com/lj1570693659/gfcq_config/internal/model/entity"
	productService "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	sLevelConfirm struct{}
)

func init() {
	productService.RegisterLevelConfirm(&sLevelConfirm{})
}

func (s *sLevelConfirm) Create(ctx context.Context, in *v1.CreateLevelConfirmReq) (*v1.CreateLevelConfirmRes, error) {
	res := &v1.CreateLevelConfirmRes{
		LevelConfirm: &v1.LevelConfirmInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.LevelConfirmInfo{
		Name:          in.GetName(),
		ScoreMin:      in.GetScoreMin(),
		ScoreMax:      in.GetScoreMax(),
		ScoreRange:    in.GetScoreRange(),
		IsNeedPm:      in.GetIsNeedPm(),
		PmDemand:      in.GetPmDemand(),
		ProductDemand: in.GetProductDemand(),
		MonitDemand:   in.GetMonitDemand(),
		IsNeedPml:     in.GetIsNeedPml(),
		Remark:        in.Remark,
	}); err != nil {
		return res, err
	}

	data := do.ProductLevelConfirm{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.ProductLevelConfirm.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.LevelConfirm.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sLevelConfirm) GetOne(ctx context.Context, in *v1.GetOneLevelConfirmReq) (*v1.GetOneLevelConfirmRes, error) {
	var info *v1.LevelConfirmInfo
	query := dao.ProductLevelConfirm.Ctx(ctx)

	// 优先级
	if len(in.GetLevelConfirm().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelConfirm.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetLevelConfirm().GetName(), "%")})
	}
	// 主键查询
	if in.GetLevelConfirm().GetId() > 0 {
		query = query.Where(dao.ProductLevelConfirm.Columns().Id, in.GetLevelConfirm().GetId())
	}
	if in.GetLevelConfirm().GetIsNeedPm() > 0 {
		query = query.Where(dao.ProductLevelConfirm.Columns().IsNeedPm, in.GetLevelConfirm().GetIsNeedPm())
	}
	if in.GetLevelConfirm().GetIsNeedPml() > 0 {
		query = query.Where(dao.ProductLevelConfirm.Columns().IsNeedPml, in.GetLevelConfirm().GetIsNeedPml())
	}

	err := query.Scan(&info)

	return &v1.GetOneLevelConfirmRes{
		LevelConfirm: info,
	}, err
}

func (s *sLevelConfirm) GetList(ctx context.Context, in *v1.GetListLevelConfirmReq) (*v1.GetListLevelConfirmRes, error) {
	res := &v1.GetListLevelConfirmRes{
		Page:      in.GetPage(),
		Size:      in.GetSize(),
		TotalSize: 0,
		Data:      make([]*v1.LevelConfirmInfo, 0),
	}
	resData := make([]*v1.LevelConfirmInfo, 0)
	confirmEntity := make([]entity.ProductLevelConfirm, 0)

	query := dao.ProductLevelConfirm.Ctx(ctx)

	// 优先级
	if len(in.GetLevelConfirm().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelConfirm.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetLevelConfirm().GetName(), "%")})
	}
	// 主键查询
	if in.GetLevelConfirm().GetId() > 0 {
		query = query.Where(dao.ProductLevelConfirm.Columns().Id, in.GetLevelConfirm().GetId())
	}
	if in.GetLevelConfirm().GetIsNeedPm() > 0 {
		query = query.Where(dao.ProductLevelConfirm.Columns().IsNeedPm, in.GetLevelConfirm().GetIsNeedPm())
	}
	if in.GetLevelConfirm().GetIsNeedPml() > 0 {
		query = query.Where(dao.ProductLevelConfirm.Columns().IsNeedPml, in.GetLevelConfirm().GetIsNeedPml())
	}
	// 主键
	if len(in.GetLevelConfirm().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductLevelConfirm.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetLevelConfirm().GetRemark(), "%")})
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

func (s *sLevelConfirm) Modify(ctx context.Context, in *v1.ModifyLevelConfirmReq) (*v1.ModifyLevelConfirmRes, error) {
	res := &v1.ModifyLevelConfirmRes{LevelConfirm: &v1.LevelConfirmInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	info := &v1.LevelConfirmInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductLevelConfirm{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductLevelConfirm.Ctx(ctx).Where(dao.ProductLevelConfirm.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.LevelConfirm = info
	return res, nil
}

func (s *sLevelConfirm) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneLevelConfirmReq{LevelConfirm: &v1.LevelConfirmInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || info == nil {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.ProductLevelConfirm.Ctx(ctx).Where(dao.ProductLevelConfirm.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sLevelConfirm) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.ProductLevelConfirm, error) {
	info := &entity.ProductLevelConfirm{}
	err := dao.ProductLevelConfirm.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sLevelConfirm) checkInputData(ctx context.Context, in *v1.LevelConfirmInfo) (*v1.LevelConfirmInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("项目优先级不能为空")
	}

	if in.GetScoreMin() < 0 {
		return in, errors.New("评分下限不能小于0")
	}

	if in.GetScoreMax() < 0 || in.GetScoreMax() > 100 {
		return in, errors.New("评分上限不能小于0，且不能大于100")
	}

	if in.GetScoreRange() != consts.ScoreRangeMin && in.GetScoreRange() != consts.ScoreRangeMax && in.GetScoreRange() != consts.ScoreRangeMinMax {
		return in, errors.New("评分区间包含关系错误")
	}
	if in.GetIsNeedPm() != consts.IsNeedPm && in.GetIsNeedPm() != consts.IsNotNeedPm {
		return in, errors.New("是否委派项目经理数据错误")
	}
	if in.GetIsNeedPm() == consts.IsNeedPm && len(in.GetPmDemand()) == 0 {
		return in, errors.New("项目经理要求不能为空")
	}
	if len(in.GetProductDemand()) == 0 {
		return in, errors.New("项目管理要求不能为空")
	}
	if in.GetIsNeedPml() != consts.IsNeedPml && in.GetIsNeedPml() != consts.IsNotNeedPml {
		return in, errors.New("是否需要项目负责人数据错误")
	}

	// 评价权重和不能大于1
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductLevelConfirm.Columns().Name): in.GetName(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前优先级已存在，请确认输入信息是否正确")
	}

	return in, nil
}
