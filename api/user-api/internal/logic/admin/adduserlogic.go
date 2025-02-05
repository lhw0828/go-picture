package admin

import (
	"context"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户
func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UserAddReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	return
}
