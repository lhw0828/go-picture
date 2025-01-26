package svc

import (
	"picture/api/space-api/internal/config"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	SpaceRpc space.SpaceServiceClient
	UserRpc  user.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		SpaceRpc: space.NewSpaceServiceClient(zrpc.MustNewClient(c.SpaceRpc).Conn()),
		UserRpc:  user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}
