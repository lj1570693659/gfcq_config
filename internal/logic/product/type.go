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
	sType struct{}
)

func init() {
	service.RegisterType(&sType{})
}

func (s *sType) GetOne(ctx context.Context, in *v1.GetOneTypeReq) (*v1.GetOneTypeRes, error) {
	var info *v1.TypeInfo
	query := dao.ProductType.Ctx(ctx)

	// 优先级
	if len(in.GetType().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductType.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetType().GetName(), "%")})
	}
	// 主键查询
	if in.GetType().GetId() > 0 {
		query = query.Where(dao.ProductType.Columns().Id, in.GetType().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneTypeRes{
		Type: info,
	}, err
}

func (s *sType) GetList(ctx context.Context, in *v1.GetListTypeReq) (*v1.GetListTypeRes, error) {
	res := &v1.GetListTypeRes{
		Page:      in.GetPage(),
		Size:      in.GetSize(),
		TotalSize: 0,
		Data:      make([]*v1.TypeInfo, 0),
	}
	resData := make([]*v1.TypeInfo, 0)
	confirmEntity := make([]entity.ProductType, 0)

	query := dao.ProductType.Ctx(ctx)

	// 优先级
	if len(in.GetType().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductType.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetType().GetName(), "%")})
	}
	// 主键查询
	if in.GetType().GetId() > 0 {
		query = query.Where(dao.ProductType.Columns().Id, in.GetType().GetId())
	}

	// 主键
	if len(in.GetType().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductType.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetType().GetRemark(), "%")})
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

func (s *sType) GetAll(ctx context.Context, in *v1.GetAllTypeReq) (*v1.GetAllTypeRes, error) {
	resData := make([]*v1.TypeInfo, 0)
	confirmEntity := make([]entity.ProductType, 0)

	query := dao.ProductType.Ctx(ctx)

	// 优先级
	if len(in.GetType().GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductType.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetType().GetName(), "%")})
	}
	// 主键查询
	if in.GetType().GetId() > 0 {
		query = query.Where(dao.ProductType.Columns().Id, in.GetType().GetId())
	}

	// 主键
	if len(in.GetType().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductType.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetType().GetRemark(), "%")})
	}

	err := query.Scan(&confirmEntity)

	levelEntityByte, _ := json.Marshal(confirmEntity)
	json.Unmarshal(levelEntityByte, &resData)

	return &v1.GetAllTypeRes{
		Data: resData,
	}, err
}

func (s *sType) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.ProductType, error) {
	info := &entity.ProductType{}
	err := dao.ProductType.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sType) checkInputData(ctx context.Context, in *v1.TypeInfo) (*v1.TypeInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("项目类型不能为空")
	}

	// 评价权重和不能大于1
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductType.Columns().Name): in.GetName(),
	}
	if in.GetId() > 0 {
		condition["id != ?"] = in.GetId()
	}
	getInfo, err := s.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	if g.IsNil(getInfo) || (!g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id)) {
		return in, errors.New("当前项目类型已存在，请确认输入信息是否正确")
	}

	return in, nil
}

func (s *sType) Create(ctx context.Context, in *v1.CreateTypeReq) (*v1.CreateTypeRes, error) {
	res := &v1.CreateTypeRes{
		Type: &v1.TypeInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.TypeInfo{
		Name:   in.GetName(),
		Remark: in.Remark,
	}); err != nil {
		return res, err
	}

	data := do.ProductType{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.ProductType.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.Type.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sType) Modify(ctx context.Context, in *v1.ModifyTypeReq) (*v1.ModifyTypeRes, error) {
	res := &v1.ModifyTypeRes{Type: &v1.TypeInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.TypeInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductType{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductType.Ctx(ctx).Where(dao.ProductMode.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.Type = info
	return res, nil
}

func (s *sType) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneTypeReq{Type: &v1.TypeInfo{Id: id}})
	if (err != nil && err.Error() == sql.ErrNoRows.Error()) || info == nil {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.ProductType.Ctx(ctx).Where(dao.ProductMode.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}
