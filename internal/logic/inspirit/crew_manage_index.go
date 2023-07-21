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
	productSvr "github.com/lj1570693659/gfcq_config/internal/service/product"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productPb "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

type (
	sCrewManageIndex struct{}
)

func init() {
	service.RegisterCrewManageIndex(&sCrewManageIndex{})
}

func (s *sCrewManageIndex) Create(ctx context.Context, in *v1.CreateCrewManageIndexReq) (*v1.CreateCrewManageIndexRes, error) {
	res := &v1.CreateCrewManageIndexRes{
		CrewManageIndex: &v1.CrewManageIndexInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.CrewManageIndexInfo{
		ScoreIndex:    in.GetScoreIndex(),
		ProductRoleId: in.GetProductRoleId(),
		Remark:        in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.CrewManageIndex{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.CrewManageIndex.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.CrewManageIndex.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sCrewManageIndex) GetOne(ctx context.Context, in *v1.GetOneCrewManageIndexReq) (*v1.GetOneCrewManageIndexRes, error) {
	var info *v1.CrewManageIndexInfo
	query := dao.CrewManageIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewManageIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().ScoreIndex, in.GetCrewManageIndex().GetScoreIndex())
	}
	if in.GetCrewManageIndex().GetProductRoleId() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().ProductRoleId, in.GetCrewManageIndex().GetProductRoleId())
	}
	// 主键查询
	if in.GetCrewManageIndex().GetId() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().Id, in.GetCrewManageIndex().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneCrewManageIndexRes{
		CrewManageIndex: info,
	}, err
}

func (s *sCrewManageIndex) GetList(ctx context.Context, in *v1.GetListCrewManageIndexReq) (*v1.GetListCrewManageIndexRes, error) {
	res := &v1.GetListCrewManageIndexRes{
		Page: in.GetPage(),
		Size: in.GetSize(),
	}
	resData := make([]*v1.CrewManageIndexInfo, 0)
	budgetEntity := make([]entity.CrewManageIndex, 0)

	query := dao.CrewManageIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewManageIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().ScoreIndex, in.GetCrewManageIndex().GetScoreIndex())
	}
	if in.GetCrewManageIndex().GetProductRoleId() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().ProductRoleId, in.GetCrewManageIndex().GetProductRoleId())
	}
	// 主键查询
	if in.GetCrewManageIndex().GetId() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().Id, in.GetCrewManageIndex().GetId())
	}
	// 备注
	if len(in.GetCrewManageIndex().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewManageIndex.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewManageIndex().GetRemark(), "%")})
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

func (s *sCrewManageIndex) GetAll(ctx context.Context, in *v1.GetAllCrewManageIndexReq) (*v1.GetAllCrewManageIndexRes, error) {
	res := &v1.GetAllCrewManageIndexRes{}
	resData := make([]*v1.CrewManageIndexInfo, 0)
	budgetEntity := make([]entity.CrewManageIndex, 0)

	query := dao.CrewManageIndex.Ctx(ctx)

	// 评价标准
	if in.GetCrewManageIndex().GetScoreIndex() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().ScoreIndex, in.GetCrewManageIndex().GetScoreIndex())
	}
	if in.GetCrewManageIndex().GetProductRoleId() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().ProductRoleId, in.GetCrewManageIndex().GetProductRoleId())
	}
	// 主键查询
	if in.GetCrewManageIndex().GetId() > 0 {
		query = query.Where(dao.CrewManageIndex.Columns().Id, in.GetCrewManageIndex().GetId())
	}
	// 备注
	if len(in.GetCrewManageIndex().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.CrewManageIndex.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetCrewManageIndex().GetRemark(), "%")})
	}

	err := query.Scan(&budgetEntity)

	levelEntityByte, _ := json.Marshal(budgetEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Data = resData
	return res, err
}

func (s *sCrewManageIndex) Modify(ctx context.Context, in *v1.ModifyCrewManageIndexReq) (*v1.ModifyCrewManageIndexRes, error) {
	res := &v1.ModifyCrewManageIndexRes{CrewManageIndex: &v1.CrewManageIndexInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.CrewManageIndexInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.CrewManageIndex{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.CrewManageIndex.Ctx(ctx).Where(dao.CrewManageIndex.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.CrewManageIndex = info
	return res, nil
}

func (s *sCrewManageIndex) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneCrewManageIndexReq{CrewManageIndex: &v1.CrewManageIndexInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.CrewManageIndex.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.CrewManageIndex.Ctx(ctx).Where(dao.CrewManageIndex.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sCrewManageIndex) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.CrewManageIndex, error) {
	info := &entity.CrewManageIndex{}
	err := dao.CrewManageIndex.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sCrewManageIndex) checkInputData(ctx context.Context, in *v1.CrewManageIndexInfo) (*v1.CrewManageIndexInfo, error) {
	if in.GetScoreIndex() < 0 {
		return in, errors.New("工时指数不能小于0")
	}
	if g.IsEmpty(in.GetProductRoleId()) {
		return in, errors.New("项目角色不能为空")
	}

	// 检测角色数据是否有效
	roles, err := productSvr.Roles().GetOne(ctx, &productPb.GetOneRolesReq{Roles: &productPb.RolesInfo{
		Id:  gconv.Int32(in.GetProductRoleId()),
		Pid: 0,
	}})
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(roles) || (!g.IsNil(roles) && g.IsEmpty(roles.Roles.Id)) {
		return in, errors.New("当前角色名称不存在，请确认输入信息是否正确")
	}

	// 管理指数值唯一 && 管理指数对应角色唯一
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.CrewManageIndex.Columns().ScoreIndex): in.GetScoreIndex(),
	}
	conditionRoles := g.Map{
		fmt.Sprintf("%s = ?", dao.CrewManageIndex.Columns().ProductRoleId): in.GetProductRoleId(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
		conditionRoles["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前工时指数已存在，请确认输入信息是否正确")
	}

	getRolesInfo, err := s.GetOneByCondition(ctx, conditionRoles)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if g.IsNil(getRolesInfo) || (!g.IsNil(getRolesInfo) && !g.IsEmpty(getRolesInfo.Id)) {
		return in, errors.New("项目角色对应工时指数已存在，请确认输入信息是否正确")
	}

	return in, nil
}
