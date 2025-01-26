package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/rpc/user-rpc/user"

	// 添加这行导入
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
	// 从上下文中获取 userId
	userId := l.ctx.Value("userId")
	if userId == nil {
		l.Logger.Error("Failed to get userId from context")
		return nil, fmt.Errorf("获取用户ID失败")
	}

	// 处理 json.Number 类型
	var id int64
	switch v := userId.(type) {
	case json.Number:
		id, err = v.Int64()
		if err != nil {
			l.Logger.Errorf("Convert userId error: %v", err)
			return nil, fmt.Errorf("用户ID格式错误")
		}
	default:
		l.Logger.Errorf("Invalid userId type: %T", userId)
		return nil, fmt.Errorf("用户ID类型错误")
	}

	l.Logger.Infof("Getting user info for userId: %d", id)

	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdRequest{
		Id: id,
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
