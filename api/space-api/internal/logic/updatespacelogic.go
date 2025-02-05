package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新团队空间使用容量
func NewUpdateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpaceLogic {
	return &UpdateSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSpaceLogic) UpdateSpace(req *types.UpdateSpaceReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
