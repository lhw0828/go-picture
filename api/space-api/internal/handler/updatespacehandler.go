package handler

import (
	"net/http"

	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/response"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新团队空间使用容量
func updateSpaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSpaceReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.WithContext(r.Context()).Error("解析请求参数失败", logx.Field("error", err))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateSpaceLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSpace(&req)
		if err != nil {
			logx.WithContext(r.Context()).Error("更新空间失败", logx.Field("error", err))
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
