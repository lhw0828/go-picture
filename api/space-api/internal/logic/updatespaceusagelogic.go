package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpaceUsageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新空间使用容量
func NewUpdateSpaceUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpaceUsageLogic {
	return &UpdateSpaceUsageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSpaceUsageLogic) UpdateSpaceUsage(req *types.UpdateSpaceUsageReq) (resp *types.UpdateSpaceUsageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
