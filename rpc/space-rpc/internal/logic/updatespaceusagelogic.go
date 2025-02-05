package logic

import (
	"context"

	"picture/common/errorx"
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
	if err := l.validateRequest(in); err != nil {
		return nil, err
	}

	if err := l.svcCtx.SpaceDao.UpdateUsage(in.SpaceId, in.Size, in.Operation); err != nil {
		l.Error("更新空间使用量失败", logx.Field("error", err))
		return nil, errorx.NewError(errorx.SystemErr)
	}

	total, used, err := l.svcCtx.SpaceDao.GetUsage(in.SpaceId)
	if err != nil {
		l.Error("获取空间使用量失败", logx.Field("error", err))
		return nil, errorx.NewError(errorx.SystemErr)
	}

	return &space.UpdateSpaceUsageResponse{
		Success:           true,
		RemainingCapacity: total - used,
	}, nil
}

func (l *UpdateSpaceUsageLogic) validateRequest(in *space.UpdateSpaceUsageRequest) error {
	if in == nil || in.SpaceId <= 0 {
		return errorx.NewError(errorx.ParamError)
	}
	if in.Size <= 0 {
		return errorx.NewErrorWithMsg(errorx.ParamError, "无效的容量大小")
	}
	if in.Operation != "add" && in.Operation != "subtract" {
		return errorx.NewErrorWithMsg(errorx.ParamError, "无效的操作类型")
	}
	return nil
}
