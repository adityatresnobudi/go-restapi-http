package server

import (
	"net/http"
	"strings"

	httpSwagger "github.com/swaggo/http-swagger"
)

func applicationJsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/swagger/") {
			w.Header().Set("Content-Type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}

func (s *server) runHTTPServer() error {
	s.mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	return http.ListenAndServe(s.cfg.Http.Port, applicationJsonResponseMiddleware(s.mux))
}
