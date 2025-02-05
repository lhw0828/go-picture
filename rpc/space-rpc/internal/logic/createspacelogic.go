package logic

import (
	"context"
	"database/sql"
	"time"

	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/model"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

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
func (l *CreateSpaceLogic) CreateSpace(in *space.CreateSpaceRequest) (*space.CreateSpaceResponse, error) {
	l.Logger.Info("=== Start creating space ===")
	l.Logger.Infof("Request params: %+v", in)

	// 参数校验
	if len(in.SpaceName) < 1 {
		return nil, errorx.NewCodeError(errorx.SpaceNameNotNull, "空间名称不能为空")
	}

	// 设置空间容量
	var maxSize int64
	switch in.SpaceLevel {
	case 0:
		maxSize = 1024 * 1024 * 1024 // 1GB
	case 1:
		maxSize = 5 * 1024 * 1024 * 1024 // 5GB
	case 2:
		maxSize = 10 * 1024 * 1024 * 1024 // 10GB
	default:
		return nil, errorx.NewCodeError(errorx.InvalidSpaceLevel, "无效的空间级别")
	}

	// 创建空间
	newSpace := &model.Space{
		SpaceName:  in.SpaceName,
		SpaceLevel: 0, // normal
		SpaceType:  0, // private
		MaxSize:    maxSize,
		MaxCount:   1000, // 默认1000张图片
		TotalSize:  0,
		TotalCount: 0,
		UserId:     in.UserId,
		CreateTime: time.Now(),
		EditTime:   time.Now(),
		UpdateTime: time.Now(),
		IsDelete:   sql.NullInt32{Int32: 0, Valid: true},
	}

	result, err := l.svcCtx.SpaceDao.Insert(newSpace)
	if err != nil {
		l.Logger.Errorf("Insert space error: %v", err)
		return nil, errorx.NewCodeError(errorx.CreateSpaceFailed, "创建空间失败")
	}

	spaceId, err := result.LastInsertId()
	if err != nil {
		l.Logger.Errorf("Get last insert id error: %v", err)
		return nil, errorx.NewCodeError(errorx.CreateSpaceFailed, "创建空间失败")
	}

	return &space.CreateSpaceResponse{
		Id: spaceId,
	}, nil
}
