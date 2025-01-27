package logic

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"picture/common/errorx"
	"picture/rpc/picture-rpc/internal/dao"
	"picture/rpc/picture-rpc/internal/svc"
	"picture/rpc/picture-rpc/picture"
	"strings"
	"time"

	"picture/rpc/space-rpc/space" // 添加 space 包导入

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPictureLogic {
	return &UploadPictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPictureLogic) UploadPicture(in *picture.UploadPictureRequest) (*picture.UploadPictureResponse, error) {
	l.Logger.Infof("开始上传图片: name=%s, size=%d, spaceId=%d, userId=%d",
		in.PicName, len(in.File), in.SpaceId, in.UserId)

	// 校验并处理文件
	if err := l.validateFile(in); err != nil {
		return nil, err
	}

	// 校验空间
	if err := l.validateSpace(in); err != nil {
		return nil, err
	}

	// 开启事务
	tx, err := l.svcCtx.DB.BeginTx(l.ctx, nil)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.DBError, "开启事务失败")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	// 上传并保存图片
	pic, err := l.uploadAndSave(in)
	if err != nil {
		return nil, err
	}

	// 更新空间使用量
	if err := l.updateSpaceUsage(in.SpaceId, pic.PicSize); err != nil {
		return nil, err
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		l.Logger.Errorf("提交事务失败: %+v", err)
		return nil, errorx.NewCodeError(errorx.DBError, "提交事务失败")
	}

	return l.buildResponse(pic), nil
}

func (l *UploadPictureLogic) validateFile(in *picture.UploadPictureRequest) error {
	if len(in.File) == 0 {
		return errorx.NewCodeError(errorx.ParamError, "文件不能为空")
	}
	if len(in.File) > int(l.svcCtx.Config.Upload.MaxSize) {
		return errorx.NewCodeError(errorx.ParamError,
			fmt.Sprintf("文件大小不能超过 %dMB", l.svcCtx.Config.Upload.MaxSize/1024/1024))
	}

	fileType := http.DetectContentType(in.File)
	allowTypes := make(map[string]bool)
	for _, t := range l.svcCtx.Config.Upload.AllowTypes {
		allowTypes[t] = true
	}
	if !allowTypes[fileType] {
		return errorx.NewCodeError(errorx.ParamError,
			fmt.Sprintf("不支持的文件类型: %s，支持的类型: %s",
				fileType, strings.Join(l.svcCtx.Config.Upload.AllowTypes, ", ")))
	}

	if in.PicName == "" {
		in.PicName = fmt.Sprintf("%s_%s.%s",
			time.Now().Format("20060102150405"),
			uuid.New().String()[:8],
			getExtFromMime(fileType),
		)
	}
	return nil
}

func (l *UploadPictureLogic) validateSpace(in *picture.UploadPictureRequest) error {
	if in.SpaceId == 0 {
		return nil
	}

	space, err := l.svcCtx.SpaceRpc.GetSpace(l.ctx, &space.GetSpaceRequest{Id: in.SpaceId})
	if err != nil {
		l.Logger.Errorf("获取空间信息失败: %+v", err)
		return errorx.NewCodeError(errorx.InternalError, "获取空间信息失败")
	}
	if space == nil {
		return errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}
	if space.TotalCount >= space.MaxCount {
		return errorx.NewCodeError(errorx.QuotaError, "空间条数不足")
	}
	if space.TotalSize >= space.MaxSize {
		return errorx.NewCodeError(errorx.QuotaError, "空间大小不足")
	}
	return nil
}

func (l *UploadPictureLogic) uploadAndSave(in *picture.UploadPictureRequest) (*dao.Picture, error) {
	picInfo, err := l.svcCtx.FileManager.UploadPicture(l.ctx, bytes.NewReader(in.File), in.PicName, in.SpaceId, in.UserId)
	if err != nil {
		l.Logger.Errorf("上传图片失败: %+v", err)
		return nil, errorx.NewCodeError(errorx.InternalError, "上传图片失败")
	}

	pic := &dao.Picture{
		Url:          picInfo.Url,
		ThumbnailUrl: picInfo.ThumbnailUrl,
		Name:         in.PicName,
		SpaceId:      in.SpaceId,
		UserId:       in.UserId,
		PicSize:      picInfo.Size,
		PicWidth:     picInfo.Width,
		PicHeight:    picInfo.Height,
		PicScale:     picInfo.Scale,
		PicFormat:    picInfo.Format,
		ReviewStatus: 0,
		CreateTime:   time.Now(),
		EditTime:     time.Now(),
		UpdateTime:   time.Now(),
	}

	err = l.svcCtx.PictureDao.Insert(l.ctx, pic)
	if err != nil {
		l.Logger.Errorf("保存图片信息失败: %+v", err)
		return nil, errorx.NewCodeError(errorx.DBError, "保存图片信息失败")
	}

	return pic, nil
}

func (l *UploadPictureLogic) updateSpaceUsage(spaceId int64, size int64) error {
	if spaceId == 0 {
		return nil
	}

	_, err := l.svcCtx.SpaceRpc.UpdateSpaceUsage(l.ctx, &space.UpdateSpaceUsageRequest{
		SpaceId:   spaceId,
		Size:      size,
		Operation: "add", // 添加 operation 字段，表示增加空间使用量
	})
	if err != nil {
		l.Logger.Errorf("更新空间使用量失败: %+v", err)
		return errorx.NewCodeError(errorx.InternalError, "更新空间使用量失败")
	}
	return nil
}

func (l *UploadPictureLogic) buildResponse(pic *dao.Picture) *picture.UploadPictureResponse {
	return &picture.UploadPictureResponse{
		Picture: &picture.Picture{
			Id:           pic.Id,
			Url:          pic.Url,
			ThumbnailUrl: pic.ThumbnailUrl,
			Name:         pic.Name,
			SpaceId:      pic.SpaceId,
			UserId:       pic.UserId,
			PicSize:      pic.PicSize,
			PicWidth:     pic.PicWidth,
			PicHeight:    pic.PicHeight,
			PicScale:     pic.PicScale,
			PicFormat:    pic.PicFormat,
			ReviewStatus: pic.ReviewStatus,
			CreateTime:   pic.CreateTime.Unix(),
			UpdateTime:   pic.UpdateTime.Unix(),
		},
	}
}

func getExtFromMime(mimeType string) string {
	switch mimeType {
	case "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	case "image/gif":
		return "gif"
	case "image/webp":
		return "webp"
	default:
		return "jpg"
	}
}
