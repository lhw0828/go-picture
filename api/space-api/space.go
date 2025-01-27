package main

import (
	"flag"
	"fmt"
	"net/http"

	"picture/api/space-api/internal/config"
	"picture/api/space-api/internal/handler"
	"picture/api/space-api/internal/svc"
	"picture/common/errorx"
	"picture/common/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/space.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "认证失败"))
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	// 使用封装的错误处理器
	httpx.SetErrorHandler(errorx.ErrorHandler)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
