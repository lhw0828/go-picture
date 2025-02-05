package logic

import (
	"context"

	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceAnalysisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpaceAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceAnalysisLogic {
	return &GetSpaceAnalysisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取空间分析数据
func (l *GetSpaceAnalysisLogic) GetSpaceAnalysis(in *space.GetSpaceAnalysisRequest) (*space.SpaceAnalysis, error) {
	// todo: add your logic here and delete this line

	return &space.SpaceAnalysis{}, nil
}
