package app

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/config"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/api"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/handlers"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/repository"
	"github.com/laa66/trippie-identity-service.git/internal/core/services"
	httpserver "github.com/laa66/trippie-identity-service.git/server"
)

type App struct {
	httpServer *httpserver.HttpServer
}

func (a *App) Run() {
	a.httpServer.Run(8080)
}

func CreateApp() *App {
	// Build components
	logger.InitLogger(slog.LevelDebug)
	config.LoadConfig("../identity_config.yaml")
	app := &App{}
	engine := gin.Default()
	engine.Use(httpserver.ErrorHandler())
	app.httpServer = httpserver.NewHttpServer(engine)

	repositories, err := repository.NewPostgresRepositories();
	if err != nil {
		err.LogStackTrace()
		panic(err)
	}

	identityService := services.NewIdentityService(repositories)
	identityHandler := handlers.NewIdentityHandler(identityService)

	// Register JSON endpoints
	api := api.NewApi(identityHandler)
	api.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))

	return app
}
