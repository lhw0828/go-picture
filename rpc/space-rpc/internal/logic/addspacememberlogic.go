package logic

import (
	"context"
	"time"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/model"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSpaceMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSpaceMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSpaceMemberLogic {
	return &AddSpaceMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加空间成员
func (l *AddSpaceMemberLogic) AddSpaceMember(in *space.AddSpaceMemberRequest) (*space.AddSpaceMemberResponse, error) {
	// 检查空间是否存在
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewDefaultError("空间不存在")
	}

	// 检查是否为团队空间
	if spaceInfo.SpaceType != 1 {
		return nil, errorx.NewDefaultError("非团队空间不能添加成员")
	}

	// 检查用户是否已经是成员
	existMember, err := l.svcCtx.SpaceMemberDao.FindBySpaceIdAndUserId(in.SpaceId, in.UserId)
	if err != nil {
		return nil, err
	}
	if existMember != nil {
		return nil, errorx.NewDefaultError("用户已是空间成员")
	}

	// 添加成员
	member := &model.SpaceMember{
		SpaceId:    in.SpaceId,
		UserId:     in.UserId,
		SpaceRole:  "viewer", // 默认为查看者角色
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	_, err = l.svcCtx.SpaceMemberDao.Insert(member)
	if err != nil {
		l.Logger.Errorf("Insert space member error: %v", err)
		return nil, errorx.NewDefaultError("添加成员失败")
	}

	return &space.AddSpaceMemberResponse{
		Success: true,
	}, nil
}
