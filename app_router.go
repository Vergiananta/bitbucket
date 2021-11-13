package main

import (
	"bitbucket/connect"
	"bitbucket/controllers"
	"bitbucket/controllers/handler"
	"bitbucket/manager"
	"bitbucket/middlewares"
	"bitbucket/utils/httpParse"
	"bitbucket/utils/httpResponse"
	"github.com/gorilla/mux"
)

type appRouter struct {
	app                  	*majooApp
	parse     				*httpParse.JsonParse
	responder 				httpResponse.IResponder
	connect   				connect.Connect
	logRequestMiddleware 	*middlewares.LogRequestMiddleware
}

type appRoutes struct {
	centerRoutes controllers.IDelivery
	mdw          []mux.MiddlewareFunc
}

func (r *appRouter) InitMainRoutes() {
	r.app.router.Use(r.logRequestMiddleware.Log)
	serviceManager := manager.NewServiceManager(r.connect)
	appRoutes := []appRoutes{
		{
			centerRoutes: handler.NewUserController(r.app.router, r.parse, r.responder, serviceManager.UserUseCase()),
			mdw:          nil,
		},
		{
			centerRoutes: handler.NewProductController(r.app.router, r.parse, r.responder, serviceManager.ProductUseCase()),
			mdw:          nil,
		},

	}

	for _, r := range appRoutes {
		r.centerRoutes.InitRoute(r.mdw...)
	}
}

func NewAppRouter(app *majooApp) *appRouter {
	return &appRouter{
		app,
		httpParse.NewJsonParse(),
		httpResponse.NewJSONResponder(),
		app.connect,
		middlewares.NewLogRequestMiddleware(),
	}
}