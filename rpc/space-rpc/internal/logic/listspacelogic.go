package logic

import (
	"context"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSpaceLogic {
	return &ListSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间列表
func (l *ListSpaceLogic) ListSpace(in *space.ListSpaceRequest) (*space.ListSpaceResponse, error) {
	// 1. 参数校验
	if in == nil {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}
	if in.Current <= 0 {
		in.Current = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}

	// 2. 获取当前用户信息并校验管理员权限
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: in.UserId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil || userInfo.UserRole != "admin" {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "仅管理员可查看")
	}

	// 3. 查询总数
	total, err := l.svcCtx.SpaceDao.Count(l.ctx, in.SpaceName, in.SpaceType)
	if err != nil {
		return nil, err
	}

	// 4. 分页查询
	spaces, err := l.svcCtx.SpaceDao.List(l.ctx, in.Current, in.PageSize, in.SpaceName, in.SpaceType)
	if err != nil {
		return nil, err
	}

	// 5. 转换为响应对象
	spaceList := make([]*space.SpaceInfo, 0, len(spaces))
	for _, s := range spaces {
		spaceList = append(spaceList, &space.SpaceInfo{
			Id:         s.Id,
			SpaceName:  s.SpaceName,
			SpaceType:  s.SpaceType,
			SpaceLevel: s.SpaceLevel,
			MaxSize:    s.MaxSize,
			MaxCount:   s.MaxCount,
			TotalSize:  s.TotalSize,
			TotalCount: s.TotalCount,
			UserId:     s.UserId,
			CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &space.ListSpaceResponse{
		List:     spaceList,
		Total:    total,
		Current:  in.Current,
		PageSize: in.PageSize,
	}, nil
}
