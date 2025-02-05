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
func (l *AddUserLogic) AddUser(in *user.UserAddRequest) (*user.RegisterResponse, error) {
	// Parameter validation
	if in == nil || len(in.UserAccount) == 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 检查账号是否已存在
	existUser, err := l.svcCtx.UserDao.FindByUserAccount(l.ctx, in.UserAccount)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, errorx.NewCodeError(errorx.UserExist, "账号已存在")
	}

	// 默认密码加密
	defaultPassword := "12345678"
	encryptPassword, err := utils.EncryptPassword(defaultPassword)
	if err != nil {
		return nil, err
	}

	// 创建用户
	newUser := &model.User{
		UserAccount:  in.UserAccount,
		UserPassword: encryptPassword,
		UserName:     sql.NullString{String: in.UserName, Valid: true},
		UserAvatar:   sql.NullString{String: in.UserAvatar, Valid: true},
		UserProfile:  sql.NullString{String: in.UserProfile, Valid: true},
		UserRole:     in.UserRole,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		IsDelete:     0, // 修改为直接使用 int32 类型
	}

	result, err := l.svcCtx.UserDao.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		Id: userId,
	}, nil
}
