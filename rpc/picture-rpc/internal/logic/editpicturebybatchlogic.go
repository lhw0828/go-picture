package logic

import (
	"context"

	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPictureByBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditPictureByBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPictureByBatchLogic {
	return &EditPictureByBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量编辑图片
func (l *EditPictureByBatchLogic) EditPictureByBatch(in *picture.EditPictureByBatchRequest) (*picture.EditPictureByBatchResponse, error) {
	// todo: add your logic here and delete this line

	return &picture.EditPictureByBatchResponse{}, nil
}
