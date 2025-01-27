package logic

import (
	"context"

	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePictureLogic {
	return &DeletePictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除图片
func (l *DeletePictureLogic) DeletePicture(in *picture.DeletePictureRequest) (*picture.DeletePictureResponse, error) {
	// todo: add your logic here and delete this line

	return &picture.DeletePictureResponse{}, nil
}
