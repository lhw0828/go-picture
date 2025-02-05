package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/errorx"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建空间
func NewCreateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSpaceLogic {
	return &CreateSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSpaceLogic) CreateSpace(req *types.CreateSpaceReq, userId int64) (resp *types.SpaceInfo, err error) {
	// 1. 参数校验
	if req == nil {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数不能为空")
	}

	// 2. 获取当前登录用户
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, errorx.NewCodeError(errorx.UserNotExist, "用户不存在")
	}

	// 3. 校验权限，非管理员只能创建普通级别的空间
	if req.SpaceLevel != 0 && userInfo.UserRole != "admin" {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "无权限创建指定级别的空间")
	}

	// 4. 调用 RPC 创建空间
	spaceInfo, err := l.svcCtx.SpaceRpc.CreateSpace(l.ctx, &space.CreateSpaceRequest{
		SpaceName:  req.SpaceName,
		SpaceType:  req.SpaceType,
		SpaceLevel: req.SpaceLevel,
		UserId:     userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SpaceInfo{
		Id:         spaceInfo.Id,
		SpaceName:  spaceInfo.SpaceName,
		SpaceType:  spaceInfo.SpaceType,
		SpaceLevel: spaceInfo.SpaceLevel,
		MaxSize:    spaceInfo.MaxSize,
		MaxCount:   spaceInfo.MaxCount,
		TotalSize:  spaceInfo.TotalSize,
		TotalCount: spaceInfo.TotalCount,
		UserId:     spaceInfo.UserId,
		CreateTime: spaceInfo.CreateTime,
		UpdateTime: spaceInfo.UpdateTime,
	}, nil
}
