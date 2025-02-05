package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserRequest) (*user.BaseResponse, error) {
	if in == nil || in.Id == 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	err := l.svcCtx.UserDao.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.BaseResponse{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
