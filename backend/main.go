package main

import (
	"booklib/app"
	"booklib/app/booklib"
	"booklib/app/storage"
	"booklib/config"
	"booklib/database"
)

func main() {
	config := config.GetConfig()
	r := app.NewRouter()

	db := database.NewMySQL(config)

	bookStorage := storage.NewBooklibDb(&db)
	bookService := booklib.NewBooklibService(bookStorage)
	bookHandler := app.NewBookHandler(bookService)

	healthHandler := func(c app.Context) {
		c.JSON("ok")
	}

	r.GET("/health", app.NewHandler(healthHandler))
	r.POST("/add-book", app.NewHandler(bookHandler.HandleAdd))
	r.PUT("/edit-book", app.NewHandler(bookHandler.HandleEdit))
	r.DELETE("/delete-book", app.NewHandler(bookHandler.HandleDelete))
	app.Run(r, config.Server.Port)

}
