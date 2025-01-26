package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSpaceMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加空间成员
func NewAddSpaceMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSpaceMemberLogic {
	return &AddSpaceMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSpaceMemberLogic) AddSpaceMember(req *types.AddSpaceMemberReq) (resp *types.AddSpaceMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
