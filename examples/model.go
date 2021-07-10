package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/leonsteinhaeuser/procroute"
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
	urlParams   map[string]string
	queryParams url.Values
	logger      procroute.Loggable
}

// Get implements the GetRoute interface
func (e *Example) Get(requestData interface{}) (interface{}, *procroute.HttpError) {
	e.logger.Info("received get request with data: %+#v", requestData)

	if requestData != nil {
		rd, ok := requestData.(MyType)
		if !ok {
			return MyType{
				Name: "Hello",
				URL:  "example.local",
				Timings: Timings{
					ID:        1,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
					DeletedAt: time.Now().UTC(),
				},
			}, nil
		}

		return rd, nil
	}

	return requestData, nil
}

// GetRoutePath implements the GetRoutePath interface
func (e *Example) GetRoutePath() string {
	return "/{id}"
}

// GetAll implements the GetAllRoute interface
func (e *Example) GetAll(requestData interface{}) ([]interface{}, *procroute.HttpError) {
	e.logger.Info("received get all request with data: %+#v", requestData)

	return []interface{}{MyType{
		Name: "Hello",
		URL:  "example.local",
		Timings: Timings{
			ID:        1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			DeletedAt: time.Now().UTC(),
		},
	},
		MyType{
			Name: "2-Hello",
			URL:  "2.example.local",
			Timings: Timings{
				ID:        2,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				DeletedAt: time.Now().UTC(),
			},
		}}, nil
}

// GetAllRoutePath implements the GetAllRoutePath interface
func (e *Example) GetAllRoutePath() string {
	return ""
}

// Post implements the PostRoute interface
func (e *Example) Post(requestData interface{}) *procroute.HttpError {
	e.logger.Info("received post request with data: %+#v", requestData)
	return nil
}

// PostRoutePath implements the PostRouteRoutePath interface
func (e *Example) PostRoutePath() string {
	return ""
}

// Update implements the UpdateRoute interface
func (e *Example) Update(requestData interface{}) *procroute.HttpError {
	e.logger.Info("received update request with data: %+#v", requestData)
	return nil
}

// UpdateRoutePath implements the UpdateRouteRoutePath interface
func (e *Example) UpdateRoutePath() string {
	return "/{id}"
}

// Delete implements the DeleteRoute interface
func (e *Example) Delete(requestData interface{}) *procroute.HttpError {
	e.logger.Info("received delete request with data: %+#v", requestData)
	return nil
}

// DeleteRoutePath implements the DeleteRouteRoutePath interface
func (e *Example) DeleteRoutePath() string {
	return "/{id}"
}

// Raw implements the RawRoute interface
func (e *Example) Raw(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug("URL values: %+v", r.URL.Query())
}

// HttpMethods implements the RawRoute interface
func (e *Example) HttpMethods() []string {
	return []string{"GET", "OPTIONS"}
}

// RawRoutePath implements the RawRouteRoutePath interface
func (e *Example) RawRoutePath() string {
	return "/{id}/raw"
}

// SetUrlParams implements the UrlParams interface
func (e *Example) SetUrlParams(args map[string]string) {
	e.urlParams = args
}

// SetQueryParams implements the QueryParams interface
func (e *Example) SetQueryParams(args url.Values) {
	e.queryParams = args
}

// WithLogger implements the WithLogger interface
func (e *Example) WithLogger(lggbl procroute.Loggable) {
	e.logger = lggbl
}
