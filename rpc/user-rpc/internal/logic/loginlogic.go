package logic

import (
	"context"
	"time"

	"picture/common/errorx"
	"picture/common/utils"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user" // 更新这个导入路径

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 添加日志
	l.Logger.Infof("Login request: %v", in)

	// 参数校验
	if len(in.UserAccount) < 4 {
		return nil, errorx.NewCodeError(errorx.LengthLess4, "账号长度不能小于4位")
	}
	if len(in.UserPassword) < 8 {
		return nil, errorx.NewCodeError(errorx.LengthLess8, "密码长度不能小于8位")
	}

	// 查询用户
	userModel, err := l.svcCtx.UserDao.FindByUserAccount(l.ctx, in.UserAccount)
	if err != nil {
		l.Logger.Errorf("Find user error: %v", err)
		return nil, err
	}
	if userModel == nil {
		return nil, errorx.NewCodeError(errorx.UserNotExist, "用户不存在")
	}

	// 校验密码
    encryptedPassword, err := utils.EncryptPassword(in.UserPassword)
    if err != nil {
        return nil, err
    }
    if userModel.UserPassword != encryptedPassword {
        return nil, errorx.NewCodeError(errorx.PasswordWrong, "密码错误")
    }

	// 生成token
	token, err := l.generateToken(userModel.Id)
	if err != nil {
		l.Logger.Errorf("Generate token error: %v", err)
		return nil, errorx.NewCodeError(errorx.GenerateTokenFail, "生成token失败")
	}

    // 添加返回日志
    // 修复返回数据的类型转换
	resp := &user.LoginResponse{
		User: &user.UserInfo{
			Id:          userModel.Id,
			UserAccount: userModel.UserAccount,
			UserName:    userModel.UserName.String,
			UserAvatar:  userModel.UserAvatar.String,
			UserProfile: userModel.UserProfile.String,
			UserRole:    userModel.UserRole,
		},
		AccessToken: token,
	}
    l.Logger.Infof("Login response: %v", resp)
    return resp, nil
}

func (l *LoginLogic) generateToken(userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("ning4256"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
