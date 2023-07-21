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
	"github.com/lj1570693659/gfcq_config/internal/consts"
	"github.com/lj1570693659/gfcq_config/internal/dao"
	"github.com/lj1570693659/gfcq_config/internal/library"
	"github.com/lj1570693659/gfcq_config/internal/model/do"
	"github.com/lj1570693659/gfcq_config/internal/model/entity"
	service "github.com/lj1570693659/gfcq_config/internal/service/inspirit"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

type (
	sStageRadio struct{}
)

func init() {
	service.RegisterStageRadio(&sStageRadio{})
}

func (s *sStageRadio) Create(ctx context.Context, in *v1.CreateStageRadioReq) (*v1.CreateStageRadioRes, error) {
	res := &v1.CreateStageRadioRes{
		StageRadio: &v1.StageRadioInfo{},
	}
	if _, err := s.checkInputData(ctx, &v1.StageRadioInfo{
		Sid:        in.GetSid(),
		QuotaRadio: in.GetQuotaRadio(),
		ScoreMin:   in.GetScoreMin(),
		ScoreMax:   in.GetScoreMax(),
		ScoreRange: in.GetScoreRange(),
		Remark:     in.GetRemark(),
	}); err != nil {
		return res, err
	}

	data := do.ProductStageRadio{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.ProductStageRadio.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return res, err
	}

	res.StageRadio.Id = gconv.Int32(lastInsertId)
	return res, nil
}

func (s *sStageRadio) GetOne(ctx context.Context, in *v1.GetOneStageRadioReq) (*v1.GetOneStageRadioRes, error) {
	var info *v1.StageRadioInfo
	query := dao.ProductStageRadio.Ctx(ctx)

	// 浮动比例
	if in.GetStageRadio().GetQuotaRadio() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().QuotaRadio, in.GetStageRadio().GetQuotaRadio())
	}
	if in.GetStageRadio().GetScoreMin() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().ScoreMin, in.GetStageRadio().GetScoreMin())
	}
	if in.GetStageRadio().GetScoreMax() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().ScoreMax, in.GetStageRadio().GetScoreMax())
	}
	if in.GetStageRadio().GetSid() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().Sid, in.GetStageRadio().GetSid())
	}
	// 主键查询
	if in.GetStageRadio().GetId() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().Id, in.GetStageRadio().GetId())
	}

	err := query.Scan(&info)

	return &v1.GetOneStageRadioRes{
		StageRadio: info,
	}, err
}

func (s *sStageRadio) GetList(ctx context.Context, in *v1.GetListStageRadioReq) (*v1.GetListStageRadioRes, error) {
	res := &v1.GetListStageRadioRes{
		Page: in.GetPage(),
		Size: in.GetSize(),
	}
	resData := make([]*v1.StageRadioInfo, 0)
	budgetEntity := make([]entity.ProductStageRadio, 0)

	query := dao.ProductStageRadio.Ctx(ctx)

	// 评价标准
	// 浮动比例
	if in.GetStageRadio().GetQuotaRadio() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().QuotaRadio, in.GetStageRadio().GetQuotaRadio())
	}
	if in.GetStageRadio().GetScoreMin() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().ScoreMin, in.GetStageRadio().GetScoreMin())
	}
	if in.GetStageRadio().GetScoreMax() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().ScoreMax, in.GetStageRadio().GetScoreMax())
	}
	if in.GetStageRadio().GetSid() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().Sid, in.GetStageRadio().GetSid())
	}
	// 备注
	if len(in.GetStageRadio().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductStageRadio.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetStageRadio().GetRemark(), "%")})
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

func (s *sStageRadio) GetAll(ctx context.Context, in *v1.GetAllStageRadioReq) (*v1.GetAllStageRadioRes, error) {
	res := &v1.GetAllStageRadioRes{}
	resData := make([]*v1.StageRadioInfo, 0)
	budgetEntity := make([]entity.ProductStageRadio, 0)

	query := dao.ProductStageRadio.Ctx(ctx)

	// 评价标准
	// 浮动比例
	if in.GetStageRadio().GetQuotaRadio() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().QuotaRadio, in.GetStageRadio().GetQuotaRadio())
	}
	if in.GetStageRadio().GetScoreMin() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().ScoreMin, in.GetStageRadio().GetScoreMin())
	}
	if in.GetStageRadio().GetScoreMax() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().ScoreMax, in.GetStageRadio().GetScoreMax())
	}
	if in.GetStageRadio().GetSid() > 0 {
		query = query.Where(dao.ProductStageRadio.Columns().Sid, in.GetStageRadio().GetSid())
	}
	// 备注
	if len(in.GetStageRadio().GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.ProductStageRadio.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetStageRadio().GetRemark(), "%")})
	}

	err := query.Scan(&budgetEntity)

	levelEntityByte, _ := json.Marshal(budgetEntity)
	json.Unmarshal(levelEntityByte, &resData)
	res.Data = resData
	return res, err
}

