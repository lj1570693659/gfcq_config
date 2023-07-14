package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"testing"
)

func Test_LevelConfirm_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelConfirmClient(conn)
		res    *v1.GetListLevelConfirmRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListLevelConfirmReq{
		Page: 1,
		Size: size,
	})
	fmt.Println("res=============", res)
	fmt.Println("TotalSize=============", res.TotalSize)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_LevelConfirm_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelConfirmClient(conn)
		res    *v1.GetOneLevelConfirmRes
		err    error
	)
	print("depert=============", depert)
	res, err = depert.GetOne(ctx, &v1.GetOneLevelConfirmReq{
		LevelConfirm: &v1.LevelConfirmInfo{
			Name: "很低",
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_LevelConfirm_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelConfirmClient(conn)
		res    *v1.CreateLevelConfirmRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateLevelConfirmReq{
		Name:          "很高2",
		ScoreMin:      80,
		ScoreMax:      100,
		ScoreRange:    v1.ScoreRangeEnum_includeMin,
		IsNeedPm:      v1.IsNeedPmEnum_NeedPm,
		PmDemand:      "12+级",
		ProductDemand: "周会",
		MonitDemand:   "周会监控进度",
		IsNeedPml:     v1.IsNeedPmlEnum_NeedPml,
		Remark:        "很高-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_LevelConfirm_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelConfirmClient(conn)
		res    *v1.ModifyLevelConfirmRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyLevelConfirmReq{
		Id:            6,
		Name:          "很高2",
		ScoreMin:      80,
		ScoreMax:      100,
		ScoreRange:    v1.ScoreRangeEnum_includeMin,
		IsNeedPm:      v1.IsNeedPmEnum_NeedPm,
		PmDemand:      "12+级",
		ProductDemand: "周会",
		MonitDemand:   "周会监控进度",
		IsNeedPml:     v1.IsNeedPmlEnum_NeedPml,
		Remark:        "很高-备注信息111111111",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_LevelConfirm_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelConfirmClient(conn)
		res    *v1.DeleteLevelConfirmRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteLevelConfirmReq{
		Id: 7,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
