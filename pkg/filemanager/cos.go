package filemanager

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"golang.org/x/image/draw"
)

type CosManager struct {
	client *cos.Client
	bucket string
}

func NewCosManager(client *cos.Client, bucket string) *CosManager {
	return &CosManager{
		client: client,
		bucket: bucket,
	}
}

func (m *CosManager) UploadPicture(ctx context.Context, file io.Reader, filename string, spaceId int64, userId int64) (*PictureInfo, error) {
	// 按照空间划分目录
	var uploadPathPrefix string
	if spaceId == 0 {
		// 公共图库
		uploadPathPrefix = fmt.Sprintf("public/%d", userId)
	} else {
		// 空间
		uploadPathPrefix = fmt.Sprintf("space/%d", spaceId)
	}

	// 生成唯一文件名
	ext := filepath.Ext(filename)
	uniqueName := fmt.Sprintf("%s_%s%s",
		time.Now().Format("20060102"),
		uuid.New().String()[:8],
		ext,
	)

	// 生成文件路径
	originalKey := path.Join(uploadPathPrefix, "images", uniqueName)
	thumbnailKey := path.Join(uploadPathPrefix, "thumbnails", uniqueName)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %v", err)
	}

	// 分析图片信息
	img, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("解析图片失败: %v", err)
	}

	// 生成缩略图
	thumbnailData, err := m.generateThumbnail(data)
	if err != nil {
		return nil, fmt.Errorf("生成缩略图失败: %v", err)
	}

	// 上传原图
	_, err = m.client.Object.Put(ctx, originalKey, bytes.NewReader(data), nil)
	if err != nil {
		return nil, fmt.Errorf("上传原图失败: %v", err)
	}

	// 上传缩略图
	_, err = m.client.Object.Put(ctx, thumbnailKey, bytes.NewReader(thumbnailData), nil)
	if err != nil {
		return nil, fmt.Errorf("上传缩略图失败: %v", err)
	}

	// 获取图片访问 URL
	imgURL := m.client.Object.GetObjectURL(originalKey).String()
	thumbnailURL := m.client.Object.GetObjectURL(thumbnailKey).String()

	return &PictureInfo{
		Url:          imgURL,
		ThumbnailUrl: thumbnailURL,
		Size:         int64(len(data)),
		Width:        int32(img.Width),
		Height:       int32(img.Height),
		Scale:        float64(img.Width) / float64(img.Height),
		Format:       format,
	}, nil
}

func (m *CosManager) generateThumbnail(data []byte) ([]byte, error) {
	// 解码原图
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 计算缩略图尺寸
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	ratio := float64(width) / float64(height)

	var newWidth, newHeight int
	if ratio > 1 {
		newWidth = 300
		newHeight = int(float64(300) / ratio)
	} else {
		newHeight = 300
		newWidth = int(float64(300) * ratio)
	}

	// 生成缩略图
	thumbnail := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.NearestNeighbor.Scale(thumbnail, thumbnail.Rect, img, bounds, draw.Over, nil)

	// 编码为 JPEG
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, thumbnail, &jpeg.Options{Quality: 80})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *CosManager) DeletePicture(ctx context.Context, url string) error {
	return nil
}
