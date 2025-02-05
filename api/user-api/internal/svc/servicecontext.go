package svc

import (
	"picture/api/user-api/internal/config"
	"picture/common/middleware"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserServiceClient
	Admin   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn()),
		Admin:   middleware.Admin,
	}
}
