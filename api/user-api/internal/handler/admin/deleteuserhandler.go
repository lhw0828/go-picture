package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/api/user-api/internal/logic/admin"
	"picture/api/user-api/internal/svc"
)

// 删除用户
func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
