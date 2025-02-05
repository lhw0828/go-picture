package svc

import (
	"picture/rpc/user-rpc/internal/config"
	"picture/rpc/user-rpc/internal/dao"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	UserDao *dao.UserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		UserDao: dao.NewUserDao(conn),
	}
}
