package main

import (
	"booklib/app"
	"booklib/config"
)

func main() {
	config := config.GetConfig()
	r := app.NewRouter()

	// r.GET("/covid/summary", app.NewHandler(covidService.HandleRequest))
	app.Run(r, config.Server.Port)

}
