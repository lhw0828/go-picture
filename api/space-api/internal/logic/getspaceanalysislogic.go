package logic

import (
	"context"

	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpaceAnalysisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取空间分析数据
func NewGetSpaceAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpaceAnalysisLogic {
	return &GetSpaceAnalysisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSpaceAnalysisLogic) GetSpaceAnalysis(req *types.SpaceAnalysisReq) (resp *types.SpaceAnalysis, err error) {
	// todo: add your logic here and delete this line

	return
}
