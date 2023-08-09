package main

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_config/internal/controller/inspirit"
	"github.com/lj1570693659/gfcq_config/internal/controller/product"
	_ "github.com/lj1570693659/gfcq_config/internal/logic/inspirit"
	_ "github.com/lj1570693659/gfcq_config/internal/logic/product"
	"google.golang.org/grpc"
	"net"
)

func main() {
	etcdLink, _ := g.Config("config.yaml").Get(context.Background(), "grpc.etcdLink")
	fmt.Println(etcdLink.String())
	//3. 设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", etcdLink.String())
	if err != nil {
		fmt.Println(err)
	}

	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	inspirit.BudgetAccessRegister(grpcServer)
	inspirit.CrewDutyIndexRegister(grpcServer)
	inspirit.CrewHoursIndexRegister(grpcServer)
	inspirit.CrewKpiRuleRegister(grpcServer)
	inspirit.CrewManageIndexRegister(grpcServer)
	inspirit.CrewOvertimeRuleRegister(grpcServer)
	inspirit.CrewSolveRuleRegister(grpcServer)
	inspirit.ProductStageRadioRegister(grpcServer)

	product.LevelAssessRegister(grpcServer)
	product.LevelConfirmRegister(grpcServer)
	product.ModeRegister(grpcServer)
	product.ModeStageRegister(grpcServer)
	product.RolesRegister(grpcServer)
	product.TypeRegister(grpcServer)

	// 4退出关闭监听
	defer listener.Close()
	//5、启动服务
	grpcServer.Serve(listener)
}
