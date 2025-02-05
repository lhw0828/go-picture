package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取空间列表
func NewListSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceLogic {
	return &ListSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSpaceLogic) ListSpace(req *types.ListReq) (resp *types.ListSpaceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
