package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	request "picture/common/types"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取空间信息
func NewGetSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceLogic {
	return &GetSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSpaceLogic) GetSpace(req *request.GetRequest, userId int64) (*types.SpaceInfo, error) {
	resp, err := l.svcCtx.SpaceRpc.GetSpace(l.ctx, &space.GetSpaceRequest{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	// 转换为 API 响应格式
	return &types.SpaceInfo{
		Id:          resp.Id,
		SpaceName:   resp.SpaceName,
		SpaceType:   resp.SpaceType,
		SpaceLevel:  resp.SpaceLevel,
		MaxSize:     resp.MaxSize,
		MaxCount:    resp.MaxCount,
		TotalSize:   resp.TotalSize,
		TotalCount:  resp.TotalCount,
		UserId:      resp.UserId,
		CreateTime:  resp.CreateTime,
		UpdateTime:  resp.UpdateTime,
		Permissions: []string{}, // 添加空的权限列表
	}, nil
}
