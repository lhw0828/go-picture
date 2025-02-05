package logic

import (
	"context"

	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceLogic {
	return &ListSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间列表
func (l *ListSpaceLogic) ListSpace(in *space.ListSpaceRequest) (*space.ListSpaceResponse, error) {
	// todo: add your logic here and delete this line

	return &space.ListSpaceResponse{}, nil
}
