package app

import "github.com/laa66/trippie-identity-service.git/api"

type App struct {

}

func (a *App) Run() {
	router := api.SetupRouter()
	router.Run(":8080")
}
