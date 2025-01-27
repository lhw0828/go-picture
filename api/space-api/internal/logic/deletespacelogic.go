package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除团队空间
func NewDeleteSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSpaceLogic {
	return &DeleteSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSpaceLogic) DeleteSpace() (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
