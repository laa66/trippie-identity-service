package app

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
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
	engine := gin.Default()
	engine.Use(http_server.ErrorHandler())
	app.httpServer = http_server.NewHttpServer(engine)

	InitLogger()
	// Register JSON endpoints
	api.RegisterIdentityEndpoints(app.httpServer.GetRouterGroup("identity"))

	return app
}

// TODO: move to logger package
func InitLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
