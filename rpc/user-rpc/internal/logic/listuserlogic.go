package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUserByPage(in *user.UserQueryRequest) (*user.UserQueryResponse, error) {
	if in == nil || in.Current < 1 || in.PageSize < 1 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 计算偏移量
	offset := (in.Current - 1) * in.PageSize

	// 查询总数
	total, err := l.svcCtx.UserDao.Count(l.ctx)
	if err != nil {
		return nil, err
	}

	// 分页查询
	users, err := l.svcCtx.UserDao.FindByPage(l.ctx, offset, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应对象
	var records []*user.UserInfo
	for _, u := range users {
		records = append(records, &user.UserInfo{
			Id:          u.Id,
			UserAccount: u.UserAccount,
			UserName:    u.UserName.String,
			UserAvatar:  u.UserAvatar.String,
			UserProfile: u.UserProfile.String,
			UserRole:    u.UserRole,
		})
	}

	return &user.UserQueryResponse{
		Total:   total,
		Records: records,
	}, nil
}
