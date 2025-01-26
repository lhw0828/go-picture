package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
)

// 获取空间成员列表
func listSpaceMembersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListSpaceMembersLogic(r.Context(), svcCtx)
		resp, err := l.ListSpaceMembers()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
