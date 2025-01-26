package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/errorx"
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

func (l *GetSpaceLogic) GetSpace(req *types.GetSpaceReq) (resp *types.GetSpaceResp, err error) {
	// 打印请求参数
	l.Logger.Infof("获取空间信息请求参数: %+v", req)

	if req == nil || req.Id <= 0 {
		return nil, errorx.NewDefaultError("无效的空间ID")
	}

	// 调用 RPC 获取空间信息
	spaceInfo, err := l.svcCtx.SpaceRpc.GetSpace(l.ctx, &space.GetSpaceRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("获取空间信息失败")
	}
	if spaceInfo == nil {
		return nil, errorx.NewDefaultError("空间不存在")
	}

	// 转换空间级别
	var spaceLevel string
	switch spaceInfo.SpaceLevel {
	case 1:
		spaceLevel = "pro"
	case 2:
		spaceLevel = "premium"
	default:
		spaceLevel = "normal"
	}

	// 转换空间类型
	spaceType := "private"
	if spaceInfo.SpaceType == 1 {
		spaceType = "team"
	}

	return &types.GetSpaceResp{
		Id:         spaceInfo.Id,
		SpaceName:  spaceInfo.SpaceName,
		SpaceType:  spaceType,
		SpaceLevel: spaceLevel,
		MaxSize:    spaceInfo.MaxSize,
		MaxCount:   spaceInfo.MaxCount,
		TotalSize:  spaceInfo.TotalSize,
		TotalCount: spaceInfo.TotalCount,
		UserId:     spaceInfo.UserId,
		CreateTime: spaceInfo.CreateTime,
		UpdateTime: spaceInfo.UpdateTime,
	}, nil
}
