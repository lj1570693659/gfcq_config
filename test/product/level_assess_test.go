package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"testing"
)

func Test_LevelAssess_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelAssessClient(conn)
		res    *v1.GetListLevelAssessRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListLevelAssessReq{
		Page: 2,
		Size: size,
	})
	fmt.Println("res=============", res.TotalSize)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_LevelAssess_GetListWithoutPage(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelAssessClient(conn)
		res    *v1.GetListWithoutLevelAssessRes
		err    error
	)
	res, err = depert.GetListWithoutPage(ctx, &v1.GetListWithoutLevelAssessReq{})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}
func Test_LevelAssess_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelAssessClient(conn)
		res    *v1.GetOneLevelAssessRes
		err    error
	)
	print("depert=============", depert)
	res, err = depert.GetOne(ctx, &v1.GetOneLevelAssessReq{
		LevelAssess: &v1.LevelAssessInfo{
			EvaluateDimensions: "人事室",
			Id:                 0,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_LevelAssess_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelAssessClient(conn)
		res    *v1.CreateLevelAssessRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateLevelAssessReq{
		EvaluateDimensions: "开发难度",
		EvaluateCriteria:   "开发周期-test-22222222222",
		EvaluateId:         16,
		// 你妹的，还得手输
		ScoreCriteria: "开发周期",
		Weight:        0.05,
		Remark:        "开发难度-开发周期-test-22222222222-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_LevelAssess_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelAssessClient(conn)
		res    *v1.ModifyLevelAssessRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyLevelAssessReq{
		Id:                 21,
		EvaluateDimensions: "开发难度",
		EvaluateCriteria:   "开发周期-test-2222",
		EvaluateId:         16,
		ScoreCriteria:      "开发周期",
		Weight:             0.06,
		Remark:             "行政室-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_LevelAssess_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewLevelAssessClient(conn)
		res    *v1.DeleteLevelAssessRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteLevelAssessReq{
		Id: 21,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
