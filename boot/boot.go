package boot

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ctx            context.Context
	JobLevelServer v1.JobLevelClient
	//DepertmentServer  v1.DepartmentClient
	//JobServer         v1.JobClient
	//EmployeeServer    v1.EmployeeClient
	//EmployeeJobServer v1.EmployeeJobClient
)

// 用于应用初始化。
func init() {
	ctx = context.Background()
	// 部门、员工基础信息服务
	organizeServerName, _ := g.Config("manifest/config/config.yaml").Get(ctx, "grpc.organize.link")
	OrganizeServer, err := grpc.Dial(gconv.String(organizeServerName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	//DepertmentServer = v1.NewDepartmentClient(OrganizeServer)
	//EmployeeServer = v1.NewEmployeeClient(OrganizeServer)
	//JobServer = v1.NewJobClient(OrganizeServer)
	JobLevelServer = v1.NewJobLevelClient(OrganizeServer)
	//EmployeeJobServer = v1.NewEmployeeJobClient(OrganizeServer)
}
func init2() {
	ctx = context.Background()
	baseServerName, err := g.Config("manifest/config/config.yaml").Get(ctx, "grpc.organize.name")
	if err != nil {
		g.Log().Error(ctx, "get organize server config name error: %v", err)
	}
	BaseServer := grpcx.Client.MustNewGrpcClientConn(baseServerName.String())
	//DepertmentServer = v1.NewDepartmentClient(BaseServer)
	//EmployeeServer = v1.NewEmployeeClient(BaseServer)
	//JobServer = v1.NewJobClient(BaseServer)
	JobLevelServer = v1.NewJobLevelClient(BaseServer)
	//EmployeeJobServer = v1.NewEmployeeJobClient(BaseServer)
}
