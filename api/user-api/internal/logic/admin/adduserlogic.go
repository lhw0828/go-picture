package admin

import (
	"context"
	"strings"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/common/errorx"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户
func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UserAddReq) (*types.RegisterResp, error) {
	// 参数校验
	if len(req.UserAccount) == 0 {
		return nil, errorx.NewCodeError(errorx.UserAccountNotNull, errorx.UserAccountNotNullMsg)
	}
	if len(req.UserRole) == 0 {
		return nil, errorx.NewCodeError(errorx.UserRoleNotNull, errorx.UserRoleNotNullMsg)
	}

	// 调用 RPC 服务
	res, err := l.svcCtx.UserRpc.AddUser(l.ctx, &user.UserAddRequest{
		UserAccount: req.UserAccount,
		UserName:    req.UserName,
		UserAvatar:  req.UserAvatar,
		UserProfile: req.UserProfile,
		UserRole:    req.UserRole,
	})
	if err != nil {
		l.Logger.Errorf("AddUser failed: %v", err)
		// 判断具体错误类型
		if strings.Contains(err.Error(), "账号已存在") {
			return nil, errorx.NewCodeError(errorx.UserExist, errorx.UserExistMsg)
		}
		return nil, errorx.NewCodeError(errorx.CreateUserFailed, errorx.CreateUserFailedMsg)
	}

	return &types.RegisterResp{
		Id: res.Id,
	}, nil
}
