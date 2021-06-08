# procroute

[![unit-tests](https://github.com/leonsteinhaeuser/procroute/actions/workflows/test.yml/badge.svg)](https://github.com/leonsteinhaeuser/procroute/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/leonsteinhaeuser/procroute/branch/main/graph/badge.svg?token=3OEL9ZLQRM)](https://codecov.io/gh/leonsteinhaeuser/procroute)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/leonsteinhaeuser/procroute)
![GitHub issues](https://img.shields.io/github/issues-raw/leonsteinhaeuser/procroute)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/leonsteinhaeuser/procroute)
[![Godoc reference](https://godoc.org/github.com/leonsteinhaeuser/procroute?status.svg)](http://godoc.org/github.com/leonsteinhaeuser/procroute)

Procroute serves the purpose of simplifying the creation of web-based applications. The main goal was to create a framework that implements all the necessary http functions so that the end user can concentrate on the business logic.

## Getting started

The main two entrypoints of the application are the following: [procroute.NewRouteMachine](routemachine.go#L30) and [procroute.NewRouteSet](routeset.go#L32).

Great it seems simple, but how do i start?

1. Implement the [procroute.Loggable](logger.go#L3-L10) interface. You can find an example in the [logger.go](examples/logger.go) file.
2. Implement the [procroute.Parser](interfaces.go#L104-L112) interface. You can find an example in the [parser.go](examples/parser.go) file.
3. Implement one or more of the [route interfaces](interfaces.go#L12-L65). You can find an example in the [model.go](examples/model.go) file.
4. Define a [routing machine](routemachine.go#L30) and [routes](routeset.go#L32). You can find an example in the [main.go](examples/main.go) file.

### Define a logger

The following example implements the *Loggable* interface, that must be implemented to run the routemachine.

```go
type ExampleLogger struct{}

func (e *ExampleLogger) buildLogerEntry(prefix, format string, v ...interface{}) {
    t := time.Now().Format(time.RFC3339)
    fmt.Printf(t+"\t"+prefix+"\t"+format+"\n", v...)
}

func (e *ExampleLogger) Trace(format string, v ...interface{}) {
    e.buildLogerEntry("TRACE", format, v...)
}

func (e *ExampleLogger) Debug(format string, v ...interface{}) {
    e.buildLogerEntry("DEBUG", format, v...)
}

func (e *ExampleLogger) Info(format string, v ...interface{}) {
    e.buildLogerEntry("INFO", format, v...)
}

func (e *ExampleLogger) Warn(format string, v ...interface{}) {
    e.buildLogerEntry("WARN", format, v...)
}

func (e *ExampleLogger) Error(format string, v ...interface{}) {
    e.buildLogerEntry("ERROR", format, v...)
}

func (e *ExampleLogger) Fatal(format string, v ...interface{}) {
    e.buildLogerEntry("FATAL", format, v...)
}
```

### Define a parser

The following example implements the *Parser* interface, that must be implemented to create a routeset.

```go
type JsonParser struct{}

func (j *JsonParser) Unmarshal(data []byte, v interface{}) error {
    return json.Unmarshal(data, v)
}

func (j *JsonParser) Marshal(v interface{}) ([]byte, error) {
    return json.Marshal(v)
}

func (j *JsonParser) MimeType() string {
    return "application/json"
}
```

### Get endpoint

The following example implements the *GetRoute* interface that is used to publish an HTTP GET endpoint.

```go
type MyModel struct {
    ID uint `json:"id"`
    Name string `json:"id"`
}

type Example struct {
    MyModel
}

func (e *Example) Get() (interface{}, *procroute.HttpError) {
    // implement your business logic and return a value
    
    // the returned value is parsed into the defined format and available at the get endpoint
    return &MyModel{
        ID: 12,
        Name: "example",
    }, nil
}

func main() {
    rm := procroute.NewRouteMachine("0.0.0.0", 8080, "/api", &ExampleLogger{})
    rm.AddRouteSet(procroute.NewRouteSet("/example", &JsonParser{}).AddRoutes(&Example{}))

    if err := rm.Start(); err != nil {
        panic(err)
    }

    select {}
}
```

### Define a custom route path, inject the Loggable and UrlParams interface

```go
type MyModel struct {
    ID uint `json:"id"`
    Name string `json:"id"`
}

type Example struct {
    MyModel

    urlParams map[string]string
    loggable procroute.Loggable
}

func (e *Example) Type() interface{} {
    return e
}

func (e *Example) Get() (interface{}, *procroute.HttpError) {
    // implement your business logic and return a value
    e.loggable.Info("id contains the value: %s", e.urlParams["id"])

    // the returned value is parsed into the defined format and available at the get endpoint
    return &MyModel{
        ID: 12,
        Name: "example",
    }, nil
}

func (e *Example) GetRoutePath() string {
    return "/{id}"
}

func (e *Example) SetUrlParams(args map[string]string) {
    e.urlParams = args
}

func (e *Example) WithLogger(loggable procroute.Loggable) {
    e.loggable = loggable
}

func main() {
    rm := procroute.NewRouteMachine("0.0.0.0", 8080, "/api", &ExampleLogger{})
    rm.AddRouteSet(procroute.NewRouteSet("/example", &JsonParser{}).AddRoutes(&Example{}))

    if err := rm.Start(); err != nil {
        panic(err)
    }

    select {}
}

```

### Defining a middleware

The following code snippet provides you a logger middleware that prints the method and endpoint path to the console.

```go
type MyExampleMiddleware struct {
    logger procroute.Loggable
}

func (m *MyExampleMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        m.logger.Debug("method: %s path: %s", r.Method, r.URL)
        next.ServeHTTP(w, r)
    })
}

func (m *MyExampleMiddleware) WithLogger(lggbl procroute.Loggable) {
    m.logger = lggbl
}

func main() {
    rm := procroute.NewRouteMachine("0.0.0.0", 8080, "/api", &ExampleLogger{})
    rm.AddMiddleware(&MyExampleMiddleware{})

    if err := rm.Start(); err != nil {
        panic(err)
    }

    select {}
}

```

### Starting the example

Since this repository contains a fully functional example, clone the repository, navigate to the examples folder, and run:

```go
go run .
```

The output generated by the example should look something like this:

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
