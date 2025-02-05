package logic

import (
	"context"
	"errors"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdReq) (*types.UserVO, error) {
	if req == nil || req.Id == 0 {
		return nil, errors.New("无效的用户ID")
	}

	// 使用 GetCurrentUser 方法
	res, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("GetUserById failed: %v", err)
		return nil, errors.New("获取用户信息失败")
	}

	if res == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserVO{
		Id:          res.Id,
		UserAccount: res.UserAccount,
		UserName:    res.UserName,
		UserAvatar:  res.UserAvatar,
		UserProfile: res.UserProfile,
		UserRole:    res.UserRole,
		CreateTime:  res.CreateTime,
	}, nil
}
