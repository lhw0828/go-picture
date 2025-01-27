package svc

import (
	"picture/api/space-api/internal/config"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	SpaceRpc space.SpaceServiceClient
	UserRpc  user.UserServiceClient
}

// Deadline implements context.Context.
func (s *ServiceContext) Deadline() (deadline time.Time, ok bool) {
	panic("unimplemented")
}

// Done implements context.Context.
func (s *ServiceContext) Done() <-chan struct{} {
	panic("unimplemented")
}

// Err implements context.Context.
func (s *ServiceContext) Err() error {
	panic("unimplemented")
}

// Value implements context.Context.
func (s *ServiceContext) Value(key any) any {
	panic("unimplemented")
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		SpaceRpc: space.NewSpaceServiceClient(zrpc.MustNewClient(c.SpaceRpc).Conn()),
		UserRpc:  user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}
