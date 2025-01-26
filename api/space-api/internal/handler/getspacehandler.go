package handler

import (
	"net/http"

	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取空间信息
func getSpaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSpaceReq

		// 从路径中获取 ID
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError("无效的空间ID"))
			return
		}

		l := logic.NewGetSpaceLogic(r.Context(), svcCtx)
		resp, err := l.GetSpace(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
