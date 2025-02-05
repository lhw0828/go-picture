package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/api/user-api/internal/logic"
	"picture/api/user-api/internal/svc"
)

// 用户注销
func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
