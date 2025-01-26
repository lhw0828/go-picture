package logic

import (
	"context"
	"fmt"

	"picture/api/internal/svc"
	"picture/api/internal/types"
	"picture/rpc/user/pb/user"

	"github.com/golang-jwt/jwt/v4" // 添加这行导入
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户信息
func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser() (resp *types.LoginResp, err error) {
	// 从上下文中获取 JWT 解析后的 claims
	claims, ok := l.ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		l.Logger.Error("Failed to get claims from context")
		return nil, fmt.Errorf("获取用户信息失败")
	}

	// 从 claims 中获取 userId
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		l.Logger.Error("Failed to get userId from claims")
		return nil, fmt.Errorf("获取用户ID失败")
	}

	userId := int64(userIdFloat)
	l.Logger.Infof("Getting user info for userId: %d", userId)

	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdRequest{
		Id: userId,
	})
	if err != nil {
		l.Logger.Errorf("GetUserById failed: %v", err)
		return nil, err
	}

	return &types.LoginResp{
		Id:          res.Id,
		UserAccount: res.UserAccount,
		UserName:    res.UserName,
		UserAvatar:  res.UserAvatar,
		UserProfile: res.UserProfile,
		UserRole:    res.UserRole,
	}, nil
}
