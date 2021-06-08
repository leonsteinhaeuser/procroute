package main

import (
	"net/http"

	"github.com/leonsteinhaeuser/procroute"
)

type MyExampleMiddleware struct {
	logger procroute.Loggable
}

func (m *MyExampleMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Debug("method: %s path: %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func (m *MyExampleMiddleware) WithLogger(lggbl procroute.Loggable) {
	m.logger = lggbl
}
