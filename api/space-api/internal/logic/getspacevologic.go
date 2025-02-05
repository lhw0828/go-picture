package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/errorx"
	request "picture/common/types"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceVOLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpaceVOLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceVOLogic {
	return &GetSpaceVOLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSpaceVOLogic) GetSpaceVO(req *request.GetRequest, userId int64) (*types.SpaceVO, error) {
	// 1. 获取空间基本信息
	resp, err := l.svcCtx.SpaceRpc.GetSpace(l.ctx, &space.GetSpaceRequest{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	// 2. 获取空间创建者信息
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{
		Id: resp.UserId, // 注意：这里使用空间的创建者ID
	})
	if err != nil {
		return nil, err
	}

	// 3. 获取当前用户权限（使用当前登录用户ID）
	permResp, err := l.svcCtx.SpaceRpc.GetSpacePermissions(l.ctx, &space.GetSpacePermissionsRequest{
		SpaceId: req.Id,
		UserId:  userId, // 使用当前登录用户ID查询权限
	})
	if err != nil {
		return nil, err
	}

	// 4. 检查是否有查看权限
	hasViewPermission := false
	for _, perm := range permResp.Permissions {
		if perm == "space:view" {
			hasViewPermission = true
			break
		}
	}
	if !hasViewPermission {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "无权查看该空间")
	}

	return &types.SpaceVO{
		SpaceInfo: types.SpaceInfo{
			Id:         resp.Id,
			SpaceName:  resp.SpaceName,
			SpaceType:  resp.SpaceType,
			SpaceLevel: resp.SpaceLevel,
			MaxSize:    resp.MaxSize,
			MaxCount:   resp.MaxCount,
			TotalSize:  resp.TotalSize,
			TotalCount: resp.TotalCount,
			UserId:     resp.UserId,
			CreateTime: resp.CreateTime,
			UpdateTime: resp.UpdateTime,
		},
		UserInfo: types.UserInfo{
			Id:         userInfo.Id,
			UserName:   userInfo.UserName,
			UserAvatar: userInfo.UserAvatar,
			UserRole:   userInfo.UserRole,
		},
		Permissions: permResp.Permissions,
	}, nil
}
