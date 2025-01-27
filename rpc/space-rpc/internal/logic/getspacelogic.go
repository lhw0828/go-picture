package logic

import (
	"context"
	"time"

	"picture/common/constants"
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
	// 打印请求参数
	l.Logger.Infof("获取空间信息请求: %+v", in)

	// 参数校验
	if in == nil || in.Id <= 0 {
		return nil, constants.NewCodeError(constants.ParamError, constants.ParamErrorMsg)
	}

	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.Id)
	if spaceInfo == nil {
		l.Logger.Infof("空间不存在, id: %d", in.Id)
		return nil, constants.NewCodeError(constants.SpaceNotExist, constants.SpaceNotExistMsg)
	}
	if err != nil {
		l.Logger.Errorf("获取空间信息失败，Find space error: %v", err)
		return nil, constants.NewCodeError(constants.GetSpaceFailed, constants.GetSpaceFailedMsg)
	}

	return &space.GetSpaceResponse{
		Id:         spaceInfo.Id,
		SpaceName:  spaceInfo.SpaceName,
		SpaceLevel: int32(spaceInfo.SpaceLevel),
		SpaceType:  int32(spaceInfo.SpaceType),
		MaxSize:    spaceInfo.MaxSize,
		MaxCount:   spaceInfo.MaxCount,
		TotalSize:  spaceInfo.TotalSize,
		TotalCount: spaceInfo.TotalCount,
		UserId:     spaceInfo.UserId,
		CreateTime: spaceInfo.CreateTime.Format(time.RFC3339),
		UpdateTime: spaceInfo.UpdateTime.Format(time.RFC3339),
	}, nil
}
