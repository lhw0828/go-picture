package logic

import (
	"context"

	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewPictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReviewPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewPictureLogic {
	return &ReviewPictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 图片审核
func (l *ReviewPictureLogic) ReviewPicture(in *picture.ReviewPictureRequest) (*picture.ReviewPictureResponse, error) {
	// todo: add your logic here and delete this line

	return &picture.ReviewPictureResponse{}, nil
}
