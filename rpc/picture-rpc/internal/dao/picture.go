package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Picture struct {
	Id           int64
	Url          string
	ThumbnailUrl string
	Name         string
	Introduction string
	Category     string
	Tags         string
	PicSize      int64
	PicWidth     int32
	PicHeight    int32
	PicScale     float64
	PicFormat    string
	PicColor     string
	UserId       int64
	SpaceId      int64
	ReviewStatus int32
	ReviewMsg    string
	ReviewerId   int64
	CreateTime   time.Time
	EditTime     time.Time
	UpdateTime   time.Time
	IsDelete     int32
}

type PictureDao struct {
	db *sql.DB
}

func NewPictureDao(db *sql.DB) *PictureDao {
	return &PictureDao{
		db: db,
	}
}

func (d *PictureDao) Insert(ctx context.Context, picture *Picture) error {
	query := `INSERT INTO picture (url, thumbnailUrl, name, spaceId, picSize, picWidth, picHeight, 
		picScale, picFormat, reviewStatus, userId, createTime, editTime, updateTime)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), NOW())`

	result, err := d.db.ExecContext(ctx, query,
		picture.Url,
		picture.ThumbnailUrl,
		picture.Name,
		picture.SpaceId,
		picture.PicSize,
		picture.PicWidth,
		picture.PicHeight,
		picture.PicScale,
		picture.PicFormat,
		picture.ReviewStatus,
		picture.UserId,  // 添加 userId 参数
	)
	if err != nil {
		return fmt.Errorf("插入图片记录失败: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取插入ID失败: %v", err)
	}

	picture.Id = id
	return nil
}

func (d *PictureDao) GetById(ctx context.Context, id int64) (*Picture, error) {
	// 明确指定需要查询的字段，避免使用 SELECT *
	query := `SELECT id, url, thumbnail_url, name, introduction, category, tags, 
		pic_size, pic_width, pic_height, pic_scale, pic_format, pic_color, 
		user_id, space_id, review_status, review_message, reviewer_id, 
		create_time, edit_time, update_time, is_delete 
		FROM picture WHERE id = ? AND is_delete = 0`

	var picture Picture
	err := d.db.QueryRowContext(ctx, query, id).Scan(
		&picture.Id,
		&picture.Url,
		&picture.ThumbnailUrl,
		&picture.Name,
		&picture.Introduction,
		&picture.Category,
		&picture.Tags,
		&picture.PicSize,
		&picture.PicWidth,
		&picture.PicHeight,
		&picture.PicScale,
		&picture.PicFormat,
		&picture.PicColor,
		&picture.UserId,
		&picture.SpaceId,
		&picture.ReviewStatus,
		&picture.ReviewMsg,
		&picture.ReviewerId,
		&picture.CreateTime,
		&picture.EditTime,
		&picture.UpdateTime,
		&picture.IsDelete,
	)
	if err != nil {
		return nil, fmt.Errorf("查询图片失败: %v", err)
	}
	return &picture, nil
}

func (d *PictureDao) Update(ctx context.Context, picture *Picture) error {
	query := `UPDATE picture SET 
		name = ?, introduction = ?, category = ?, tags = ?,
		review_status = ?, review_message = ?, reviewer_id = ?,
		review_time = NOW(), edit_time = NOW(), update_time = NOW()
		WHERE id = ? AND is_delete = 0`

	result, err := d.db.ExecContext(ctx, query,
		picture.Name,
		picture.Introduction,
		picture.Category,
		picture.Tags,
		picture.ReviewStatus,
		picture.ReviewMsg,
		picture.ReviewerId,
		picture.Id,
	)
	if err != nil {
		return fmt.Errorf("更新图片信息失败: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("图片不存在或已删除")
	}

	return nil
}

func (d *PictureDao) Delete(ctx context.Context, id int64) error {
	query := `UPDATE picture SET is_delete = 1, update_time = NOW() WHERE id = ?`
	_, err := d.db.ExecContext(ctx, query, id)
	return err
}

func (d *PictureDao) ListBySpaceId(ctx context.Context, spaceId int64) ([]*Picture, error) {
	query := `SELECT id, url, thumbnail_url, name, introduction, category, tags, 
		pic_size, pic_width, pic_height, pic_scale, pic_format, pic_color, 
		user_id, space_id, review_status, review_message, reviewer_id, 
		create_time, edit_time, update_time, is_delete 
		FROM picture WHERE space_id = ? AND is_delete = 0 ORDER BY create_time DESC`

	rows, err := d.db.QueryContext(ctx, query, spaceId)
	if err != nil {
		return nil, fmt.Errorf("查询空间图片列表失败: %v", err)
	}
	defer rows.Close()

	var pictures []*Picture
	for rows.Next() {
		var pic Picture
		err := rows.Scan(
			&pic.Id,
			&pic.Url,
			&pic.ThumbnailUrl,
			&pic.Name,
			&pic.Introduction,
			&pic.Category,
			&pic.Tags,
			&pic.PicSize,
			&pic.PicWidth,
			&pic.PicHeight,
			&pic.PicScale,
			&pic.PicFormat,
			&pic.PicColor,
			&pic.UserId,
			&pic.SpaceId,
			&pic.ReviewStatus,
			&pic.ReviewMsg,
			&pic.ReviewerId,
			&pic.CreateTime,
			&pic.EditTime,
			&pic.UpdateTime,
			&pic.IsDelete,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描图片数据失败: %v", err)
		}
		pictures = append(pictures, &pic)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历图片数据失败: %v", err)
	}

	return pictures, nil
}

func (d *PictureDao) BatchDelete(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	// 构建 IN 查询参数
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`UPDATE picture SET is_delete = 1, update_time = NOW() 
		WHERE id IN (%s) AND is_delete = 0`, strings.Join(placeholders, ","))

	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("批量删除图片失败: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("没有图片被删除")
	}

	return nil
}
