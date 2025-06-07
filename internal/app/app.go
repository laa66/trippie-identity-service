package app

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/api"
	"github.com/laa66/trippie-identity-service.git/config"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/repository"
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
	config.LoadConfig("../identity_config.yaml")
	app := &App{}
	engine := gin.Default()
	engine.Use(http_server.ErrorHandler())
	app.httpServer = http_server.NewHttpServer(engine)

	repository.NewPostgresRepositories()

	// Register JSON endpoints
	api.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))
	return app
}
