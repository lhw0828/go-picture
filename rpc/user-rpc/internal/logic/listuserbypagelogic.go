package logic

import (
	"context"

	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserByPageLogic {
	return &ListUserByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserByPageLogic) ListUserByPage(in *user.UserQueryRequest) (*user.UserQueryResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserQueryResponse{}, nil
}
