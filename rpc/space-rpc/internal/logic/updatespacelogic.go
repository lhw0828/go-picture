package logic

import (
	"context"

	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpaceLogic {
	return &UpdateSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新空间信息
func (l *UpdateSpaceLogic) UpdateSpace(in *space.UpdateSpaceRequest) (*space.BaseResponse, error) {
	// todo: add your logic here and delete this line

	return &space.BaseResponse{}, nil
}
