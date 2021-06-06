package routemod

import (
	"errors"
	"net/http"
)

var (
	ErrHttpResponseWriterNotSet = errors.New("http response writer is not set")
)

// GetRoute provides the interface that must be implemented to create a Get endpoint.
type GetRoute interface {
	Get() (interface{}, *HttpError)
	Typer
}

// GetRoutePath defines an optional child interface that is used to customize route path.
type GetRoutePath interface {
	GetRoutePath() string
}

// GetAllRoute provides the interface that must be implemented to create a Get All endpoint.
type GetAllRoute interface {
	GetAll() ([]interface{}, *HttpError)
	Typer
}

// GetAllRoutePath defines an optional child interface that is used to customize route path.
type GetAllRoutePath interface {
	GetAllRoutePath() string
}

// PostRoute provides the interface that must be implemented to create a Post endpoint.
type PostRoute interface {
	Post() *HttpError
	Typer
}

// PostRouteRoutePath defines an optional child interface that is used to customize route path.
type PostRouteRoutePath interface {
	PostRoutePath() string
}

// UpdateRoute provides the interface that must be implemented to create an Update endpoint.
type UpdateRoute interface {
	Update() *HttpError
	Typer
}

// UpdateRouteRoutePath defines an optional child interface that is used to customize route path.
type UpdateRouteRoutePath interface {
	UpdateRoutePath() string
}

// DeleteRoute provides the interface that must be implemented to create a Delete endpoint.
type DeleteRoute interface {
	Delete() *HttpError
	Typer
}

// DeleteRouteRoutePath defines an optional child interface that is used to customize route path.
type DeleteRouteRoutePath interface {
	DeleteRoutePath() string
}

// Typer represents an interface used to reference back to an object. Make sure that the returned type is a pointer to the object.
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

// write marshals the error message and sends it back to the client
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
	// SetUrlParam represents a method with which url-params can be passed that can later be used to identify resources.
	SetUrlParams(args map[string]string)
}

// Parser provides the interface that must be implemented to marshal and unmarshal the data sent during http request and http responses
type Parser interface {
	// Unmarshal parses the encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an error.
	Unmarshal(data []byte, v interface{}) error
	// Marshal returns the encoded data as byte slice.
	Marshal(v interface{}) ([]byte, error)
	// MimeType returns the associated mime type in string representation. A list of available MIME types can be found at: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	MimeType() string
}

// Middleware defines an interface that is used to inject a middleware
type Middleware interface {
	Middleware(http.Handler) http.Handler
}

// WithLogger provides an optional interface that is used to attach the logger to the route or middleware
type WithLogger interface {
	WithLogger(Loggable)
}
