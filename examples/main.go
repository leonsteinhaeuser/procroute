package main

import (
	"github.com/leonsteinhaeuser/procroute"
)

func main() {
	rm := procroute.NewRouteMachine("0.0.0.0", 8080, "/api", &ExampleLogger{})
	rm.AddRouteSet(procroute.NewRouteSet("/example", &JsonParser{}).AddRoutes(&Example{}))
	rm.AddMiddleware(&MyExampleMiddleware{})

	if err := rm.Start(); err != nil {
		panic(err)
	}

	select {}
}
