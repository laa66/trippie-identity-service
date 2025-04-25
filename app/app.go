package app

import (
	"github.com/laa66/trippie-identity-service.git/api"
	http_server "github.com/laa66/trippie-identity-service.git/server"
)

type App struct {
	httpServer *http_server.HttpServer
}

func (a *App) Run() {
	a.httpServer.Run()
}

func CreateApp() *App {
	// Build components
	app := &App{}
	app.httpServer = http_server.NewHttpServer()

	// Register JSON endpoints
	api.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))

	return app
}
