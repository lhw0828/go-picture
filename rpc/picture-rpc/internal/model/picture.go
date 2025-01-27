package model

import (
	"time"
)

type Picture struct {
	Id            int64     `db:"id"`
	Url          string    `db:"url"`
	ThumbnailUrl string    `db:"thumbnail_url"`
	Name         string    `db:"name"`
	Introduction string    `db:"introduction"`
	Category     string    `db:"category"`
	Tags         string    `db:"tags"`
	PicSize      int64     `db:"pic_size"`
	PicWidth     int32     `db:"pic_width"`
	PicHeight    int32     `db:"pic_height"`
	PicScale     float64   `db:"pic_scale"`
	PicFormat    string    `db:"pic_format"`
	PicColor     string    `db:"pic_color"`
	UserId       int64     `db:"user_id"`
	SpaceId      int64     `db:"space_id"`
	ReviewStatus int32     `db:"review_status"`
	ReviewMsg    string    `db:"review_message"`
	ReviewerId   int64     `db:"reviewer_id"`
	CreateTime   time.Time `db:"create_time"`
	EditTime     time.Time `db:"edit_time"`
	UpdateTime   time.Time `db:"update_time"`
	IsDelete     int32     `db:"is_delete"`
}