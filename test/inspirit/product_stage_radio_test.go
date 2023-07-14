package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_StageRadio_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewStageRadioClient(conn)
		res    *v1.GetListStageRadioRes
		err    error
		size   int32 = 10
	)
	res, err = depert.GetList(ctx, &v1.GetListStageRadioReq{
		Page:       1,
		Size:       size,
		StageRadio: &v1.StageRadioInfo{
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

func Test_StageRadio_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewStageRadioClient(conn)
		res    *v1.GetOneStageRadioRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneStageRadioReq{
		StageRadio: &v1.StageRadioInfo{
			Id: 3,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_StageRadio_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewStageRadioClient(conn)
		res    *v1.CreateStageRadioRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateStageRadioReq{
		QuotaRadio: 1.0,
		ScoreMin:   90,
		ScoreMax:   100,
		ScoreRange: 1,
		Remark:     "超出标准工时比例大于0.3",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_StageRadio_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewStageRadioClient(conn)
		res    *v1.ModifyStageRadioRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyStageRadioReq{
		Id:         5,
		QuotaRadio: 1.0,
		ScoreMin:   90,
		ScoreMax:   100,
		ScoreRange: 1,
		Remark:     "满分绩效",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_StageRadio_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewStageRadioClient(conn)
		res    *v1.DeleteStageRadioRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteStageRadioReq{
		Id: 6,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
