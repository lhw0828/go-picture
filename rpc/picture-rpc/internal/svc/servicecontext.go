package svc

import (
	"database/sql"
	"net/http"
	"net/url"
	"picture/pkg/filemanager"
	"picture/rpc/picture-rpc/internal/config"
	"picture/rpc/picture-rpc/internal/dao"
	"picture/rpc/space-rpc/space" // 修改为正确的导入路径

	_ "github.com/go-sql-driver/mysql"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	DB          *sql.DB
	FileManager filemanager.FileManager
	PictureDao  *dao.PictureDao
	SpaceRpc    space.SpaceServiceClient
}

// 修改 RPC 客户端初始化方式
func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	db, err := sql.Open("mysql", c.MySQL.DataSource)
	if err != nil {
		panic(err)
	}

	// 初始化 COS 客户端
	u, _ := url.Parse(c.COS.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	cosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.COS.SecretID,
			SecretKey: c.COS.SecretKey,
		},
	})

	// 初始化 SpaceRpc 客户端
	spaceRpc := space.NewSpaceServiceClient(zrpc.MustNewClient(c.SpaceRpc).Conn())

	return &ServiceContext{
		Config:      c,
		DB:          db,
		FileManager: filemanager.NewCosManager(cosClient, c.COS.Bucket),
		PictureDao:  dao.NewPictureDao(db),
		SpaceRpc:    spaceRpc,
	}
}
