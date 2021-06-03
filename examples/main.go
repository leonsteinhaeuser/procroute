package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/leonsteinhaeuser/go-routemod"
)

func main() {
	logger := &ExampleLogger{}

	rm := routemod.NewRouteMachine("127.0.0.1", 8080, "/api", logger)
	rm.AddRouteSet(routemod.NewRouteSet("application/json", "/example", &JsonParser{}).AddRoutes(&Example{}))

	if err := rm.Start(); err != nil {
		panic(err)
	}

	select {}
}

// MyType

type Timings struct {
	ID        uint      `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type MyType struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
	Timings
}

// Example data type

type Example struct {
	MyType

	urlParams map[string]string `json:"-"`
}

func (e *Example) Get() (interface{}, *routemod.HttpError) {
	fmt.Println("main.go:", e.urlParams)
	return &Example{
		MyType: MyType{
			Name: "Hello",
			URL:  "example.local",
			Timings: Timings{
				ID:        1,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				DeletedAt: time.Now().UTC(),
			},
		},
	}, nil
}

func (e *Example) GetRoutePath() string {
	return "/{id}"
}

func (e *Example) GetAll() ([]interface{}, *routemod.HttpError) {
	return []interface{}{
		Example{
			MyType: MyType{
				Name: "test1",
				URL:  "test1",
				Timings: Timings{
					ID:        0,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: time.Now(),
				},
			},
		},
		Example{
			MyType: MyType{
				Name: "test2",
				URL:  "test2",
				Timings: Timings{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: time.Now(),
				},
			},
		},
	}, nil
}

func (e *Example) SetUrlParams(args map[string]string) {
	e.urlParams = args
}

func (e *Example) Type() interface{} {
	return e
}

// Example logger

type ExampleLogger struct{}

func (e *ExampleLogger) Trace(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tTRACE\t"+format+"\n", v...)
}

func (e *ExampleLogger) Debug(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tDEBUG\t"+format+"\n", v...)
}

func (e *ExampleLogger) Info(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tINFO\t"+format+"\n", v...)
}

func (e *ExampleLogger) Warn(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tWARN\t"+format+"\n", v...)
}

func (e *ExampleLogger) Error(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tERROR\t"+format+"\n", v...)
}

func (e *ExampleLogger) Fatal(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tFATAL\t"+format+"\n", v...)
}

// Example parser

type JsonParser struct{}

func (j *JsonParser) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j *JsonParser) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
