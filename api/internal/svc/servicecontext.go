package svc

import (
	"picture/api/internal/config"
	"picture/rpc/user/pb/user" // 更新这个导入路径

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn())}
}
