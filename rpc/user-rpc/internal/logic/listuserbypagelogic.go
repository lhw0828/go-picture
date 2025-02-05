package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/user-rpc/internal/svc"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserByPageLogic {
	return &ListUserByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserByPageLogic) ListUserByPage(in *user.UserQueryRequest) (*user.UserQueryResponse, error) {
	// 参数校验
	if in == nil {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 计算偏移量
	offset := (in.Current - 1) * in.PageSize

	// 查询总数
	total, err := l.svcCtx.UserDao.Count(l.ctx, in.UserRole)
	if err != nil {
		l.Logger.Errorf("Count users error: %v", err)
		return nil, err
	}

	// 查询用户列表
	users, err := l.svcCtx.UserDao.ListByPage(l.ctx, offset, in.PageSize, in.UserRole)
	if err != nil {
		l.Logger.Errorf("List users error: %v", err)
		return nil, err
	}

	// 转换响应数据
	var records []*user.UserInfo
	for _, u := range users {
		records = append(records, &user.UserInfo{
			Id:          u.Id,
			UserAccount: u.UserAccount,
			UserName:    u.UserName.String,
			UserAvatar:  u.UserAvatar.String,
			UserProfile: u.UserProfile.String,
			UserRole:    u.UserRole,
			CreateTime:  u.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &user.UserQueryResponse{
		Total:   total,
		Records: records,
	}, nil
}
