package main

import (
	"flag"
	"fmt"

	"picture/rpc/picture-rpc/internal/config"
	"picture/rpc/picture-rpc/internal/server"
	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/picture.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		picture.RegisterPictureServiceServer(grpcServer, server.NewPictureServiceServer(ctx))
		reflection.Register(grpcServer) // 确保这行代码存在
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
