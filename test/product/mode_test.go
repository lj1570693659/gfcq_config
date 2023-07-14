package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"testing"
)

func Test_Mode_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeClient(conn)
		res    *v1.GetListModeRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListModeReq{
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

func Test_Mode_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeClient(conn)
		res    *v1.GetOneModeRes
		err    error
	)
	print("depert=============", depert)
	res, err = depert.GetOne(ctx, &v1.GetOneModeReq{
		Mode: &v1.ModeInfo{
			Name: "很低",
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_Mode_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeClient(conn)
		res    *v1.CreateModeRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateModeReq{
		Name:   "代加工2",
		Factor: 0.7,
		Remark: "很高-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_Mode_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeClient(conn)
		res    *v1.ModifyModeRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyModeReq{
		Id:     5,
		Name:   "代加工2444",
		Factor: 0.2,
		Remark: "很高-备注信息111111111",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_Mode_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeClient(conn)
		res    *v1.DeleteModeRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteModeReq{
		Id: 5,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
