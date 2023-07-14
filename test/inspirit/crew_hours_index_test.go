package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_CrewHoursIndex_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewHoursIndexClient(conn)
		res    *v1.GetListCrewHoursIndexRes
		err    error
		size   int32 = 10
	)
	res, err = depert.GetList(ctx, &v1.GetListCrewHoursIndexReq{
		Page: 1,
		Size: size,
		CrewHoursIndex: &v1.CrewHoursIndexInfo{
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

func Test_CrewHoursIndex_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewHoursIndexClient(conn)
		res    *v1.GetOneCrewHoursIndexRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneCrewHoursIndexReq{
		CrewHoursIndex: &v1.CrewHoursIndexInfo{
			ScoreIndex: 1,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewHoursIndex_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewHoursIndexClient(conn)
		res    *v1.CreateCrewHoursIndexRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateCrewHoursIndexReq{
		ScoreIndex: 6,
		ScoreMin:   0.8,
		ScoreMax:   1,
		ScoreRange: 2,
		Remark:     "test-info",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewHoursIndex_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewHoursIndexClient(conn)
		res    *v1.ModifyCrewHoursIndexRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyCrewHoursIndexReq{
		Id:         3,
		ScoreIndex: 3,
		ScoreMin:   0.4,
		ScoreMax:   0.6,
		ScoreRange: 2,
		Remark:     "test-",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewHoursIndex_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewHoursIndexClient(conn)
		res    *v1.DeleteCrewHoursIndexRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteCrewHoursIndexReq{
		Id: 6,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
