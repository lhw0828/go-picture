package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/api/user-api/internal/logic/admin"
	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
)

// 创建用户
func AddUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewAddUserLogic(r.Context(), svcCtx)
		resp, err := l.AddUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
