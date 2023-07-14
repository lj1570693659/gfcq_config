package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_BudgetAccess_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewBudgetAssessClient(conn)
		res    *v1.GetListBudgetAssessRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListBudgetAssessReq{
		Page: 2,
		Size: size,
		BudgetAssess: &v1.BudgetAssessInfo{
			ScoreMin: 50,
		},
	})
	fmt.Println("TotalSize=============", res.TotalSize)
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_BudgetAccess_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewBudgetAssessClient(conn)
		res    *v1.GetOneBudgetAssessRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneBudgetAssessReq{
		BudgetAssess: &v1.BudgetAssessInfo{
			Id: 6,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_BudgetAccess_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewBudgetAssessClient(conn)
		res    *v1.CreateBudgetAssessRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateBudgetAssessReq{
		ScoreMin:    100,
		ScoreMax:    100,
		ScoreRange:  1,
		BudgetMin:   70,
		BudgetMax:   100,
		BudgetRange: 2,
		Remark:      "test-info",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_BudgetAccess_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewBudgetAssessClient(conn)
		res    *v1.ModifyBudgetAssessRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyBudgetAssessReq{
		Id:          7,
		ScoreMin:    33,
		ScoreMax:    44,
		ScoreRange:  1,
		BudgetMin:   70,
		BudgetMax:   100,
		BudgetRange: 2,
		Remark:      "test-info8888888888888",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_BudgetAccess_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewBudgetAssessClient(conn)
		res    *v1.DeleteBudgetAssessRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteBudgetAssessReq{
		Id: 7,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
