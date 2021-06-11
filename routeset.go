package procroute

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

var (
	ErrGetRouteIsNil    = errors.New("get route is nil")
	ErrGetAllRouteIsNil = errors.New("get all route is nil")
	ErrPostRouteIsNil   = errors.New("post route is nil")
	ErrUpdateRouteIsNil = errors.New("update route is nil")
	ErrDeleteRouteIsNil = errors.New("delete route is nil")
)

// RouteSet defines a structure that is used to create an endpoint set based on the base path
type RouteSet struct {
	parser   Parser
	router   *mux.Router
	basePath string

	routeSet []interface{}
	logger   Loggable
}

// NewRouteSet defines a new route set that is used to genereate http endpoints
func NewRouteSet(basePath string, parser Parser) *RouteSet {
	return &RouteSet{
		parser:   parser,
		basePath: basePath,
	}
}

// withRouter provides a method that initializes the router within the routeset. This method must be called, before you register any route type
func (rs *RouteSet) withRouter(rt *mux.Router) *RouteSet {
	rs.router = rt
	return rs
}

// withLogger provides a method that sets the logger to the route set
func (rs *RouteSet) withLogger(logger Loggable) *RouteSet {
	rs.logger = logger
	return rs
}

// withRouterBasePath provides a method that prefixes the route set endpoints with the passed in base path
func (rs *RouteSet) withRouterBasePath(basePath string) *RouteSet {
	rs.basePath = path.Join(basePath, rs.basePath)
	return rs
}

// AddRoutes provides a method that adds routes to the route set.
// When calling AddRoutes, ensure that your types do not overwrite each other.
//
// Bad example:
//  type Model struct {
//  	Name string `json:"name,omitempty"`
//  	URL  string `json:"url,omitempty"`
//  }
//
//  type MyType struct {
//  	Model
//  }
//
//  func (m *MyType) Type() interface{} {
//  	return &m.Model
//  }
//
//  func (m *MyType) Get() (interface{}, *HttpError) {
//      // do something
//  	return Model{
//			Name: "example",
//			URL: "example.url",
//      }, nil
//  }
//
//  type MyType2 struct {
//  	Model
//  }
//
//  func (m *MyType2) Type() interface{} {
//  	return &m.Model
//  }
//
//  func (m *MyType2) Get() (interface{}, *HttpError) {
//      // do something
//  	return Model{
//			Name: "example",
//			URL: "example.url",
//      }, nil
//  }
//
//  rs.AddRoutes(&MyType{}, &MyType2{})
func (rs *RouteSet) AddRoutes(routes ...interface{}) *RouteSet {
	rs.routeSet = append(rs.routeSet, routes...)
	return rs
}

// buildPath is used to combine the basePath plus the uriPath
//
// Example:
//  basePath: "/foo".
//  uriPath: "/bar".
//
// Result = /foo/bar
func (rs *RouteSet) buildPath(uriPath ...string) string {
	return path.Join(rs.basePath, path.Join(uriPath...))
}

