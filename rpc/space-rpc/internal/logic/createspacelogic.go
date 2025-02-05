package logic

import (
	"context"
	"time"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/model"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
	"picture/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSpaceLogic {
	return &CreateSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建空间
func (l *CreateSpaceLogic) CreateSpace(in *space.CreateSpaceRequest) (*space.SpaceInfo, error) {
	// 1. 参数校验
	if in == nil || in.SpaceName == "" {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 2. 检查用户是否存在
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: in.UserId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, errorx.NewCodeError(errorx.UserNotExist, "用户不存在")
	}

	// 3. 检查用户是否已有同类型空间
	exists, err := l.svcCtx.SpaceDao.ExistsByUserIdAndType(l.ctx, in.UserId, in.SpaceType)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorx.NewCodeError(errorx.OnlyCreateOneSpaceEachType, "每个用户每类空间只能创建一个")
	}

	// 4. 创建空间
	now := time.Now()
	spaceModel := &model.Space{
		SpaceName:  in.SpaceName,
		SpaceType:  in.SpaceType,
		SpaceLevel: in.SpaceLevel,
		UserId:     in.UserId,
		CreateTime: now,
		EditTime:   now,
		UpdateTime: now,
	}

	// 5. 根据空间级别设置配额
	l.fillSpaceByLevel(spaceModel)

	// 6. 保存到数据库
	result, err := l.svcCtx.SpaceDao.Insert(l.ctx, spaceModel)
	if err != nil {
		return nil, err
	}

	// 获取插入记录的ID
	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	spaceModel.Id = lastId

	// 7. 如果是团队空间，创建空间成员记录
	if spaceModel.SpaceType == 1 {
		_, err = l.svcCtx.SpaceMemberDao.Insert(l.ctx, &model.SpaceMember{
			SpaceId:    spaceModel.Id,
			UserId:     in.UserId,
			SpaceRole:  "admin",
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
		if err != nil {
			return nil, err
		}
	}

	return &space.SpaceInfo{
		Id:         spaceModel.Id,
		SpaceName:  spaceModel.SpaceName,
		SpaceType:  spaceModel.SpaceType,
		SpaceLevel: spaceModel.SpaceLevel,
		MaxSize:    spaceModel.MaxSize,
		MaxCount:   spaceModel.MaxCount,
		TotalSize:  spaceModel.TotalSize,
		TotalCount: spaceModel.TotalCount,
		UserId:     spaceModel.UserId,
		CreateTime: spaceModel.CreateTime.Format(time.RFC3339),
		UpdateTime: spaceModel.UpdateTime.Format(time.RFC3339),
	}, nil
}

// fillSpaceByLevel 根据空间级别设置配额
func (l *CreateSpaceLogic) fillSpaceByLevel(space *model.Space) {
	switch space.SpaceLevel {
	case 0: // 普通版
		space.MaxSize = 1 << 30 // 1GB
		space.MaxCount = 1000   // 1000个文件
	case 1: // 专业版
		space.MaxSize = 10 << 30 // 10GB
		space.MaxCount = 10000   // 10000个文件
	case 2: // 旗舰版
		space.MaxSize = 100 << 30 // 100GB
		space.MaxCount = 100000   // 100000个文件
	default:
		space.MaxSize = 1 << 30 // 默认1GB
		space.MaxCount = 1000   // 默认1000个文件
	}
	// 初始化使用量
	space.TotalSize = 0
	space.TotalCount = 0
}
