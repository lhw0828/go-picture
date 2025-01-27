package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建空间
func NewCreateSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSpaceLogic {
	return &CreateSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSpaceLogic) CreateSpace(req *types.CreateSpaceReq) (resp *types.CreateSpaceResp, err error) {
	l.Logger.Infof("Create space request: %+v", req)

	// 从上下文获取用户ID
	userId := l.ctx.Value("userId")
	if userId == nil {
		l.Logger.Error("Failed to get userId from context")
		return nil, fmt.Errorf("获取用户ID失败")
	}
	l.Logger.Infof("UserId from context: %v, type: %T", userId, userId)

	var uid int64
	switch v := userId.(type) {
	case json.Number:
		uid, err = v.Int64()
		if err != nil {
			l.Logger.Errorf("Convert userId error: %v", err)
			return nil, fmt.Errorf("用户ID格式错误")
		}
	default:
		l.Logger.Errorf("Invalid userId type: %T", userId)
		return nil, fmt.Errorf("用户ID类型错误")
	}

	l.Logger.Infof("Calling SpaceRpc.CreateSpace with userId: %d", uid)
	// 调用 RPC 服务
	res, err := l.svcCtx.SpaceRpc.CreateSpace(l.ctx, &space.CreateSpaceRequest{
		SpaceName:  req.SpaceName,
		SpaceType:  req.SpaceType,
		SpaceLevel: req.SpaceLevel,
		UserId:     uid,
	})
	if err != nil {
		l.Logger.Errorf("Create space error: %v", err)
		return nil, err
	}

	l.Logger.Infof("Space created successfully with id: %d", res.Id)
	return &types.CreateSpaceResp{
		Id: res.Id,
	}, nil
}
