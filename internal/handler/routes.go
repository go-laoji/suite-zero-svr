// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"wecom-go-3rd-app-demo/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/callback/data",
				Handler: suiteDataCallbackGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/callback/data",
				Handler: suiteDataCallbackPostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/callback/cmd",
				Handler: suiteCmdCallbackGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/callback/cmd",
				Handler: suiteCmdCallbackPostHandler(serverCtx),
			},
		},
	)
}
