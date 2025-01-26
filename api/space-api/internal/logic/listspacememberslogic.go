package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取空间成员列表
func NewListSpaceMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceMembersLogic {
	return &ListSpaceMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSpaceMembersLogic) ListSpaceMembers() (resp *types.ListSpaceMembersResp, err error) {
	// todo: add your logic here and delete this line

	return
}
