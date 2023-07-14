package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"testing"
)

func Test_ModeStage_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeStageClient(conn)
		res    *v1.GetListModeStageRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListModeStageReq{
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

func Test_ModeStage_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeStageClient(conn)
		res    *v1.GetOneModeStageRes
		err    error
	)
	print("depert=============", depert)
	res, err = depert.GetOne(ctx, &v1.GetOneModeStageReq{
		ModeStage: &v1.ModeStageInfo{
			Name: "很低",
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_ModeStage_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeStageClient(conn)
		res    *v1.CreateModeStageRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateModeStageReq{
		Tid:        1,
		Name:       "G00",
		QuotaRadio: 0.05,
		Remark:     "项目验收-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_ModeStage_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeStageClient(conn)
		res    *v1.ModifyModeStageRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyModeStageReq{
		Id:         4,
		Tid:        1,
		Name:       "G2",
		QuotaRadio: 0.1,
		Remark:     "工程认可-备注信息4444",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_ModeStage_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewModeStageClient(conn)
		res    *v1.DeleteModeStageRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteModeStageReq{
		Id: 7,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
