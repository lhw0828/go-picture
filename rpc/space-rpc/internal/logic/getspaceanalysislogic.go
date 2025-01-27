package logic

import (
	"context"

	"picture/common/errorx"
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
func (l *GetSpaceAnalysisLogic) GetSpaceAnalysis(in *space.GetSpaceAnalysisRequest) (*space.GetSpaceAnalysisResponse, error) {
	// 获取文件统计信息
	fileStats, err := l.svcCtx.SpaceAnalysisDao.GetFileStats(in.SpaceId)
	if err != nil {
		return nil, err
	}

	// 获取文件类型分布
	typeStats, err := l.svcCtx.SpaceAnalysisDao.GetFileTypeDistribution(in.SpaceId)
	if err != nil {
		return nil, err
	}

	// 获取使用趋势
	var days int
	switch in.TimeRange {
	case "day":
		days = 1
	case "week":
		days = 7
	case "month":
		days = 30
	default:
		return nil, errorx.NewCodeError(errorx.InvalidTimeRange, "无效的时间范围")
	}

	trends, err := l.svcCtx.SpaceAnalysisDao.GetUsageTrends(in.SpaceId, days)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp := &space.GetSpaceAnalysisResponse{
		TotalFiles:           fileStats.FileCount,
		TotalSize:            fileStats.TotalSize,
		FileTypeDistribution: make(map[string]int64),
	}

	// 添加文件类型分布
	for _, stat := range typeStats {
		resp.FileTypeDistribution[stat.FileType] = stat.TotalSize
	}

	// 添加使用趋势
	for _, trend := range trends {
		resp.UsageTrends = append(resp.UsageTrends, &space.SpaceUsageTrend{
			Date:  trend.Date,
			Usage: trend.Usage,
		})
	}

	return resp, nil
}
