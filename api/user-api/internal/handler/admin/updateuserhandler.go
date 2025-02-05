package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/api/user-api/internal/logic/admin"
	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
)

// 更新用户
func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewUpdateUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
