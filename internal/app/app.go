package app

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/api"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
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
	logger.InitLogger(slog.LevelDebug)
	app := &App{}
	engine := gin.Default()
	engine.Use(http_server.ErrorHandler())
	app.httpServer = http_server.NewHttpServer(engine)

	// Register JSON endpoints
	api.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))

	return app
}

