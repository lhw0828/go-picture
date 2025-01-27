package main

import (
	"flag"
	"fmt"

	"picture/api/user-api/internal/config"
	"picture/api/user-api/internal/handler"
	"picture/api/user-api/internal/svc"
	"picture/common/errorx"
	"picture/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	// 使用封装的错误处理器
	httpx.SetErrorHandler(errorx.ErrorHandler)

	handler.RegisterHandlers(server, ctx)

	// 添加错误处理中间件
	server.Use(middleware.ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
