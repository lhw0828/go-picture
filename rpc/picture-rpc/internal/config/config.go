package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL struct {
		DataSource string
	}
	DB struct {
		DataSource string
	}
	COS struct {
		SecretID  string
		SecretKey string
		BucketURL string
		Bucket    string
		Region    string
	}
	Upload struct {
		MaxSize    int64
		AllowTypes []string
	}
	FileStorage struct {
		Type string
	}
	SpaceRpc zrpc.RpcClientConf
}
