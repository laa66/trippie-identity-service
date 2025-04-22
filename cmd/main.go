package main

import (
	"github.com/laa66/trippie-identity-service.git/app"
)

func main() {
	app := app.Create()
	app.Run()
}
