package main

import "github.com/laa66/trippie-identity-service.git/internal/app"

func main() {
	app := app.CreateApp()
	app.Run()
}
