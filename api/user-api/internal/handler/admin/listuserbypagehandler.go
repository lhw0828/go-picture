package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/api/user-api/internal/logic/admin"
	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"
)

// 分页获取用户列表
func ListUserByPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewListUserByPageLogic(r.Context(), svcCtx)
		resp, err := l.ListUserByPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
