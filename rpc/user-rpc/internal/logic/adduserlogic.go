package logic

import (
	"context"
	"database/sql"
	"time"

	"picture/common/errorx"
	"picture/common/utils"
	"picture/rpc/user-rpc/internal/model"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 管理功能
func (l *AddUserLogic) AddUser(in *user.AddUserRequest) (*user.AddUserResponse, error) {
	// Parameter validation
	if in == nil || len(in.UserAccount) == 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 默认密码加密
	defaultPassword := "12345678"
	encryptPassword, err := utils.EncryptPassword(defaultPassword)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{ // 修改变量名避免冲突
		UserAccount:  in.UserAccount,
		UserPassword: encryptPassword,
		UserName:     in.UserName,
		UserAvatar:   sql.NullString{String: in.UserAvatar, Valid: true},
		UserProfile:  sql.NullString{String: in.UserProfile, Valid: true},
		UserRole:     in.UserRole,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	result, err := l.svcCtx.UserDao.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &user.AddUserResponse{
		Id: userId,
	}, nil
}
