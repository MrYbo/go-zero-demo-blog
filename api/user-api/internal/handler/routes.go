// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/blog/v1/api/user-api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/login",
				Handler: LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user",
				Handler: createHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/:id",
				Handler: findOneHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users",
				Handler: findAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/:id",
				Handler: updateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/:id",
				Handler: destroyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api"),
	)
}
