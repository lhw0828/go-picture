package model

import (
	"time"
)

type SpaceMember struct {
	Id         int64     `db:"id"`
	SpaceId    int64     `db:"spaceId"`
	UserId     int64     `db:"userId"`
	SpaceRole  string    `db:"spaceRole"`
	CreateTime time.Time `db:"createTime"`
	UpdateTime time.Time `db:"updateTime"`
}
