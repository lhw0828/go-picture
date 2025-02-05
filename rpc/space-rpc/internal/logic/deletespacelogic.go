package logic

import (
	"context"
	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSpaceLogic {
	return &DeleteSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSpaceLogic) DeleteSpace(in *space.DeleteSpaceRequest) (*space.BaseResponse, error) {
	// 1. 参数校验
	if in == nil || in.Id <= 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 2. 检查空间是否存在
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}

	// 3. 获取当前用户信息
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: in.UserId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, errorx.NewCodeError(errorx.UserNotExist, "用户不存在")
	}

	// 4. 权限校验（仅本人或管理员可删除）
	if spaceInfo.UserId != in.UserId && userInfo.UserRole != "admin" {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "无权限删除该空间")
	}

	// 5. 删除空间（软删除）
	err = l.svcCtx.SpaceDao.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &space.BaseResponse{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
