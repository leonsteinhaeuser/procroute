package procroute

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	ErrAddressNotSet             = errors.New("address not set")
	ErrRouteSetNotPresent        = errors.New("missing routesets")
	ErrNilRouteSetIsNotAllowed   = errors.New("empty route set is not supported")
	ErrNilMiddlewareIsNotAllowed = errors.New("nil middleware is not supported")
)

// RouteMachine represents the manager type used to create and operate endpoints.
type RouteMachine struct {
	server      *http.Server
	routeSets   []*RouteSet
	router      *mux.Router
	middlewares []mux.MiddlewareFunc

	basePath string
	logger   Loggable
}

// NewRouteMachine is a constructor that creates a route machine based on the settings passed as parameters.
// If the port or loggable is not set correctly, you will get errors during execution.
func NewRouteMachine(addr string, port uint16, basePath string, loggable Loggable) *RouteMachine {
	return &RouteMachine{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", addr, port),
		},
		basePath: basePath,
		router:   mux.NewRouter(),
		logger:   loggable,
	}
}

// AddRouteSet provides a method to register a new RouteSet within the route machine
func (rm *RouteMachine) AddRouteSet(routeSet *RouteSet) error {
	if routeSet == nil {
		return ErrNilRouteSetIsNotAllowed
	}

	routeSet.withLogger(rm.logger).withRouterBasePath(rm.basePath).withRouter(rm.router)

	if err := routeSet.build(); err != nil {
		return err
	}

	rm.routeSets = append(rm.routeSets, routeSet)
	return nil
}

// SetReadTimeout provides a method that changes the read timeout within the http server
func (rm *RouteMachine) SetReadTimeout(timeout time.Duration) *RouteMachine {
	rm.server.ReadTimeout = timeout
	return rm
}

// SetReadHeaderTimeout provides a method that changes the read header timeout within the http server
func (rm *RouteMachine) SetReadHeaderTimeout(timeout time.Duration) *RouteMachine {
	rm.server.ReadHeaderTimeout = timeout
	return rm
}

// SetIdleTimeout provides a method that changes the idle timeout within the http server
func (rm *RouteMachine) SetIdleTimeout(timeout time.Duration) *RouteMachine {
	rm.server.IdleTimeout = timeout
	return rm
}

// AddMiddleware injects a middleware just before an endpoint is touched.
func (rm *RouteMachine) AddMiddleware(next Middleware) error {
	if next == nil {
		return ErrNilMiddlewareIsNotAllowed
	}

	if lgg, ok := next.(WithLogger); ok {
		rm.logger.Info("adding logger to middleware")
		lgg.WithLogger(rm.logger)
	}
	rm.middlewares = append(rm.middlewares, next.Middleware)
	return nil
}

// Start provides a method that starts a go routine with the http server
//
// Possible errors:
//  - ErrAddressNotSet
//  - ErrRouteSetNotPresent
//  - socket related errors
func (rm *RouteMachine) Start() error {
	var err error

	// check if starting requirements are set
	if rm.server == nil || rm.server.Addr == "" {
		return ErrAddressNotSet
	}

	// check if routeset requirements are set
	if len(rm.routeSets) < 1 {
		return ErrRouteSetNotPresent
	}

	// initialize middlewares
	rm.router.Use(rm.middlewares...)

	// assign the router to each route set
	for _, routeSet := range rm.routeSets {
		routeSet.router = rm.router
	}

	// assign the router
	rm.server.Handler = rm.router

	afterCh := time.After(500 * time.Millisecond)
	go func() {
		if err = rm.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			rm.logger.Error("server closed unexpectedly with error: %w", err)
			return
		}
	}()
	<-afterCh
	if err != nil {
		return err
	}
	rm.logger.Info("server started on: %s", rm.server.Addr)
	return nil
}

// Stop delegates the stop signal to http.server.Shutdown
func (rm *RouteMachine) Stop() error {
	return rm.server.Shutdown(context.Background())
}
