package logic

import (
	context "context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/constants"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceLogic {
	return &GetSpaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSpaceLogic) GetSpace(req *types.GetSpaceReq) (resp *types.GetSpaceResp, err error) {
	// 打印请求参数
	l.Logger.Infof("获取空间信息请求参数: %+v", req)

	if req == nil || req.Id <= 0 {
		return nil, constants.NewCodeError(constants.ParamError, constants.ParamErrorMsg)
	}

	// 调用 RPC 获取空间信息
	spaceInfo, err := l.svcCtx.SpaceRpc.GetSpace(l.ctx, &space.GetSpaceRequest{
		Id: req.Id,
	})
	if spaceInfo == nil {
		return nil, constants.NewCodeError(constants.SpaceNotExist, constants.SpaceNotExistMsg)
	}
	if err != nil {
		l.Logger.Errorf("获取空间信息失败: %v", err)
		return nil, constants.NewCodeError(constants.SystemErr, constants.SystemErrMsg)
	}

	return &types.GetSpaceResp{
		Id:         spaceInfo.Id,
		SpaceName:  spaceInfo.SpaceName,
		SpaceType:  spaceInfo.SpaceType,
		SpaceLevel: spaceInfo.SpaceLevel,
		MaxSize:    spaceInfo.MaxSize,
		MaxCount:   spaceInfo.MaxCount,
		TotalSize:  spaceInfo.TotalSize,
		TotalCount: spaceInfo.TotalCount,
		UserId:     spaceInfo.UserId,
		CreateTime: spaceInfo.CreateTime,
		UpdateTime: spaceInfo.UpdateTime,
	}, nil
}
