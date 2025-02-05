package admin

import (
	"context"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取用户列表
func NewListUserByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserByPageLogic {
	return &ListUserByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserByPageLogic) ListUserByPage(req *types.UserQueryReq) (resp *types.UserQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
