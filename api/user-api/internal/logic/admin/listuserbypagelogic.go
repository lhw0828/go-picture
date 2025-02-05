package admin

import (
	"context"

	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/rpc/user-rpc/user"

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
	// 参数校验
	if req.Current <= 0 {
		req.Current = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 调用 RPC 服务
	res, err := l.svcCtx.UserRpc.ListUserByPage(l.ctx, &user.UserQueryRequest{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserRole: req.UserRole,
	})
	if err != nil {
		l.Logger.Errorf("ListUserByPage failed: %v", err)
		return nil, err
	}

	// 转换响应数据
	var records []types.UserVO
	for _, u := range res.Records {
		records = append(records, types.UserVO{
			Id:          u.Id,
			UserAccount: u.UserAccount,
			UserName:    u.UserName,
			UserAvatar:  u.UserAvatar,
			UserProfile: u.UserProfile,
			UserRole:    u.UserRole,
			CreateTime:  u.CreateTime,
		})
	}

	resp = &types.UserQueryResp{
		Total:   res.Total,
		Records: records,
	}
	return resp, nil
}
