package app

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/laa66/trippie-identity-service.git/config"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/handlers"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/http"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/repository"
	"github.com/laa66/trippie-identity-service.git/internal/core/services"
	"github.com/laa66/trippie-identity-service.git/server"
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

	if _, err := repository.NewPostgresRepositories(); err != nil {
		err.LogStackTrace()
		panic(err)
	}

	identityService := services.NewIdentityService()
	identityHandler := handlers.NewIdentityHandler(identityService)
	httpServer := http.NewHTTPServer(identityHandler)

	// Register JSON endpoints
	httpServer.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))

	return app
}
