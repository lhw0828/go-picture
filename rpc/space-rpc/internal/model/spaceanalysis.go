package model

import (
	"time"
)

type SpaceUsageRecord struct {
	Id         int64     `db:"id"`
	SpaceId    int64     `db:"spaceId"`
	FileSize   int64     `db:"fileSize"`
	FileType   string    `db:"fileType"`
	Operation  string    `db:"operation"`
	CreateTime time.Time `db:"createTime"`
}

type FileTypeStats struct {
	FileType  string `db:"fileType"`
	TotalSize int64  `db:"totalSize"`
	FileCount int64  `db:"fileCount"`
}

type DailyUsage struct {
	Date  string `db:"date"`
	Usage int64  `db:"usage"`
}
