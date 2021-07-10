package procroute

import (
	"net/http"
	"net/url"
)

// GetRoute provides the interface that must be implemented to create a Get endpoint.
type GetRoute interface {
	// Get represents the method that contains the business logic for receiving a resource.
	//
	// Example
	//  type Model struct {
	//  	Name string `json:"name,omitempty"`
	//  	URL  string `json:"url,omitempty"`
	//  }
	//
	//  type MyType struct {
	//  	Model
	//  }
	//
	//  func (m *MyType) Get(requestData interface{}) (interface{}, *HttpError) {
	//      // do something
	//  	return Model{
	//			Name: "example",
	//			URL: "example.url",
	//      }, nil
	//  }
	Get(requestData interface{}) (interface{}, *HttpError)
}

// GetRoutePath defines an optional child interface that is used to customize route path.
type GetRoutePath interface {
	// GetRoutePath represents an optional method that can be set to define a custom path for the get route.
	//
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) GetRoutePath() string {
	//  	return "/get-all"
	//  }
	GetRoutePath() string
}

// GetAllRoute provides the interface that must be implemented to create a Get All endpoint.
type GetAllRoute interface {
	// GetAll represents the method that contains the business logic for receiving all resources.
	//
	// Example
	//  type Model struct {
	//  	Name string `json:"name,omitempty"`
	//  	URL  string `json:"url,omitempty"`
	//  }
	//
	//  type MyType struct {
	//  	Model
	//  }
	//
	//  func (m *MyType) GetAll(requestData interface{}) ([]interface{}, *HttpError) {
	//      // do something
	//  	return []interface{}{
	//          Model{
	//				Name: "example",
	//				URL: "example.url",
	//          },
	//      }, nil
	//  }
	GetAll(requestData interface{}) ([]interface{}, *HttpError)
}

// GetAllRoutePath defines an optional child interface that is used to customize route path.
type GetAllRoutePath interface {
	// GetAllRoutePath represents an optional method that can be set to define a custom path for the get all route.
	//
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) GetAllRoutePath() string {
	//  	return "/get-all"
	//  }
	GetAllRoutePath() string
}

// PostRoute provides the interface that must be implemented to create a Post endpoint.
type PostRoute interface {
	// Post represents the method that contains the business logic for creating a resource.
	//
	// Example
	//  type Model struct {
	//  	Name string `json:"name,omitempty"`
	//  	URL  string `json:"url,omitempty"`
	//  }
	//
	//  type MyType struct {
	//  	Model
	//  }
	//
	//  func (m *MyType) Post(requestData interface{}) *HttpError {
	//      // do something
	//      fmt.Printf("%+v\n", requestData)
	//  	return nil
	//  }
	Post(requestData interface{}) *HttpError
}

// PostRouteRoutePath defines an optional child interface that is used to customize route path.
type PostRouteRoutePath interface {
	// PostRoutePath represents an optional method that can be set to define a custom path for the post route.
	//
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) PostRoutePath() string {
	//  	return "/post"
	//  }
	PostRoutePath() string
}

// UpdateRoute provides the interface that must be implemented to create an Update endpoint.
type UpdateRoute interface {
	// Update represents the method that contains the business logic for updating a resource.
	//
	// Example
	//  type Model struct {
	//  	Name string `json:"name,omitempty"`
	//  	URL  string `json:"url,omitempty"`
	//  }
	//
	//  type MyType struct {
	//  	Model
	//  }
	//
	//  func (m *MyType) Update(requestData interface{}) *HttpError {
	//      // do something
	//      fmt.Printf("%+v\n", m.Model)
	//  	return nil
	//  }
	Update(requestData interface{}) *HttpError
}

// UpdateRouteRoutePath defines an optional child interface that is used to customize route path.
type UpdateRouteRoutePath interface {
	// UpdateRoutePath represents an optional method that can be set to define a custom path for the update route.
	//
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) UpdateRoutePath() string {
	//  	return "/update"
	//  }
	UpdateRoutePath() string
}

// DeleteRoute provides the interface that must be implemented to create a Delete endpoint.
type DeleteRoute interface {
	// Delete represents the method that contains the business logic for deleting a resource.
	//
	// Example
	//  type Model struct {
	//  	Name string `json:"name,omitempty"`
	//  	URL  string `json:"url,omitempty"`
	//  }
	//
	//  type MyType struct {
	//  	Model
	//  }
	//
	//  func (m *MyType) Delete(requestData interface{}) *HttpError {
	//      // do something
	//      fmt.Printf("%+v\n", m.Model)
	//  	return nil
	//  }
	Delete(requestData interface{}) *HttpError
}

// DeleteRouteRoutePath defines an optional child interface that is used to customize route path.
type DeleteRouteRoutePath interface {
	// DeleteRoutePath represents an optional method that can be set to define a custom path for the delete route.
	//
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) DeleteRoutePath() string {
	//  	return "/delete"
	//  }
	DeleteRoutePath() string
}

// RawRoute provides the interface that must be implemented to create a Raw endpoint.
type RawRoute interface {
	// Raw represents the method that does nothing for you.
	// Any logic must be handled by the user itself.
	// This type can be used when a functionality is missing by this framework.
	Raw(w http.ResponseWriter, r *http.Request)
	// HttpMethods must returns a slice of http methods procroute should use to register the route.
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) HttpMethods() string {
	//  	return []string{"GET","OPTIONS"}
	//  }
	HttpMethods() []string
}

// RawRouteRoutePath defines an optional child interface that is used to customize route path.
type RawRouteRoutePath interface {
	// RawRoutePath represents an optional method that can be set to define a custom path for the delete route.
	//
	// Example:
	//  type MyType struct {}
	//
	//  func (m *MyType) RawRoutePath() string {
	//  	return "/raw"
	//  }
	RawRoutePath() string
}

// UrlParams represents an interface that must be implemented if the route works with url parameters.
type UrlParams interface {
	// SetUrlParam represents a method to pass url params that can be used later to identify resources.
	//
	// Example:
	//  type MyType struct {
	//  	urlParams map[string]string
	//  }
	//
	//  func (m *MyType) SetUrlParams(args map[string]string) {
	//  	m.urlParams = args
	//  }
	SetUrlParams(args map[string]string)
}

type QueryParams interface {
	// SetQueryParams represents a method to pass query params.
	//
	// Example:
	//  type MyType struct {
	//  	queryParams url.Values
	//  }
	//
	//  func (m *MyType) SetQueryParams(args url.Values) {
	//  	m.queryParams = args
	//  }
	SetQueryParams(args url.Values)
}
