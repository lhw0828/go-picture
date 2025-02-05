package admin

import (
	"context"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
