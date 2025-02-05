package logic

import (
	"context"
	"errors"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注销
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.BaseResp, err error) {
	// 从上下文中获取用户ID
	userId, ok := l.ctx.Value("userId").(int64)
	if !ok || userId == 0 {
		return nil, errors.New("用户未登录")
	}

	res, err := l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutRequest{
		UserId: userId,
	})
	if err != nil {
		l.Logger.Errorf("Logout failed: %v", err)
		return nil, errors.New("注销失败")
	}

	return &types.BaseResp{
		Code:    int(res.Code),
		Message: res.Msg,
	}, nil
}
