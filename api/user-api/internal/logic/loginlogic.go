package logic

import (
	"context"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/rpc/user-rpc/user" // 更新这个导入路径

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 调用 RPC 服务
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		UserAccount:  req.UserAccount,
		UserPassword: req.UserPassword,
	})
	if err != nil {
		logx.Errorf("登陆失败，Login failed: %v", err)
		return nil, err
	}

	return &types.LoginResp{
		Id:          res.User.Id,
		UserAccount: res.User.UserAccount,
		UserName:    res.User.UserName,
		UserAvatar:  res.User.UserAvatar,
		UserProfile: res.User.UserProfile,
		UserRole:    res.User.UserRole,
		AccessToken: res.AccessToken,
	}, nil
}
