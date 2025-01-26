package main

import (
	"flag"
	"fmt"

	"picture/rpc/space-rpc/internal/config"
	"picture/rpc/space-rpc/internal/server"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/space.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		space.RegisterSpaceServiceServer(grpcServer, server.NewSpaceServiceServer(ctx))

		// 注册服务
		server.NewSpaceServiceServer(ctx)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
