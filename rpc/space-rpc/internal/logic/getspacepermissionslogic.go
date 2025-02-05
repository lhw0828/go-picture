package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpacePermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpacePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpacePermissionsLogic {
	return &GetSpacePermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间权限
func (l *GetSpacePermissionsLogic) GetSpacePermissions(in *space.GetSpacePermissionsRequest) (*space.GetSpacePermissionsResponse, error) {
	// 1. 参数校验
	if in == nil || in.SpaceId <= 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 2. 获取空间信息
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}

	// 3. 获取用户信息
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: in.UserId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "用户不存在")
	}

	// 4. 获取权限列表
	permissions := make([]string, 0)

	// 4.1 如果是管理员，拥有所有权限
	if userInfo.UserRole == "admin" {
		permissions = append(permissions,
			"space:view",     // 查看空间
			"space:edit",     // 编辑空间
			"space:delete",   // 删除空间
			"space:upload",   // 上传文件
			"space:download", // 下载文件
			"space:share",    // 分享文件
			"space:manage",   // 空间管理
		)
		return &space.GetSpacePermissionsResponse{
			Permissions: permissions,
		}, nil
	}

	// 4.2 如果是空间所有者，拥有除管理权限外的所有权限
	if spaceInfo.UserId == userInfo.Id {
		permissions = append(permissions,
			"space:view",     // 查看空间
			"space:edit",     // 编辑空间
			"space:delete",   // 删除空间
			"space:upload",   // 上传文件
			"space:download", // 下载文件
			"space:share",    // 分享文件
		)
		return &space.GetSpacePermissionsResponse{
			Permissions: permissions,
		}, nil
	}

	// 4.3 如果是团队空间，查询用户在空间中的权限
	if spaceInfo.SpaceType == 1 {
		permissions, err = l.svcCtx.SpaceDao.GetUserPermissions(l.ctx, in.SpaceId, in.UserId)
		if err != nil {
			return nil, err
		}
	}

	// 4.4 如果是私有空间且不是所有者，只有查看权限
	if spaceInfo.SpaceType == 0 {
		permissions = append(permissions, "space:view")
	}

	return &space.GetSpacePermissionsResponse{
		Permissions: permissions,
	}, nil
}
