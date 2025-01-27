package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"

	"picture/api/user-api/internal/config"
	"picture/api/user-api/internal/handler"
	"picture/api/user-api/internal/svc"
	"picture/common/constants"
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

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		var e *constants.CodeError
		switch {
		case errors.As(err, &e):
			return http.StatusOK, constants.Fail(e)
		default:
			return http.StatusInternalServerError, nil
		}
	})

	handler.RegisterHandlers(server, ctx)

	// 添加错误处理中间件
	server.Use(middleware.ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
