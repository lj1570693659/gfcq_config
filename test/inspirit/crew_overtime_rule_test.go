package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"testing"
)

func Test_CrewOvertimeRule_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewOvertimeRuleClient(conn)
		res    *v1.GetListCrewOvertimeRuleRes
		err    error
		size   int32 = 10
	)
	res, err = depert.GetList(ctx, &v1.GetListCrewOvertimeRuleReq{
		Page:             1,
		Size:             size,
		CrewOvertimeRule: &v1.CrewOvertimeRuleInfo{
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

func Test_CrewOvertimeRule_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewOvertimeRuleClient(conn)
		res    *v1.GetOneCrewOvertimeRuleRes
		err    error
	)

	res, err = depert.GetOne(ctx, &v1.GetOneCrewOvertimeRuleReq{
		CrewOvertimeRule: &v1.CrewOvertimeRuleInfo{
			Id: 3,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewOvertimeRule_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewOvertimeRuleClient(conn)
		res    *v1.CreateCrewOvertimeRuleRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateCrewOvertimeRuleReq{
		Redio:      1.1,
		ScoreMin:   0.3,
		ScoreMax:   1,
		ScoreRange: 2,
		Remark:     "超出标准工时比例大于0.3",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewOvertimeRule_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewOvertimeRuleClient(conn)
		res    *v1.ModifyCrewOvertimeRuleRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyCrewOvertimeRuleReq{
		Id:         6,
		Redio:      0.99,
		ScoreMin:   0.3,
		ScoreMax:   1,
		ScoreRange: 2,
		Remark:     "超出标准工时比例大于0.3",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_CrewOvertimeRule_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewCrewOvertimeRuleClient(conn)
		res    *v1.DeleteCrewOvertimeRuleRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteCrewOvertimeRuleReq{
		Id: 6,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
