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
	sRoles struct{}
)

func init() {
	service.RegisterRoles(&sRoles{})
}

func (s *sRoles) Create(ctx context.Context, in *v1.CreateRolesReq) (*v1.CreateRolesRes, error) {
	res := &v1.CreateRolesRes{
		Roles: &v1.RolesInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.RolesInfo{
		Name:    in.GetName(),
		Pid:     in.GetPid(),
		Explain: in.GetExplain(),
		Remark:  in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.ProductRoles{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.ProductRoles.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.Roles.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sRoles) GetOne(ctx context.Context, in *v1.GetOneRolesReq) (*v1.GetOneRolesRes, error) {
	info := &v1.RolesInfo{}
	query := dao.ProductRoles.Ctx(ctx)

	// 角色名称
	if len(in.GetRoles().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductRoles.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetRoles().GetName(), "%")})
	}
	// 角色与职责说明
	if len(in.GetRoles().GetExplain()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductRoles.Columns().Explain), g.Slice{fmt.Sprintf("%s%s", in.GetRoles().GetExplain(), "%")})
	}
	// 主键查询
	if in.GetRoles().GetId() > 0 {
		query = query.Where(dao.ProductRoles.Columns().Id, in.GetRoles().GetId())
	}
	// 上级角色
	if in.GetRoles().GetPid() > 0 {
		query = query.Where(dao.ProductRoles.Columns().Pid, in.GetRoles().GetPid())
	}

	err := query.Scan(&info)

	return &v1.GetOneRolesRes{
		Roles: info,
	}, err
}

func (s *sRoles) GetList(ctx context.Context, in *v1.GetListRolesReq) (*v1.GetListRolesRes, error) {
	res := &v1.GetListRolesRes{
		Page:      in.GetPage(),
		Size:      in.GetSize(),
		TotalSize: 0,
		Data:      make([]*v1.RolesInfo, 0),
	}
	resData := make([]*v1.RolesInfo, 0)
	confirmEntity := make([]entity.ProductRoles, 0)

	query := dao.ProductRoles.Ctx(ctx)

	// 角色名称
	if len(in.GetRoles().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductRoles.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetRoles().GetName(), "%")})
	}
	// 角色与职责说明
	if len(in.GetRoles().GetExplain()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductRoles.Columns().Explain), g.Slice{fmt.Sprintf("%s%s", in.GetRoles().GetExplain(), "%")})
	}
	// 主键查询
	if in.GetRoles().GetId() > 0 {
		query = query.Where(dao.ProductRoles.Columns().Id, in.GetRoles().GetId())
	}
	// 上级角色
	if in.GetRoles().GetPid() > 0 {
		query = query.Where(dao.ProductRoles.Columns().Pid, in.GetRoles().GetPid())
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

func (s *sRoles) GetAll(ctx context.Context, in *v1.GetAllRolesReq) (*v1.GetAllRolesRes, error) {
	res := &v1.GetAllRolesRes{
		Data: make([]*v1.RolesInfo, 0),
	}
	resData := make([]*v1.RolesInfo, 0)
	confirmEntity := make([]entity.ProductRoles, 0)

	query := dao.ProductRoles.Ctx(ctx)

	// 角色名称
	if len(in.GetRoles().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductRoles.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetRoles().GetName(), "%")})
	}
	// 角色与职责说明
	if len(in.GetRoles().GetExplain()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductRoles.Columns().Explain), g.Slice{fmt.Sprintf("%s%s", in.GetRoles().GetExplain(), "%")})
	}
	// 主键查询
	if in.GetRoles().GetId() > 0 {
		query = query.Where(dao.ProductRoles.Columns().Id, in.GetRoles().GetId())
	}
	// 上级角色
	if in.GetRoles().GetPid() > 0 {
		query = query.Where(dao.ProductRoles.Columns().Pid, in.GetRoles().GetPid())
	} else if in.GetRoles().GetPid() == -1 {
		query = query.Where(dao.ProductRoles.Columns().Pid, 0)
	}

	err := query.Scan(&confirmEntity)

	levelEntityByte, _ := json.Marshal(confirmEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Data = resData
	return res, err
}

func (s *sRoles) Modify(ctx context.Context, in *v1.ModifyRolesReq) (*v1.ModifyRolesRes, error) {
	res := &v1.ModifyRolesRes{Roles: &v1.RolesInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.RolesInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductRoles{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductRoles.Ctx(ctx).Where(dao.ProductRoles.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.Roles = info
	return res, nil
}

func (s *sRoles) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneRolesReq{Roles: &v1.RolesInfo{Id: id}})
	if (err != nil && err.Error() == sql.ErrNoRows.Error()) || (!g.IsNil(info) && g.IsEmpty(info.Roles.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	fmt.Println("info-----44444444444444-----------", info)
	// 删除父级数据时检查子级数据是否删除完全
	if g.IsEmpty(info.Roles.GetPid()) {
		info, err := s.GetOne(ctx, &v1.GetOneRolesReq{Roles: &v1.RolesInfo{Pid: id}})
		fmt.Println("info-----33333-----------", info)
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, err.Error(), err
		}
		if !g.IsNil(info) && !g.IsEmpty(info.Roles.Id) {
			fmt.Println("info-----111111111-----------", info)
			return false, "请先删除下级元素", errors.New("存在下级元素未删除完全")
		}
	}
	fmt.Println("info-----2222222222-----------", info)
	_, err = dao.ProductRoles.Ctx(ctx).Where(dao.ProductRoles.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sRoles) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.ProductRoles, error) {
	info := &entity.ProductRoles{}
	err := dao.ProductRoles.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sRoles) checkInputData(ctx context.Context, in *v1.RolesInfo) (*v1.RolesInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("角色名称不能为空")
	}

	// 角色名称唯一
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductRoles.Columns().Name): in.GetName(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前角色名称已存在，请确认输入信息是否正确")
	}

	return in, nil
}