func (s *sStageRadio) Modify(ctx context.Context, in *v1.ModifyStageRadioReq) (*v1.ModifyStageRadioRes, error) {
	res := &v1.ModifyStageRadioRes{StageRadio: &v1.StageRadioInfo{}}
	if g.IsEmpty(in.GetId()) {
		return res, errors.New("编辑信息对象不能为空")
	}

	// 输入数据校验
	info := &v1.StageRadioInfo{}
	inputByte, _ := json.Marshal(in)
	json.Unmarshal(inputByte, &info)
	if _, err := s.checkInputData(ctx, info); err != nil {
		return res, err
	}

	data := do.ProductStageRadio{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	data.UpdateTime = gtime.Now()
	_, err = dao.ProductStageRadio.Ctx(ctx).Where(dao.ProductStageRadio.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return res, err
	}

	res.StageRadio = info
	return res, nil
}

func (s *sStageRadio) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneStageRadioReq{StageRadio: &v1.StageRadioInfo{Id: id}})
	if (err != nil && err == sql.ErrNoRows) || (!g.IsNil(info) && g.IsEmpty(info.StageRadio.Id)) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	_, err = dao.ProductStageRadio.Ctx(ctx).Where(dao.ProductStageRadio.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除等级评估配置数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sStageRadio) GetQuotaRadioByScore(ctx context.Context, in *v1.GetQuotaRadioByScoreReq) (*v1.GetQuotaRadioByScoreRes, error) {
	res := &v1.GetQuotaRadioByScoreRes{}
	if g.IsEmpty(in.Score) {
		return res, errors.New("接收到的得分为空")
	}

	allScore, err := s.GetAll(ctx, &v1.GetAllStageRadioReq{})
	if !g.IsNil(err) {
		return res, err
	}
	if len(allScore.GetData()) == 0 {
		return res, errors.New("请先完善应发激励占比配置信息")
	}
	for _, v := range allScore.GetData() {
		switch v.ScoreRange {
		case consts.ScoreRangeMin:
			// 左闭右开
			if v.ScoreMin <= in.Score && in.Score < v.ScoreMax {
				res.QuotaRadio = v.GetQuotaRadio()
				break
			}
		case consts.ScoreRangeMax:
			// 左开右闭
			if v.ScoreMin < in.Score && in.Score <= v.ScoreMax {
				res.QuotaRadio = v.GetQuotaRadio()
				break
			}
		case consts.ScoreRangeMinMax:
			// 左闭右闭
			if v.ScoreMin <= in.Score && in.Score <= v.ScoreMax {
				res.QuotaRadio = v.GetQuotaRadio()
				break
			}
		}
	}
	return res, nil
}

func (s *sStageRadio) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.ProductStageRadio, error) {
	info := &entity.ProductStageRadio{}
	err := dao.ProductStageRadio.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *sStageRadio) checkInputData(ctx context.Context, in *v1.StageRadioInfo) (*v1.StageRadioInfo, error) {
	if in.GetQuotaRadio() < 0 {
		return in, errors.New("分配比例不能小于0")
	}
	if in.GetQuotaRadio() > 1 {
		return in, errors.New("分配比例不能大于1")
	}
	min := library.Decimal(gconv.Float64(in.GetScoreMin()))
	max := library.Decimal(gconv.Float64(in.GetScoreMax()))
	if min < 0 {
		return in, errors.New("评分下限不能小于0")
	}
	if max > 100 {
		return in, errors.New("评分上限不能大于100")
	}
	if min > max {
		return in, errors.New("评分下限不能超过上限")
	}
	if in.GetScoreRange().Number() != consts.ScoreRangeMin && in.GetScoreRange().Number() != consts.ScoreRangeMax && in.GetScoreRange().Number() != consts.ScoreRangeMinMax {
		return in, errors.New("评分区间包含关系错误")
	}

	return in, nil
}
