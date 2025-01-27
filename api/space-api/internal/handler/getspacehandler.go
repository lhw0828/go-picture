package handler

import (
	"net/http"

	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/constants"
	"picture/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取空间信息
func getSpaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSpaceReq

		// 从路径中获取 ID
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, constants.NewCodeError(constants.ParamError, "无效的空间ID"))
			return
		}

		l := logic.NewGetSpaceLogic(r.Context(), svcCtx)
		resp, err := l.GetSpace(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
