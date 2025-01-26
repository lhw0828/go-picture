package logic

import (
	"context"

	"picture/api/internal/svc"
	"picture/api/internal/types"
	"picture/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		UserAccount:   req.UserAccount,
		UserPassword:  req.UserPassword,
		CheckPassword: req.CheckPassword,
	})
	if err != nil {
		l.Logger.Errorf("Register failed: %v", err)
		return nil, err
	}

	return &types.RegisterResp{
		Id: res.GetId(),
	}, nil
}
