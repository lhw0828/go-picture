package logic

import (
	"context"
	"time"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceLogic {
	return &GetSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间信息
func (l *GetSpaceLogic) GetSpace(in *space.GetSpaceRequest) (*space.GetSpaceResponse, error) {
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.Id)
	if err != nil {
		l.Logger.Errorf("Find space error: %v", err)
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewDefaultError("空间不存在")
	}

	return &space.GetSpaceResponse{
		Id:         spaceInfo.Id,
		SpaceName:  spaceInfo.SpaceName,
		SpaceLevel: int32(spaceInfo.SpaceLevel), // 转换为 int32
		SpaceType:  int32(spaceInfo.SpaceType),  // 转换为 int32
		MaxSize:    spaceInfo.MaxSize,
		MaxCount:   spaceInfo.MaxCount,
		TotalSize:  spaceInfo.TotalSize,
		TotalCount: spaceInfo.TotalCount,
		UserId:     spaceInfo.UserId,
		CreateTime: spaceInfo.CreateTime.Format(time.RFC3339),
		UpdateTime: spaceInfo.UpdateTime.Format(time.RFC3339),
	}, nil
}
