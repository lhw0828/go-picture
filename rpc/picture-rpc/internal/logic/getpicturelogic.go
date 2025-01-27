package logic

import (
	"context"

	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPictureLogic {
	return &GetPictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取图片信息
func (l *GetPictureLogic) GetPicture(in *picture.GetPictureRequest) (*picture.GetPictureResponse, error) {
	// todo: add your logic here and delete this line

	return &picture.GetPictureResponse{}, nil
}
