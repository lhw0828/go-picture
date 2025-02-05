package dao

import (
	"picture/rpc/space-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SpaceAnalysisDao struct {
	conn sqlx.SqlConn
}

func NewSpaceAnalysisDao(conn sqlx.SqlConn) *SpaceAnalysisDao {
	return &SpaceAnalysisDao{conn: conn}
}

func (d *SpaceAnalysisDao) GetFileStats(spaceId int64) (*model.FileTypeStats, error) {
	var stats model.FileTypeStats
	query := `
        SELECT COUNT(*) as fileCount,
               SUM(fileSize) as totalSize
        FROM space_file
        WHERE spaceId = ? AND isDelete = 0`
	err := d.conn.QueryRow(&stats, query, spaceId)
	return &stats, err
}

func (d *SpaceAnalysisDao) GetFileTypeDistribution(spaceId int64) ([]model.FileTypeStats, error) {
	var stats []model.FileTypeStats
	query := `
        SELECT fileType, 
               COUNT(*) as fileCount, 
               SUM(fileSize) as totalSize 
        FROM space_file 
        WHERE spaceId = ? AND isDelete = 0 
        GROUP BY fileType`
	err := d.conn.QueryRows(&stats, query, spaceId)
	return stats, err
}

func (d *SpaceAnalysisDao) GetUsageTrends(spaceId int64, days int) ([]model.DailyUsage, error) {
	var trends []model.DailyUsage
	query := `
        SELECT DATE(createTime) as date, 
               SUM(CASE WHEN operation = 'add' THEN fileSize 
                        WHEN operation = 'delete' THEN -fileSize 
                        ELSE 0 END) as usage 
        FROM space_usage_record 
        WHERE spaceId = ? 
        AND createTime >= DATE_SUB(CURDATE(), INTERVAL ? DAY) 
        GROUP BY DATE(createTime) 
        ORDER BY date`
	err := d.conn.QueryRows(&trends, query, spaceId, days)
	return trends, err
}
