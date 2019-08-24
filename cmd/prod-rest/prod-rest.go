package main

import (
	"gitlab.com/baroprime/prod-rest/app"
)

func main() {
	// init the router
	router := app.NewRouter()
	// init the db
	db := app.NewDB()
	//init the logger
	logger := app.NewLogger()

	// init the app
	a := App{
		Logger: logger,
		DB:     db,
		Router: *router,
	}
	a.Run(":8080")
}
