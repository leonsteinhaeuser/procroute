package main

import (
	"github.com/leonsteinhaeuser/go-routemod"
)

func main() {
	logger := &ExampleLogger{}

	rm := routemod.NewRouteMachine("0.0.0.0", 8080, "/api", logger)
	rm.AddRouteSet(routemod.NewRouteSet("/example", &JsonParser{}).AddRoutes(&Example{}))
	rm.AddMiddleware(&MyExampleMiddleware{})

	if err := rm.Start(); err != nil {
		panic(err)
	}

	select {}
}
