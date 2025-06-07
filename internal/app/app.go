package app

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/api"
	"github.com/laa66/trippie-identity-service.git/config"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/repository"
	"github.com/laa66/trippie-identity-service.git/server"
)

type App struct {
	httpServer *httpserver.HttpServer
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
	engine.Use(httpserver.ErrorHandler())
	app.httpServer = httpserver.NewHttpServer(engine)

	if _, err := repository.NewPostgresRepositories(); err != nil {
		err.LogStackTrace()
		panic(err)
	}

	// Register JSON endpoints
	api.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))
	return app
}
