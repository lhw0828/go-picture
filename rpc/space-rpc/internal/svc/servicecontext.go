package svc

import (
	"picture/rpc/space-rpc/internal/config"
	"picture/rpc/space-rpc/internal/dao"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	SpaceDao         *dao.SpaceDao
	SpaceMemberDao   *dao.SpaceMemberDao
	SpaceAnalysisDao *dao.SpaceAnalysisDao
	UserRpc          user.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:           c,
		SpaceDao:         dao.NewSpaceDao(conn),
		SpaceMemberDao:   dao.NewSpaceMemberDao(conn),
		SpaceAnalysisDao: dao.NewSpaceAnalysisDao(conn),
		UserRpc:          user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}
