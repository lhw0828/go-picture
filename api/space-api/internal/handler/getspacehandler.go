package handler

import (
	"net/http"
	"strconv"

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
		idStr := r.URL.Path[len("/api/v1/space/"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError("无效的空间ID"))
			return
		}
		req.Id = id

		l := logic.NewGetSpaceLogic(r.Context(), svcCtx)
		resp, err := l.GetSpace(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
