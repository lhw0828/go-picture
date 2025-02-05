package main

import (
	"flag"
	"fmt"

	"picture/api/user-api/internal/config"
	"picture/api/user-api/internal/handler"
	"picture/api/user-api/internal/svc"
	"picture/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	// 注册全局中间件
	server.Use(middleware.Recovery())
	server.Use(middleware.Cors())
	server.Use(middleware.RequestLog)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
