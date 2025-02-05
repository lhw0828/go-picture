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

// 更新空间使用容量
func updateSpaceUsageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSpaceUsageReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.WithContext(r.Context()).Error("解析请求参数失败", logx.Field("error", err))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateSpaceUsageLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSpaceUsage(&req)
		if err != nil {
			logx.WithContext(r.Context()).Error("更新空间使用量失败", logx.Field("error", err))
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
