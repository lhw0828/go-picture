package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/common/types"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除空间
func NewDeleteSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSpaceLogic {
	return &DeleteSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSpaceLogic) DeleteSpace(req *types.DeleteRequest, userId int64) (resp *space.BaseResponse, err error) {
	resp, err = l.svcCtx.SpaceRpc.DeleteSpace(l.ctx, &space.DeleteSpaceRequest{
		Id:     int64(req.Id),
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &space.BaseResponse{
		Code: resp.Code,
		Msg:  resp.Msg,
	}, err
}
