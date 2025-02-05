package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"picture/api/user-api/internal/logic"
	"picture/api/user-api/internal/svc"
	"picture/api/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户信息
func getUserByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从路径中获取id参数
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) < 5 {
			httpx.ErrorCtx(r.Context(), w, errors.New("无效的用户ID"))
			return
		}

		id, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		req := &types.GetUserByIdReq{
			Id: id,
		}

		l := logic.NewGetUserByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserById(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
