package main

import (
	"fmt"
	"time"

	"github.com/leonsteinhaeuser/go-routemod"
)

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

// Define the type that implements the interfaces

type Example struct {
	MyType

	urlParams map[string]string
}

func (e *Example) Get() (interface{}, *routemod.HttpError) {
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
					ID:        1,
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
					ID:        2,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: time.Now(),
				},
			},
		},
	}, nil
}

func (e *Example) Post() *routemod.HttpError {
	fmt.Printf("post: %+v\n", e.MyType)
	return nil
}

func (e *Example) Update() *routemod.HttpError {
	fmt.Printf("put: %+v\n", e.MyType)
	return nil
}

func (e *Example) Delete() *routemod.HttpError {
	fmt.Printf("delete: %+v\n", e.MyType)
	return nil
}

func (e *Example) DeleteRoutePath() string {
	return "/{id}"
}

func (e *Example) SetUrlParams(args map[string]string) {
	e.urlParams = args
}

func (e *Example) Type() interface{} {
	return e
}
