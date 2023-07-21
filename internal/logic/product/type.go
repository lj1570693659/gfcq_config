package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
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

func (s *sType) checkInputData(ctx context.Context, in *v1.ModeInfo) (*v1.ModeInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("研发模式不能为空")
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
