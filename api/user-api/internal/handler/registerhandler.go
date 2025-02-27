package handler

import (
	"net/http"

	"picture/api/user-api/internal/logic"
	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
	"picture/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户注册
func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
