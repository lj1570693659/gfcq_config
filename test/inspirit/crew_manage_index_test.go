package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_CrewManageIndex_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewManageIndexClient(conn)
		res    *v1.GetListCrewManageIndexRes
		err    error
		size   int32 = 10
	)
	res, err = depert.GetList(ctx, &v1.GetListCrewManageIndexReq{
		Page:            1,
		Size:            size,
		CrewManageIndex: &v1.CrewManageIndexInfo{
			//ScoreIndex: 2,
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

func Test_CrewManageIndex_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewManageIndexClient(conn)
		res    *v1.GetOneCrewManageIndexRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneCrewManageIndexReq{
		CrewManageIndex: &v1.CrewManageIndexInfo{
			ScoreIndex: 2,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewManageIndex_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewManageIndexClient(conn)
		res    *v1.CreateCrewManageIndexRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateCrewManageIndexReq{
		ScoreIndex:    3,
		ProductRoleId: 4,
		Remark:        "test-info",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewManageIndex_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewManageIndexClient(conn)
		res    *v1.ModifyCrewManageIndexRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyCrewManageIndexReq{
		Id:            3,
		ScoreIndex:    2,
		ProductRoleId: 7,
		Remark:        "test-",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewManageIndex_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewManageIndexClient(conn)
		res    *v1.DeleteCrewManageIndexRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteCrewManageIndexReq{
		Id: 4,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
