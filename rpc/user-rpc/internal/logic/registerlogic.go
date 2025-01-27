package logic

import (
	"context"
	"time"

	"picture/common/errorx"
	"picture/common/utils"
	"picture/rpc/user-rpc/internal/model"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user" // 更新这个导入路径

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 参数校验
	if len(in.UserAccount) < 4 {
		return nil, errorx.NewCodeError(errorx.LengthLess4, "账号长度不能小于4位")
	}
	if len(in.UserPassword) < 8 || len(in.CheckPassword) < 8 {
		return nil, errorx.NewCodeError(errorx.LengthLess8, "密码长度不能小于8位")
	}
	if in.UserPassword != in.CheckPassword {
		return nil, errorx.NewCodeError(errorx.PasswordNotMatch, "两次输入的密码不一致")
	}

	// 判断账号是否已存在
	existUser, err := l.svcCtx.UserDao.FindByUserAccount(in.UserAccount)
	if err != nil {
		l.Logger.Errorf("查询用户错误，Find user error: %v", err)
		return nil, err
	}
	if existUser != nil {
		return nil, errorx.NewCodeError(errorx.UserExist, "账号已存在")
	}

	// 创建用户
	newUser := &model.User{
		UserAccount:  in.UserAccount,
		UserPassword: utils.EncryptPassword(in.UserPassword),
		UserName:     "用户" + in.UserAccount,
		UserRole:     "user",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	result, err := l.svcCtx.UserDao.Insert(newUser)
	if err != nil {
		l.Logger.Errorf("插入到数据库错误，Insert user error: %v", err)
		return nil, errorx.NewCodeError(errorx.RegisterFail, "注册失败")
	}

	userId, err := result.LastInsertId()
	if err != nil {
		l.Logger.Errorf("获取最近插入用户id失败，Get last insert id error: %v", err)
		return nil, errorx.NewCodeError(errorx.RegisterFail, "注册失败")
	}

	l.Logger.Infof("注册用户成功，Register success, userId: %d", userId)
	return &user.RegisterResponse{
		Id: userId,
	}, nil
}
