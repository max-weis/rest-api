# Production REST API

A simple sample REST API which uses `gorilla/mux`, `sirupsen/logrus` and Postgres

- [Project Structure](#project-structure)
- [Main](#main)
- [package app](#package-app)
    - [app](#package-app-app)
    - [controller](#package-app-controller)
    - [db](#package-app-db)
- [package models](#package-models)

## Project Structure <a id="project-structure"></a>
```
├── app                     // contains the main parts of the app
│   ├── app.go              // defines and inits all components of the app
│   ├── controller.go       // contains the REST endpoints
│   └── db.go               // defines functions for the CRUD operations
├── models                  // stores the models for the app
│   └── book.go
└── resources               // sql stuff
    ├── 1-schema.sql
    └── 2-data.sql
├── docker-compose.yml
├── Dockerfile              
├── go.mod
├── go.sum
├── main.go                 // starts the application
└── README.md
```

## The Main <a id="main"></a>

The main is pretty self explanatory.

```go
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
```

First we init the __config__ and __app__.
 The __config__ houses all of the environment variables (specificly for the db connection).
In the next step, the __config__ is used to initialize the application.
And the last step starts the API with the specified port

## Package `app` <a id="package-app"></a>

### `app.go` <a id="package-app-app"></a>

Here we define all the components the application needs. Like the __config__, __logger__, __db__ und __router__.

`newLogger()`, `newDB()`, `newRouter()` and `NewApp()` init there components. With the `Init` func all those functions get summarized into one.

`Run()` fianlly starts the app

### `controller.go` <a id="package-app-controller"></a>

This file contains all the `HandlerFuncs` for the API.
They return a __HandlerFunc__ to make it possible for some configuration or resilience .

### `db.go` <a id="package-app-db"></a>

Here does all the magic happen.
All these function implement the CRUD operations on the database.

## Package `models` <a id="package-models"></a>

In this package all the models are defined. In this example its only one.