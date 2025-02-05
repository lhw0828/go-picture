package handler

import (
	"net/http"

	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
	"picture/api/space-api/internal/types"
	"picture/common/response"
	"picture/common/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新空间信息
func updateSpaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSpaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取当前登录用户
		userId, err := utils.GetCurrentUserId(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateSpaceLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSpace(&req, userId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
