package logic

import (
	"context"
	"database/sql"
	"time"

	"picture/common/errorx"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UserUpdateRequest) (*user.BaseResponse, error) {
	if in == nil || in.Id == 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	existUser, err := l.svcCtx.UserDao.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if existUser == nil {
		return nil, errorx.NewCodeError(errorx.UserNotExist, "用户不存在")
	}

	existUser.UserName = sql.NullString{String: in.UserName, Valid: true}
	existUser.UserAvatar = sql.NullString{String: in.UserAvatar, Valid: true}
	existUser.UserProfile = sql.NullString{String: in.UserProfile, Valid: true}
	// 如果传入的用户角色为空，不修改，否则修改
	if in.UserRole != "" {
		existUser.UserRole = in.UserRole
	}
	existUser.UpdateTime = time.Now()

	err = l.svcCtx.UserDao.Update(l.ctx, existUser)
	if err != nil {
		return nil, err
	}

	return &user.BaseResponse{
		Code: 0,
		Msg:  "更新成功",
	}, nil
}
