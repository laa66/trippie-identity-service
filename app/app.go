package app

import (
	"fmt"
	"time"

	"github.com/laa66/trippie-identity-service.git/api"
	"github.com/laa66/trippie-identity-service.git/ctx"
)

type App struct {
	Api *api.Api
}

func (a *App) Run() {
	a.Api.Run()
}

type Object struct {
	Name string
	Date time.Time
}

func Create() *App {
	app := &App{}
	app.Api = api.NewApi()

	endpoint := api.GenericEndpoint[Object]{
		MethodStr: "POST",
		PathStr:   "/test",
		HandlerFn: func(c ctx.AppContext, body Object) (int, any, error) {
			fmt.Println("Received body:", body)
			return 200, map[string]string{"message": "OK"}, nil
		},
	}

	app.Api.RegisterEndpoints([]api.Endpoint{endpoint})
	return app
}
