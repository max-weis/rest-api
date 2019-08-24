package main

import (
	"gitlab.com/baroprime/prod-rest/app"
)

func main() {
	// get the app config
	config := app.NewConfig()
	// create new app
	app := app.NewApp()
	// init app
	app.Init(config)
	// run app
	app.Run(":80")
}
