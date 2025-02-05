package handler

import (
	"net/http"

	"picture/api/space-api/internal/logic"
	"picture/api/space-api/internal/svc"
	"picture/common/errorx"
	"picture/common/types"
	"picture/common/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getSpaceVOHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从路径参数中获取空间ID
		var req types.GetRequest
		err := httpx.Parse(r, &req) // 解析参数
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errorx.NewCodeError(errorx.ParamError, "无效的空间ID"))
			return
		}

		// 获取当前登录用户
		userId, err := utils.GetCurrentUserId(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetSpaceVOLogic(r.Context(), svcCtx)
		resp, err := l.GetSpaceVO(&req, userId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
