package routemod

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

type RouteMachine struct {
	server      *http.Server
	routeSets   []*RouteSet
	router      *mux.Router
	middlewares []mux.MiddlewareFunc

	basePath string
	logger   Loggable
}

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
func (r *RouteMachine) AddRouteSet(routeSet *RouteSet) error {
	if routeSet == nil {
		return ErrNilRouteSetIsNotAllowed
	}

	routeSet.withLogger(r.logger).withRouterBasePath(r.basePath).withRouter(r.router)

	if err := routeSet.build(); err != nil {
		return err
	}

	r.routeSets = append(r.routeSets, routeSet)
	return nil
}

// SetReadTimeout provides a method that changes the read timeout within the http server
func (r *RouteMachine) SetReadTimeout(timeout time.Duration) *RouteMachine {
	r.server.ReadTimeout = timeout
	return r
}

// SetReadHeaderTimeout provides a method that changes the read header timeout within the http server
func (r *RouteMachine) SetReadHeaderTimeout(timeout time.Duration) *RouteMachine {
	r.server.ReadHeaderTimeout = timeout
	return r
}

// SetIdleTimeout provides a method that changes the idle timeout within the http server
func (r *RouteMachine) SetIdleTimeout(timeout time.Duration) *RouteMachine {
	r.server.IdleTimeout = timeout
	return r
}

// AddMiddleware injects a middleware just before an endpoint is touched.
func (r *RouteMachine) AddMiddleware(next Middleware) error {
	if next == nil {
		return ErrNilMiddlewareIsNotAllowed
	}

	if lgg, ok := next.(WithLogger); ok {
		r.logger.Info("adding logger to middleware")
		lgg.WithLogger(r.logger)
	}
	r.middlewares = append(r.middlewares, next.Middleware)
	return nil
}

// Start provides a method that starts a go routine to provides the http server
func (r *RouteMachine) Start() error {
	var err error

	// check if starting requirements are set
	if r.server == nil || r.server.Addr == "" {
		return ErrAddressNotSet
	}

	// check if routeset requirements are set
	if len(r.routeSets) < 1 {
		return ErrRouteSetNotPresent
	}

	// initialize middlewares
	r.router.Use(r.middlewares...)

	// assign the router to each route set
	for _, routeSet := range r.routeSets {
		routeSet.router = r.router
	}

	// assign the router
	r.server.Handler = r.router

	afterCh := time.After(500 * time.Millisecond)
	go func() {
		if err = r.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			r.logger.Error("server closed unexpectedly with error: %w", err)
			return
		}
	}()
	<-afterCh
	if err != nil {
		return err
	}
	r.logger.Info("server started on: %s", r.server.Addr)
	return nil
}

// Stop delegates the stop signal to http.Server.Shutdown
func (r *RouteMachine) Stop() error {
	return r.server.Shutdown(context.Background())
}
