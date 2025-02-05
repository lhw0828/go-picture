package logic

import (
	"context"

	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSpaceMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceMembersLogic {
	return &ListSpaceMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间成员列表
func (l *ListSpaceMembersLogic) ListSpaceMembers(in *space.ListSpaceMembersRequest) (*space.ListSpaceMembersResponse, error) {
	// todo: add your logic here and delete this line

	return &space.ListSpaceMembersResponse{}, nil
}
