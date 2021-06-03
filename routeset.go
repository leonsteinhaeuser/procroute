package routemod

import (
	"errors"
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
	parser      Parser
	contentType string
	router      *mux.Router
	basePath    string

	routeSet []interface{}
	logger   Loggable
}

// NewRouteSet defines a new route machine that can be used to create http endpoints
func NewRouteSet(contentType, basePath string, parser Parser) *RouteSet {
	return &RouteSet{
		parser:      parser,
		contentType: contentType,
		basePath:    basePath,
	}
}

// withRouter provides a method that initializes the router within the routeset. This method must be called, before you register any route type
func (rm *RouteSet) withRouter(rt *mux.Router) *RouteSet {
	rm.router = rt
	return rm
}

// withLogger provides a method that sets the logger to the route set
func (rm *RouteSet) withLogger(logger Loggable) *RouteSet {
	rm.logger = logger
	return rm
}

// withRouterBasePath provides a method that prefixes the current base path with the passed in pase path
func (rm *RouteSet) withRouterBasePath(basePath string) *RouteSet {
	rm.basePath = path.Join(basePath, rm.basePath)
	return rm
}

// AddRoutes provides a method that adds routes to the route set
func (rm *RouteSet) AddRoutes(routes ...interface{}) *RouteSet {
	rm.routeSet = append(rm.routeSet, routes...)
	return rm
}

// buildPath is used to combine the basePath plus the uriPath
//
// Example:
//  basePath: "/foo".
//  uriPath: "/bar".
//
// Result = /foo/bar
func (rm *RouteSet) buildPath(uriPath ...string) string {
	return path.Join(rm.basePath, path.Join(uriPath...))
}

func (rm *RouteSet) build() error {
	rm.logger.Debug("compiling routes")
	for _, routeSet := range rm.routeSet {
		// check if the routeset implements the GetRoute interface and if so, register such route
		if rts, ok := routeSet.(GetRoute); ok {
			if err := rm.registerGetRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the GetAllRoute interface and if so, register such route
		if rts, ok := routeSet.(GetAllRoute); ok {
			if err := rm.registerGetAllRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the PostRoute interface and if so, register such route
		if rts, ok := routeSet.(PostRoute); ok {
			if err := rm.registerPostRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the UpdateRoute interface and if so, register such route
		if rts, ok := routeSet.(UpdateRoute); ok {
			if err := rm.registerUpdateRoute(rts); err != nil {
				return err
			}
		}

		// check if the routeset implements the DeleteRoute interface and if so, register such route
		if rts, ok := routeSet.(DeleteRoute); ok {
			if err := rm.registerDeleteRoute(rts); err != nil {
				return err
			}
		}
	}
	return nil
}

// registerPostRoute creates a new post route
func (rm *RouteSet) registerPostRoute(rt PostRoute) error {
	if rt == nil {
		return ErrPostRouteIsNil
	}

	path := rm.buildPath()
	if subPath, ok := rt.(PostRouteRoutePath); ok {
		path = rm.buildPath(subPath.PostRoutePath())
	}
	rm.logger.Info("registered post route at: %s", path)

	rm.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rm.definePostRoute(w, r, rt)
	}).Methods("POST", "OPTIONS")

	return nil
}

// definePostRoute defines the structure used for post routes
func (rm *RouteSet) definePostRoute(w http.ResponseWriter, r *http.Request, rt PostRoute) {
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
		e.write(rm.contentType, rm.parser, w)
		return
	}

	myErr := rm.unmarshal(bts, rt)
	if myErr != nil {
		//log.Panicf("%+v", myErr)
		myErr.write(rm.contentType, rm.parser, w)
		return
	}

	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	httpErr := rt.Post()
	if httpErr != nil {
		httpErr.write(rm.contentType, rm.parser, w)
		return
	}

	w.Header().Add("Content-Type", rm.contentType)
	w.WriteHeader(http.StatusCreated)
}

// registerGetRoute creates a new get route
func (rm *RouteSet) registerGetRoute(rt GetRoute) error {
	if rt == nil {
		return ErrGetRouteIsNil
	}

	path := rm.buildPath()
	if subPath, ok := rt.(GetRoutePath); ok {
		path = rm.buildPath(subPath.GetRoutePath())
	}
	rm.logger.Info("registered get route at: %s", path)

	rm.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rm.defineGetRoute(w, r, rt)
	}).Methods("GET")

	return nil
}

