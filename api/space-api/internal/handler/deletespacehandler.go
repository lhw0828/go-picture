package handler

import (
	"net/http"

	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
	"picture/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除团队空间
func deleteSpaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteSpaceLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSpace()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
