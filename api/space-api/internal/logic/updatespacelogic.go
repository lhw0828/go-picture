package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新空间信息
func NewUpdateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpaceLogic {
	return &UpdateSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSpaceLogic) UpdateSpace(req *types.UpdateSpaceReq, userId int64) (*types.BaseResp, error) {
	resp, err := l.svcCtx.SpaceRpc.UpdateSpace(l.ctx, &space.UpdateSpaceRequest{
		Id:         req.Id,
		SpaceName:  req.SpaceName,
		SpaceLevel: req.SpaceLevel,
		MaxSize:    req.MaxSize,
		MaxCount:   req.MaxCount,
		UserId:     userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseResp{
		Code: resp.Code,
		Msg:  resp.Msg,
	}, nil
}
