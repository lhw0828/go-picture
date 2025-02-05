package admin

import (
	"errors"
	"net/http"
	"strings"

	"picture/api/user-api/internal/logic/admin"
	"picture/api/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除用户
func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从路径中获取id参数
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) < 5 {
			httpx.ErrorCtx(r.Context(), w, errors.New("无效的用户ID"))
			return
		}

		id := parts[len(parts)-1]
		l := admin.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser(id)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
