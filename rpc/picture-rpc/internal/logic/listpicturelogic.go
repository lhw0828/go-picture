package logic

import (
	"context"

	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPictureLogic {
	return &ListPictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取图片列表
func (l *ListPictureLogic) ListPicture(in *picture.ListPictureRequest) (*picture.ListPictureResponse, error) {
	// todo: add your logic here and delete this line

	return &picture.ListPictureResponse{}, nil
}
