package logic

import (
	"context"

	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceLogic {
	return &GetSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间信息
func (l *GetSpaceLogic) GetSpace(in *space.GetSpaceRequest) (*space.SpaceInfo, error) {
	// todo: add your logic here and delete this line

	return &space.SpaceInfo{}, nil
}
