# routemod

[![unit-tests](https://github.com/leonsteinhaeuser/go-routemod/actions/workflows/test.yml/badge.svg)](https://github.com/leonsteinhaeuser/go-routemod/actions/workflows/test.yml)

Routemod serves the purpose of simplifying the creation of web-based applications. The main goal was to create a framework that implements all the necessary http functions so that the end user can concentrate on the business logic.

## Getting started

The main two entrypoints of the application are the following two: `routemod.NewRouteMachine` and `routemod.NewRouteSet`.

Great it seems simple, but how do i start?

1. Implement the `routemod.Loggable` interface. You can find an example in the [logger.go](examples/logger.go) file.
2. Implement the `routemod.Parser` interface. You can find an example in the [parser.go](examples/parser.go) file.
3. Implement one or more of the [route interfaces](interfaces.go#L12-L65). You can find an example in the [model.go](examples/model.go) file.
4. Define a routing machine and routes. You can find an example in the [main.go](examples/main.go) file.

When you start the example, you get an output like the following:

```txt
2021-06-06T01:39:22+02:00       DEBUG   compiling routes
2021-06-06T01:39:22+02:00       INFO    registered get route at: /api/example/{id}
2021-06-06T01:39:22+02:00       INFO    registered get all route at: /api/example
2021-06-06T01:39:22+02:00       INFO    registered post route at: /api/example
2021-06-06T01:39:22+02:00       INFO    registered update route at: /api/example
2021-06-06T01:39:22+02:00       INFO    registered delete route at: /api/example/{id}
2021-06-06T01:39:22+02:00       INFO    server started on: 0.0.0.0:8080
```

This indicates that the sample application is running at `0.0.0.0:8080`. You can confirm this by running a curl request against the endpoint `get all`:

```bash
curl -i -X GET http://localhost:8080/api/example
```

The answer should be the same as the following text snipped:

```txt
HTTP/1.1 200 OK
Content-Type: application/json
Date: xxx, xxx Jun 2021 xx:xx:xx GMT
Content-Length: 358

[
    {
        "name": "test1",
        "url": "test1",
        "id": 1,
        "createdAt": "2021-06-06T01:51:11.0862583+02:00",
        "updatedAt": "2021-06-06T01:51:11.086261+02:00",
        "deletedAt": "2021-06-06T01:51:11.0862635+02:00"
    }, 
    {
        "name": "test2",
        "url": "test2",
        "id": 2,
        "createdAt": "2021-06-06T01:51:11.0862661+02:00",
        "updatedAt": "2021-06-06T01:51:11.0862687+02:00",
        "deletedAt": "2021-06-06T01:51:11.0862713+02:00"
    }
]
```
