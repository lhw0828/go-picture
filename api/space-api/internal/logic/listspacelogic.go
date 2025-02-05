package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取空间列表
func NewListSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceLogic {
	return &ListSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSpaceLogic) ListSpace(req *types.ListReq, userId int64) (*types.ListSpaceResp, error) {
	resp, err := l.svcCtx.SpaceRpc.ListSpace(l.ctx, &space.ListSpaceRequest{
		Current:   req.Current,
		PageSize:  req.PageSize,
		SpaceName: req.SpaceName,
		SpaceType: req.SpaceType,
		UserId:    userId,
	})
	if err != nil {
		return nil, err
	}

	spaceList := make([]types.SpaceInfo, 0, len(resp.List))
	for _, s := range resp.List {
		spaceList = append(spaceList, types.SpaceInfo{
			Id:         s.Id,
			SpaceName:  s.SpaceName,
			SpaceType:  s.SpaceType,
			SpaceLevel: s.SpaceLevel,
			MaxSize:    s.MaxSize,
			MaxCount:   s.MaxCount,
			TotalSize:  s.TotalSize,
			TotalCount: s.TotalCount,
			UserId:     s.UserId,
			CreateTime: s.CreateTime,
			UpdateTime: s.UpdateTime,
		})
	}

	return &types.ListSpaceResp{
		List:     spaceList,
		Total:    resp.Total,
		Current:  resp.Current,
		PageSize: resp.PageSize,
	}, nil
}
