package logic

import (
	"context"
	"errors"

	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetCurrentUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser(in *user.GetUserByIdRequest) (*user.UserInfo, error) {
	// 添加日志
	l.Logger.Infof("GetCurrentUser - Request: %+v", in)

	// 从数据库获取用户信息
	userInfo, err := l.svcCtx.UserDao.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			l.Logger.Errorf("User not found: %v", err)
			return nil, errors.New("用户不存在")
		}
		l.Logger.Errorf("Database error: %v", err)
		return nil, err
	}

	// 添加日志
	l.Logger.Infof("GetCurrentUser - Found user: %+v", userInfo)

	return &user.UserInfo{
		Id:          userInfo.Id,
		UserAccount: userInfo.UserAccount,
		UserName:    userInfo.UserName.String,
		UserAvatar:  userInfo.UserAvatar.String,
		UserProfile: userInfo.UserProfile.String,
		UserRole:    userInfo.UserRole,
		CreateTime:  userInfo.CreateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
