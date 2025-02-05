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

type UpdateSpaceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpaceLogic {
	return &UpdateSpaceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新空间信息
func (l *UpdateSpaceLogic) UpdateSpace(in *space.UpdateSpaceRequest) (*space.BaseResponse, error) {
	// 1. 参数校验
	if in == nil || in.Id <= 0 {
		return nil, errorx.NewCodeError(errorx.ParamError, "参数错误")
	}

	// 2. 检查空间是否存在
	spaceInfo, err := l.svcCtx.SpaceDao.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	if spaceInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}

	// 3. 获取当前用户信息并校验管理员权限
	userInfo, err := l.svcCtx.UserRpc.GetCurrentUser(l.ctx, &user.GetUserByIdRequest{Id: in.UserId})
	if err != nil {
		return nil, err
	}
	if userInfo == nil || userInfo.UserRole != "admin" {
		return nil, errorx.NewCodeError(errorx.ForbiddenErr, "仅管理员可更新空间")
	}

	// 4. 数据校验
	if err := l.validSpace(in); err != nil {
		return nil, err
	}

	// 5. 更新空间信息
	spaceInfo.SpaceName = in.SpaceName
	spaceInfo.SpaceLevel = in.SpaceLevel
	spaceInfo.MaxSize = in.MaxSize
	spaceInfo.MaxCount = in.MaxCount
	spaceInfo.UpdateTime = time.Now()

	// 6. 自动填充数据
	l.fillSpaceBySpaceLevel(spaceInfo)

	// 7. 执行更新
	err = l.svcCtx.SpaceDao.Update(l.ctx, spaceInfo)
	if err != nil {
		return nil, err
	}

	return &space.BaseResponse{
		Code: 0,
		Msg:  "更新空间成功",
	}, nil
}

// 数据校验
func (l *UpdateSpaceLogic) validSpace(in *space.UpdateSpaceRequest) error {
	if len(in.SpaceName) > 30 {
		return errorx.NewCodeError(errorx.ParamError, "空间名称过长")
	}
	if !isValidSpaceLevel(in.SpaceLevel) {
		return errorx.NewCodeError(errorx.ParamError, "空间级别不存在")
	}
	return nil
}

// 检查空间级别是否有效
func isValidSpaceLevel(level int32) bool {
	return level >= 0 && level <= 1 // 目前支持普通版(0)和专业版(1)
}

// 根据空间级别填充数据
func (l *UpdateSpaceLogic) fillSpaceBySpaceLevel(spaceInfo *model.Space) {
	if spaceInfo.SpaceLevel == 0 { // 普通版
		if spaceInfo.MaxSize == 0 {
			spaceInfo.MaxSize = 5 * 1024 * 1024 * 1024 // 5GB
		}
		if spaceInfo.MaxCount == 0 {
			spaceInfo.MaxCount = 1000
		}
	} else if spaceInfo.SpaceLevel == 1 { // 专业版
		if spaceInfo.MaxSize == 0 {
			spaceInfo.MaxSize = 10 * 1024 * 1024 * 1024 // 10GB
		}
		if spaceInfo.MaxCount == 0 {
			spaceInfo.MaxCount = 10000
		}
	}
	// ... 其他级别
}
