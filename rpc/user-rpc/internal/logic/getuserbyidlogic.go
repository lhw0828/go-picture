package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	// Add logging
	l.Logger.Infof("rpc调用，GetUserById request: %v", in)

	// Get user from database
	userModel, err := l.svcCtx.UserDao.FindById(in.Id)
	if err != nil {
		l.Logger.Errorf("Find user error: %v", err)
		return nil, err
	}
	if userModel == nil {
		return nil, errorx.NewCodeError(errorx.UserNotExist, "用户不存在")
	}

	return &user.GetUserByIdResponse{
		Id:          userModel.Id,
		UserAccount: userModel.UserAccount,
		UserName:    userModel.UserName,
		UserAvatar:  userModel.UserAvatar.String,
		UserProfile: userModel.UserProfile.String,
		UserRole:    userModel.UserRole,
	}, nil
}
