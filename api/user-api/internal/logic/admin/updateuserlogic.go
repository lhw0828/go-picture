package admin

import (
	"context"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/common/errorx"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserUpdateReq) (resp *types.BaseResp, err error) {
	// 参数校验
	if req.Id == 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "用户ID不能为空")
	}

	// 调用 RPC 服务
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UserUpdateRequest{
		Id:          req.Id,
		UserName:    req.UserName,
		UserAvatar:  req.UserAvatar,
		UserProfile: req.UserProfile,
	})
	if err != nil {
		l.Logger.Errorf("UpdateUser failed: %v", err)
		return nil, err
	}

	resp = &types.BaseResp{
		Code:    0,
		Message: "更新成功",
	}
	return resp, nil
}
