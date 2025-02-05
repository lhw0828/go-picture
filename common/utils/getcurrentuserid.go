package utils

import (
	"encoding/json"
	"net/http"
)

// GetCurrentUserId 从请求上下文中获取当前用户ID
func GetCurrentUserId(r *http.Request) (int64, error) {
	userIdNumber := r.Context().Value("userId").(json.Number)
	return userIdNumber.Int64()
}