package logic

import (
	"context"

	"picture/common/constants"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpaceUsageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSpaceUsageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpaceUsageLogic {
	return &UpdateSpaceUsageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新空间使用容量
func (l *UpdateSpaceUsageLogic) UpdateSpaceUsage(in *space.UpdateSpaceUsageRequest) (*space.UpdateSpaceUsageResponse, error) {
	// 参数校验
	if in == nil || in.SpaceId <= 0 {
		return nil, constants.NewCodeError(constants.ParamError, constants.ParamErrorMsg)
	}
	if in.Size <= 0 {
		return nil, constants.NewCodeError(constants.ParamError, "无效的容量大小")
	}
	if in.Operation != "add" && in.Operation != "subtract" {
		return nil, constants.NewCodeError(constants.ParamError, "无效的操作类型")
	}

	// 更新使用容量
	err := l.svcCtx.SpaceDao.UpdateUsage(in.SpaceId, in.Size, in.Operation)
	if err != nil {
		l.Logger.Errorf("Update space usage error: %v", err)
		return nil, err
	}

	// 获取更新后的容量信息
	total, used, err := l.svcCtx.SpaceDao.GetUsage(in.SpaceId)
	if err != nil {
		return nil, err
	}

	return &space.UpdateSpaceUsageResponse{
		Success:           true,
		RemainingCapacity: total - used,
	}, nil
}