func (rs *RouteSet) build() error {
	rs.logger.Debug("compiling routes")
	for _, routeSet := range rs.routeSet {
		// check if the routeset implements the WithLogger interface
		if wl, ok := routeSet.(WithLogger); ok {
			wl.WithLogger(rs.logger)
		}

		// check if the routeset implements the GetRoute interface and if so, register such route
		if rts, ok := routeSet.(GetRoute); ok {
			if err := rs.registerGetRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the GetAllRoute interface and if so, register such route
		if rts, ok := routeSet.(GetAllRoute); ok {
			if err := rs.registerGetAllRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the PostRoute interface and if so, register such route
		if rts, ok := routeSet.(PostRoute); ok {
			if err := rs.registerPostRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the UpdateRoute interface and if so, register such route
		if rts, ok := routeSet.(UpdateRoute); ok {
			if err := rs.registerUpdateRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the DeleteRoute interface and if so, register such route
		if rts, ok := routeSet.(DeleteRoute); ok {
			if err := rs.registerDeleteRoute(rts); err != nil {
				return err
			}
		}
	}
	return nil
}

// registerPostRoute creates a new post route
func (rs *RouteSet) registerPostRoute(rt PostRoute) error {
	if rt == nil {
		return ErrPostRouteIsNil
	}

	path := rs.buildPath()
	if subPath, ok := rt.(PostRouteRoutePath); ok {
		path = rs.buildPath(subPath.PostRoutePath())
	}
	rs.logger.Info("registered post route at: %s", path)

	rs.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rs.definePostRoute(w, r, rt)
	}).Methods("POST", "OPTIONS")

	return nil
}

// definePostRoute defines the structure used for post routes
func (rs *RouteSet) definePostRoute(w http.ResponseWriter, r *http.Request, rt PostRoute) {
	err := rs.readBody(r.Body, rt)
	if err != nil {
		err.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	httpErr := rt.Post()
	if httpErr != nil {
		httpErr.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	w.Header().Add("Content-Type", rs.parser.MimeType())
	w.WriteHeader(http.StatusCreated)
}

// registerGetRoute creates a new get route
func (rs *RouteSet) registerGetRoute(rt GetRoute) error {
	if rt == nil {
		return ErrGetRouteIsNil
	}

	path := rs.buildPath()
	if subPath, ok := rt.(GetRoutePath); ok {
		path = rs.buildPath(subPath.GetRoutePath())
	}
	rs.logger.Info("registered get route at: %s", path)

	rs.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rs.defineGetRoute(w, r, rt)
	}).Methods("GET")

	return nil
}

// defineGetRoute defines the structure used for get routes
func (rs *RouteSet) defineGetRoute(w http.ResponseWriter, r *http.Request, rt GetRoute) {
	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	data, httpErr := rt.Get()
	if httpErr != nil {
		httpErr.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	bts, err := rs.marshal(data)
	if err != nil {
		err.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	w.Header().Add("Content-Type", rs.parser.MimeType())
	w.WriteHeader(http.StatusOK)
	w.Write(bts)
}

// registerGetAllRoute creates a new get all route
func (rs *RouteSet) registerGetAllRoute(rt GetAllRoute) error {
	if rt == nil {
		return ErrGetAllRouteIsNil
	}

	path := rs.buildPath()
	if subPath, ok := rt.(GetAllRoutePath); ok {
		path = rs.buildPath(subPath.GetAllRoutePath())
	}
	rs.logger.Info("registered get all route at: %s", path)

	rs.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rs.defineGetAllRoute(w, r, rt)
	}).Methods("GET")

	return nil
}

// defineGetAllRoute defines the structure used for get all routes
func (rs *RouteSet) defineGetAllRoute(w http.ResponseWriter, r *http.Request, rt GetAllRoute) {
	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	data, httpErr := rt.GetAll()
	if httpErr != nil {
		httpErr.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	bts, err := rs.parser.Marshal(data)
	if err != nil {
		e := HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
		e.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	w.Header().Add("Content-Type", rs.parser.MimeType())
	w.WriteHeader(http.StatusOK)
	w.Write(bts)
}

// registerUpdateRoute creates a new update route
func (rs *RouteSet) registerUpdateRoute(rt UpdateRoute) error {
	if rt == nil {
		return ErrUpdateRouteIsNil
	}

	path := rs.buildPath()
	if subPath, ok := rt.(UpdateRouteRoutePath); ok {
		path = rs.buildPath(subPath.UpdateRoutePath())
	}
	rs.logger.Info("registered update route at: %s", path)

	rs.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rs.defineUpdateRoute(w, r, rt)
	}).Methods("PUT", "OPTIONS")

	return nil
}

// defineUpdateRoute defines the structure used for update routes
func (rs *RouteSet) defineUpdateRoute(w http.ResponseWriter, r *http.Request, rt UpdateRoute) {
	err := rs.readBody(r.Body, rt)
	if err != nil {
		err.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	httpErr := rt.Update()
	if httpErr != nil {
		httpErr.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	w.Header().Add("Content-Type", rs.parser.MimeType())
	w.WriteHeader(http.StatusOK)
}

// registerDeleteRoute creates a new delete route
func (rs *RouteSet) registerDeleteRoute(rt DeleteRoute) error {
	if rt == nil {
		return ErrDeleteRouteIsNil
	}

	path := rs.buildPath()
	if subPath, ok := rt.(DeleteRouteRoutePath); ok {
		path = rs.buildPath(subPath.DeleteRoutePath())
	}
	rs.logger.Info("registered delete route at: %s", path)

	rs.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rs.defineDeleteRoute(w, r, rt)
	}).Methods("DELETE", "OPTIONS")

	return nil
}

// defineUpdateRoute defines the structure used for update routes
func (rs *RouteSet) defineDeleteRoute(w http.ResponseWriter, r *http.Request, rt DeleteRoute) {
	err := rs.readBody(r.Body, rt)
	if err != nil {
		err.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	httpErr := rt.Delete()
	if httpErr != nil {
		httpErr.write(rs.parser.MimeType(), rs.parser, w)
		return
	}

	w.Header().Add("Content-Type", rs.parser.MimeType())
	w.WriteHeader(http.StatusOK)
}

// unmarshal unmarshals the byte slice into the provided Typer interface and writes an error back to the client, if the marshalling failed
func (rs *RouteSet) unmarshal(bts []byte, typer Typer) *HttpError {
	if err := rs.parser.Unmarshal(bts, typer.Type()); err != nil {
		return &HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
	}
	return nil
}

// marshal marshals the interface into a byte slice
func (rs *RouteSet) marshal(data interface{}) ([]byte, *HttpError) {
	bts, err := rs.parser.Marshal(data)
	if err != nil {
		return nil, &HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
	}
	return bts, nil
}

func (rs *RouteSet) readBody(r io.Reader, rt Typer) *HttpError {
	bts, err := ioutil.ReadAll(r)
	if err != nil {
		return &HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
	}

	return rs.unmarshal(bts, rt)
}
