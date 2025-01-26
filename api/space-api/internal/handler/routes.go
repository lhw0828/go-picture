// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"picture/api/space-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 创建空间
				Method:  http.MethodPost,
				Path:    "/space",
				Handler: createSpaceHandler(serverCtx),
			},
			{
				// 获取空间信息
				Method:  http.MethodGet,
				Path:    "/space/:id",
				Handler: getSpaceHandler(serverCtx),
			},
			{
				// 添加空间成员
				Method:  http.MethodPost,
				Path:    "/space/:id/member",
				Handler: addSpaceMemberHandler(serverCtx),
			},
			{
				// 获取空间成员列表
				Method:  http.MethodGet,
				Path:    "/space/:id/members",
				Handler: listSpaceMembersHandler(serverCtx),
			},
			{
				// 更新空间使用容量
				Method:  http.MethodPut,
				Path:    "/space/:id/usage",
				Handler: updateSpaceUsageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
