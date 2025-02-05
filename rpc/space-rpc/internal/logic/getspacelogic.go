package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

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
func (l *GetSpaceLogic) GetSpace(in *space.GetSpaceRequest) (*space.SpaceInfo, error) {
	// 1. 参数校验
	if in == nil || in.Id <= 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 2. 获取当前用户信息并校验管理员权限
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: in.UserId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil || userInfo.UserRole != "admin" {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "仅管理员可查看")
	}

	// 3. 查询空间信息
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}

	// 4. 转换为响应对象
	return &space.SpaceInfo{
		Id:         spaceInfo.Id,
		SpaceName:  spaceInfo.SpaceName,
		SpaceType:  spaceInfo.SpaceType,
		SpaceLevel: spaceInfo.SpaceLevel,
		MaxSize:    spaceInfo.MaxSize,
		MaxCount:   spaceInfo.MaxCount,
		TotalSize:  spaceInfo.TotalSize,
		TotalCount: spaceInfo.TotalCount,
		UserId:     spaceInfo.UserId,
		CreateTime: spaceInfo.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: spaceInfo.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
