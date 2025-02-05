package handler

import (
	"net/http"

	"picture/api/user-api/internal/logic"
	"picture/api/user-api/internal/svc"
	"picture/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取当前用户信息
func getCurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCurrentUserLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
