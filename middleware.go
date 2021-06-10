package procroute

import "net/http"

// Middleware defines an interface that is used to inject a middleware
type Middleware interface {
	// Middleware represents the function that must be implemented to assign a new Middleware to the RouteMachine.
	// In most cases the implementation is similar to:
	//
	// Example:
	//  func (m *MyType) Middleware(http.Handler) http.Handler {
	//      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//          // do something
	//          next.ServeHTTP(w, r)
	//      })
	//  }
	Middleware(http.Handler) http.Handler
}
