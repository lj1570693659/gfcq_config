package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_CrewDutyIndex_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewDutyIndexClient(conn)
		res    *v1.GetListCrewDutyIndexRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListCrewDutyIndexReq{
		Page: 1,
		Size: size,
		CrewDutyIndex: &v1.CrewDutyIndexInfo{
			ScoreIndex: 2,
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

func Test_CrewDutyIndex_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewDutyIndexClient(conn)
		res    *v1.GetOneCrewDutyIndexRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneCrewDutyIndexReq{
		CrewDutyIndex: &v1.CrewDutyIndexInfo{
			ScoreIndex: 1,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewDutyIndex_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewDutyIndexClient(conn)
		res    *v1.CreateCrewDutyIndexRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateCrewDutyIndexReq{
		ScoreIndex: 2,
		JobLevelId: 30,
		Arith:      v1.ArithEnum_eq,
		Remark:     "test-info",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewDutyIndex_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewDutyIndexClient(conn)
		res    *v1.ModifyCrewDutyIndexRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyCrewDutyIndexReq{
		Id:         4,
		ScoreIndex: 2,
		JobLevelId: 30,
		Arith:      v1.ArithEnum_eq,
		Remark:     "test-",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewDutyIndex_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewDutyIndexClient(conn)
		res    *v1.DeleteCrewDutyIndexRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteCrewDutyIndexReq{
		Id: 4,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
