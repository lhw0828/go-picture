package handler

import (
	"net/http"

	"picture/api/user-api/internal/logic"
	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登录
func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
