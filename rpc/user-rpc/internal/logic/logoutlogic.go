package logic

import (
	"context"

	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutRequest) (*user.BaseResponse, error) {
	// 这里可以添加登出相关的逻辑，比如清除缓存等
	return &user.BaseResponse{
		Code: 0,
		Msg:  "退出登录成功",
	}, nil
}
