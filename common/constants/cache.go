package constants

import (
	"fmt"
	"time"
)

// 缓存键前缀
const (
	PrefixUser        = "user:"
	PrefixSpace       = "space:"
	PrefixToken       = "token:"
	PrefixPicture     = "picture:"
	PrefixSpaceStats  = "space:stats:"
	PrefixSpaceUsage  = "space:usage:"
	PrefixUserSession = "user:session:"
)

// 缓存过期时间（秒）
const (
	ExpireUser        = 3600 * 24  // 用户信息缓存 24 小时
	ExpireSpace       = 3600 * 24  // 空间信息缓存 24 小时
	ExpireToken       = 3600 * 24  // token 缓存 24 小时
	ExpirePicture     = 3600 * 72  // 图片信息缓存 72 小时
	ExpireSpaceStats  = 3600 * 1   // 空间统计信息缓存 1 小时
	ExpireSpaceUsage  = 300        // 空间使用量缓存 5 分钟
	ExpireUserSession = 3600 * 24  // 用户会话缓存 24 小时
)

// 缓存分类
const (
	CacheTypeUser    = "user"
	CacheTypeSpace   = "space"
	CacheTypePicture = "picture"
)

// 获取缓存键
func GetUserKey(userId int64) string {
	return fmt.Sprintf("%s%d", PrefixUser, userId)
}

func GetSpaceKey(spaceId int64) string {
	return fmt.Sprintf("%s%d", PrefixSpace, spaceId)
}

func GetTokenKey(token string) string {
	return fmt.Sprintf("%s%s", PrefixToken, token)
}

func GetPictureKey(pictureId int64) string {
	return fmt.Sprintf("%s%d", PrefixPicture, pictureId)
}

func GetSpaceStatsKey(spaceId int64) string {
	return fmt.Sprintf("%s%d", PrefixSpaceStats, spaceId)
}

func GetSpaceUsageKey(spaceId int64) string {
	return fmt.Sprintf("%s%d", PrefixSpaceUsage, spaceId)
}

func GetUserSessionKey(userId int64, deviceId string) string {
	return fmt.Sprintf("%s%d:%s", PrefixUserSession, userId, deviceId)
}

// 获取带日期的缓存键
func GetDailyKey(prefix string, id int64) string {
	return fmt.Sprintf("%s%d:%s", prefix, id, time.Now().Format("20060102"))
}

// 获取带时间范围的缓存键
func GetTimeRangeKey(prefix string, id int64, start, end time.Time) string {
	return fmt.Sprintf("%s%d:%s:%s", prefix, id, start.Format("20060102"), end.Format("20060102"))
}

// 获取分页缓存键
func GetPageKey(prefix string, page, pageSize int) string {
	return fmt.Sprintf("%s:page:%d:%d", prefix, page, pageSize)
}