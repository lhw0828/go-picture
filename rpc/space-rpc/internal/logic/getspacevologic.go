package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceVOLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpaceVOLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceVOLogic {
	return &GetSpaceVOLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSpaceVOLogic) GetSpaceVO(in *space.GetSpaceVORequest) (*space.SpaceVO, error) {
	// 1. 参数校验
	if in == nil || in.Id <= 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 2. 获取空间信息
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}

	// 3. 获取创建者信息
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: spaceInfo.UserId})
	if err != nil {
		return nil, err
	}

	// 4. 获取当前用户权限
	permissions, err := l.svcCtx.SpaceDao.GetUserPermissions(l.ctx, in.Id, in.UserId)
	if err != nil {
		return nil, err
	}

	// 5. 检查是否有查看权限
	hasViewPermission := false
	for _, perm := range permissions {
		if perm == "space:view" {
			hasViewPermission = true
			break
		}
	}
	if !hasViewPermission {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "无权查看该空间")
	}

	// 6. 组装返回数据
	return &space.SpaceVO{
		SpaceInfo: &space.SpaceInfo{
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
		},
		User: &space.UserInfo{
			Id:         userInfo.Id,
			Username:   userInfo.UserName,
			UserAvatar: userInfo.UserAvatar,
			UserRole:   userInfo.UserRole,
		},
		Permissions: permissions,
	}, nil
}
