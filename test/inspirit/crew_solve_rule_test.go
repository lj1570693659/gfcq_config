package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_CrewSolveRule_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewSolveRuleClient(conn)
		res    *v1.GetListCrewSolveRuleRes
		err    error
		size   int32 = 10
	)
	res, err = depert.GetList(ctx, &v1.GetListCrewSolveRuleReq{
		Page:          1,
		Size:          size,
		CrewSolveRule: &v1.CrewSolveRuleInfo{
			//Redio: 2,
		},
	})
	fmt.Println("TotalSize=============", res.TotalSize)
	fmt.Println("res=============", res.Data)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_CrewSolveRule_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewSolveRuleClient(conn)
		res    *v1.GetOneCrewSolveRuleRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneCrewSolveRuleReq{
		CrewSolveRule: &v1.CrewSolveRuleInfo{
			Demand: v1.DemandEnum_highlight,
		},
	})
	fmt.Println("res=============", res.GetCrewSolveRule())
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewSolveRule_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewSolveRuleClient(conn)
		res    *v1.CreateCrewSolveRuleRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateCrewSolveRuleReq{
		Redio:  0.9,
		Demand: v1.DemandEnum_highlight,
		Remark: "有突出贡献",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewSolveRule_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewSolveRuleClient(conn)
		res    *v1.ModifyCrewSolveRuleRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyCrewSolveRuleReq{
		Id:     4,
		Redio:  0.8,
		Demand: v1.DemandEnum_highlight,
		Remark: "有突出贡献",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewSolveRule_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewSolveRuleClient(conn)
		res    *v1.DeleteCrewSolveRuleRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteCrewSolveRuleReq{
		Id: 4,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
