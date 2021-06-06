package main

import (
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
	logger    routemod.Loggable
}

func (e *Example) Get() (interface{}, *routemod.HttpError) {
	e.logger.Info("get route hit")
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
	e.logger.Info("get all route hit")
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
	e.logger.Info("post route hit: %+v", e.MyType)
	return nil
}

func (e *Example) Update() *routemod.HttpError {
	e.logger.Info("update route hit: %+v", e.MyType)
	return nil
}

func (e *Example) Delete() *routemod.HttpError {
	e.logger.Info("delete route hit: %+v", e.MyType)
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

func (e *Example) WithLogger(lggbl routemod.Loggable) {
	e.logger = lggbl
}
