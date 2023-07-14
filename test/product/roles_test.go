package product

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"testing"
)

func Test_Roles_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewRolesClient(conn)
		res    *v1.GetListRolesRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListRolesReq{
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

func Test_Roles_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewRolesClient(conn)
		res    *v1.GetOneRolesRes
		err    error
	)
	print("depert=============", depert)
	res, err = depert.GetOne(ctx, &v1.GetOneRolesReq{
		Roles: &v1.RolesInfo{
			Name: "很低",
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_Roles_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewRolesClient(conn)
		res    *v1.CreateRolesRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateRolesReq{
		Name:    "工艺工程师2",
		Pid:     7,
		Explain: "团队成员-工艺工程师2",
		Remark:  "很高-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_Roles_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewRolesClient(conn)
		res    *v1.ModifyRolesRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyRolesReq{
		Id:      9,
		Name:    "工艺工程师2333",
		Pid:     7,
		Explain: "团队成员-工艺工程师2",
		Remark:  "很高-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}

func Test_Roles_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("gfcq_config")
		depert = v1.NewRolesClient(conn)
		res    *v1.DeleteRolesRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteRolesReq{
		Id: 10,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get inspirit list failed: %+v`, err)
	}

}
