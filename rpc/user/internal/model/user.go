package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int64          `db:"id"`
	UserAccount  string         `db:"userAccount"`
	UserPassword string         `db:"userPassword"`
	UserName     string         `db:"userName"`
	UserAvatar   sql.NullString `db:"userAvatar"`  // 修改为 sql.NullString
	UserProfile  sql.NullString `db:"userProfile"` // 修改为 sql.NullString
	UserRole     string         `db:"userRole"`
	CreateTime   time.Time      `db:"createTime"`
	UpdateTime   time.Time      `db:"updateTime"`
	IsDelete     sql.NullInt32  `db:"isDelete"`
}
