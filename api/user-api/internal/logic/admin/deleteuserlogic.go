package admin

import (
	"context"
	"strconv"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/common/errorx"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(id string) (resp *types.BaseResp, err error) {
	// 参数校验
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ParamError, "无效的用户ID")
	}

	// 调用 RPC 服务
	_, err = l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserRequest{
		Id: userId,
	})
	if err != nil {
		l.Logger.Errorf("DeleteUser failed: %v", err)
		return nil, err
	}

	resp = &types.BaseResp{
		Code:    0,
		Message: "删除成功",
	}
	return resp, nil
}
