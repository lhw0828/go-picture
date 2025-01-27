package filemanager

import (
	"context"
	"io"
)

type PictureInfo struct {
	Url          string
	ThumbnailUrl string
	Size         int64
	Width        int32
	Height       int32
	Scale        float64
	Format       string
}

type FileManager interface {
	UploadPicture(ctx context.Context, file io.Reader, filename string, spaceId int64, userId int64) (*PictureInfo, error)
	DeletePicture(ctx context.Context, url string) error
}