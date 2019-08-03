package main

import (
	"gitlab.com/baroprime/prod-rest/internal"
	. "gitlab.com/baroprime/prod-rest/app"
)

func main()  {
	// init the environment configuration
	conf := internal.NewConfiguration()
	// init the resolver
	resolv := internal.NewResolver(conf)
	// init the router
	router := internal.NewRouter()

	// init the repo with the help of the resolver
	repo := Repo{
	   DB: resolv.ResolveDB("postgres"),
	}

	// init the app
	a := App{
	   Logger: resolv.ResolveLogger(),
	   Repo: repo,
	   Router: *router,
	}
	a.Run()
}