// defineGetRoute defines the structure used for get routes
func (rm *RouteSet) defineGetRoute(w http.ResponseWriter, r *http.Request, rt GetRoute) {
	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	data, httpErr := rt.Get()
	if httpErr != nil {
		httpErr.write(rm.contentType, rm.parser, w)
		return
	}

	bts, err := rm.marshal(data)
	if err != nil {
		err.write(rm.contentType, rm.parser, w)
		return
	}

	w.Header().Add("Content-Type", rm.contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(bts)
}

// registerGetAllRoute creates a new get all route
func (rm *RouteSet) registerGetAllRoute(rt GetAllRoute) error {
	if rt == nil {
		return ErrGetAllRouteIsNil
	}

	path := rm.buildPath()
	if subPath, ok := rt.(GetAllRoutePath); ok {
		path = rm.buildPath(subPath.GetAllRoutePath())
	}
	rm.logger.Info("registered get all route at: %s", path)

	rm.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rm.defineGetAllRoute(w, r, rt)
	}).Methods("GET")

	return nil
}

// defineGetAllRoute defines the structure used for get all routes
func (rm *RouteSet) defineGetAllRoute(w http.ResponseWriter, r *http.Request, rt GetAllRoute) {
	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	data, httpErr := rt.GetAll()
	if httpErr != nil {
		httpErr.write(rm.contentType, rm.parser, w)
		return
	}

	bts, err := rm.parser.Marshal(data)
	if err != nil {
		e := HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
		e.write(rm.contentType, rm.parser, w)
		return
	}

	w.Header().Add("Content-Type", rm.contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(bts)
}

// registerUpdateRoute creates a new update route
func (rm *RouteSet) registerUpdateRoute(rt UpdateRoute) error {
	if rt == nil {
		return ErrUpdateRouteIsNil
	}

	path := rm.buildPath()
	if subPath, ok := rt.(UpdateRouteRoutePath); ok {
		path = rm.buildPath(subPath.UpdateRoutePath())
	}
	rm.logger.Info("registered update route at: %s", path)

	rm.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rm.defineUpdateRoute(w, r, rt)
	}).Methods("UPDATE", "OPTIONS")

	return nil
}

// defineUpdateRoute defines the structure used for update routes
func (rm *RouteSet) defineUpdateRoute(w http.ResponseWriter, r *http.Request, rt UpdateRoute) {
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
		e.write(rm.contentType, rm.parser, w)
		return
	}

	myErr := rm.unmarshal(bts, rt)
	if myErr != nil {
		myErr.write(rm.contentType, rm.parser, w)
		return
	}

	if m, ok := rt.(UrlParams); ok {
		m.SetUrlParams(mux.Vars(r))
	}

	httpErr := rt.Update()
	if httpErr != nil {
		httpErr.write(rm.contentType, rm.parser, w)
		return
	}

	w.Header().Add("Content-Type", rm.contentType)
	w.WriteHeader(http.StatusOK)
}

// registerDeleteRoute creates a new delete route
func (rm *RouteSet) registerDeleteRoute(rt DeleteRoute) error {
	if rt == nil {
		return ErrDeleteRouteIsNil
	}

	path := rm.buildPath()
	if subPath, ok := rt.(DeleteRouteRoutePath); ok {
		path = rm.buildPath(subPath.DeleteRoutePath())
	}
	rm.logger.Info("registered update route at: %s", path)

	rm.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		rm.defineDeleteRoute(w, r, rt)
	}).Methods("DELETE", "OPTIONS")

	return nil
}

// defineUpdateRoute defines the structure used for update routes
func (rm *RouteSet) defineDeleteRoute(w http.ResponseWriter, r *http.Request, rt DeleteRoute) {
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
		e.write(rm.contentType, rm.parser, w)
		return
	}

	myErr := rm.unmarshal(bts, rt)
	if myErr != nil {
		myErr.write(rm.contentType, rm.parser, w)
		return
	}

	httpErr := rt.Delete()
	if httpErr != nil {
		httpErr.write(rm.contentType, rm.parser, w)
		return
	}

	w.Header().Add("Content-Type", rm.contentType)
	w.WriteHeader(http.StatusOK)
}

// unmarshal unmarshals the byte slice into the provided Typer interface and writes an error back to the client, if the marshalling failed
func (rm *RouteSet) unmarshal(bts []byte, typer Typer) *HttpError {
	if err := rm.parser.Unmarshal(bts, typer.Type()); err != nil {
		return &HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
	}
	return nil
}

// marshal marshals the interface into a byte slice
func (rm *RouteSet) marshal(data interface{}) ([]byte, *HttpError) {
	bts, err := rm.parser.Marshal(data)
	if err != nil {
		return nil, &HttpError{
			Status:    http.StatusInternalServerError,
			ErrorCode: "",
			Message:   err.Error(),
		}
	}
	return bts, nil
}
