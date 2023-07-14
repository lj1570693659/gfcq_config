package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/lj1570693659/gfcq_config/internal/cmd"
	_ "github.com/lj1570693659/gfcq_config/internal/logic/inspirit"
	_ "github.com/lj1570693659/gfcq_config/internal/logic/product"
)

func main() {
	//etcdLink, _ := g.Config("config.yaml").Get(context.Background(), "grpc.etcdLink")
	//grpcx.Resolver.Register(etcd.New(etcdLink.String()))
	cmd.Main.Run(gctx.New())
}

func main2() {
	//etcdLink, _ := g.Config("config.yaml").Get(context.Background(), "grpc.etcdLink")
	//etcdName, _ := g.Config("config.yaml").Get(context.Background(), "grpc.name")
	//
	//// 注册ETCD
	//service, err := cmd.NewLocalDefNamingService(etcdLink.String(), etcdName.String())
	//log.Fatalf("failed to create NamingService: %v", service)
	//if err != nil {
	//	log.Fatalf("failed to create NamingService: %v", err)
	//}
	//resolver, err := service.NewEtcdResolver()
	//if err != nil {
	//	log.Fatalf("Create etcd resolver error: %v", err)
	//}
	//// 连接服务端
	//conn, err := grpc.Dial("etcd://localhost:23792/"+service.GetPathServerName("s1"), grpc.WithInsecure(), grpc.WithResolvers(resolver),
	//	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	//if err != nil {
	//	log.Fatalf("Conn server error: %v", err)
	//}
	//log.Printf("Conn success: %v", conn.GetState())
	//// 执行完方法自动关闭资源
	//defer func() {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatalf("Close conn error: %v", err)
	//		return
	//	}
	//	log.Println("Close conn success")
	//}()
	//
	//cmd.Main.Run(gctx.New())
}
