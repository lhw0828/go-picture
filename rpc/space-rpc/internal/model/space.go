package model

import (
	"database/sql"
	"time"
)

type Space struct {
	Id         int64         `db:"id"`
	SpaceName  string        `db:"spaceName"`
	SpaceLevel int           `db:"spaceLevel"` // 0-普通版 1-专业版 2-旗舰版
	SpaceType  int           `db:"spaceType"`  // 0-私有 1-团队
	MaxSize    int64         `db:"maxSize"`    // 最大总大小
	MaxCount   int64         `db:"maxCount"`   // 最大数量
	TotalSize  int64         `db:"totalSize"`  // 当前总大小
	TotalCount int64         `db:"totalCount"` // 当前数量
	UserId     int64         `db:"userId"`
	CreateTime time.Time     `db:"createTime"`
	EditTime   time.Time     `db:"editTime"`
	UpdateTime time.Time     `db:"updateTime"`
	IsDelete   sql.NullInt32 `db:"isDelete"`
}
