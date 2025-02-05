package logic

import (
	"context"
	"time"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSpaceMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceMembersLogic {
	return &ListSpaceMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间成员列表
func (l *ListSpaceMembersLogic) ListSpaceMembers(in *space.ListSpaceMembersRequest) (*space.ListSpaceMembersResponse, error) {
	// 检查空间是否存在
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewCodeError(errorx.SpaceNotExist, "空间不存在")
	}

	// 获取成员列表
	members, err := l.svcCtx.SpaceMemberDao.FindBySpaceId(in.SpaceId)
	if err != nil {
		l.Logger.Errorf("Find space members error: %v", err)
		return nil, err
	}

	// 构建响应
	var respMembers []*space.SpaceMember
	for _, member := range members {
		// 获取用户信息
		userInfo, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdRequest{
			Id: member.UserId,
		})
		if err != nil {
			l.Logger.Errorf("Get user info error: %v", err)
			continue
		}

		respMembers = append(respMembers, &space.SpaceMember{
			Id:         member.Id,
			SpaceId:    member.SpaceId,
			UserId:     member.UserId,
			UserName:   userInfo.UserName,
			UserAvatar: userInfo.UserAvatar,
			Role:       member.SpaceRole,
			JoinTime:   member.CreateTime.Format(time.RFC3339),
		})
	}

	return &space.ListSpaceMembersResponse{
		Members: respMembers,
	}, nil
}
