package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	// Mysql 配置
	Mysql struct {
		DataSource string
	}
}
