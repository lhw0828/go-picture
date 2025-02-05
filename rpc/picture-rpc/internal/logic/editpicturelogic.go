package logic

import (
	"context"

	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPictureLogic {
	return &EditPictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 编辑图片
func (l *EditPictureLogic) EditPicture(in *picture.EditPictureRequest) (*picture.EditPictureResponse, error) {
	// todo: add your logic here and delete this line

	return &picture.EditPictureResponse{}, nil
}
