package routemod

import (
	"errors"
	"net/http"
)

var (
	ErrHttpResponseWriterNotSet = errors.New("http response write is not set")
)

// GetRoute provides the interface that is used for http get routes
type GetRoute interface {
	Get() (interface{}, *HttpError)
	Typer
}

// GetRoutePath defines an optional child path for the get endpoint
type GetRoutePath interface {
	GetRoutePath() string
}

// GetAllRoute provides the interface that is used for http get all routes
type GetAllRoute interface {
	GetAll() ([]interface{}, *HttpError)
	Typer
}

// GetAllRoutePath defines an optional child path for the get all endpoint
type GetAllRoutePath interface {
	GetAllRoutePath() string
}

// PostRoute provides the interface that is used for http post routes
type PostRoute interface {
	Post() *HttpError
	Typer
}

// PostRouteRoutePath defines an optional child path for the post endpoint
type PostRouteRoutePath interface {
	PostRoutePath() string
}

// UpdateRoute provides the interface that is used for http update routes
type UpdateRoute interface {
	Update() *HttpError
	Typer
}

// UpdateRouteRoutePath defines an optional child path for the update endpoint
type UpdateRouteRoutePath interface {
	UpdateRoutePath() string
}

// DeleteRoute provides the interface that is used for http delete routes
type DeleteRoute interface {
	Delete() *HttpError
	Typer
}

// DeleteRouteRoutePath defines an optional child path for the delete endpoint
type DeleteRouteRoutePath interface {
	DeleteRoutePath() string
}

// Typer provides a method that is used to create a new object based on the returned type. Ensure that the returned type is a pointer.
//
// Example:
//  func (r *Example) Type() interface{} {return &MyModelType{}}
type Typer interface {
	Type() interface{}
}

// HttpError represents the datatype used as error response
type HttpError struct {
	Status    int
	ErrorCode string
	Message   string
}

func (h *HttpError) write(contentType string, parser Parser, w http.ResponseWriter) error {
	if w == nil {
		return ErrHttpResponseWriterNotSet
	}

	bts, err := parser.Marshal(h)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(h.Status)
	w.Write(bts)
	return nil
}

// UrlParams represents an interface that must be implemented, if the route handles url params
type UrlParams interface {
	SetUrlParams(args map[string]string)
}

// Parser provides the interface that must be implemented to marshal and unmarshal the data sent during http request and http responses
type Parser interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}
