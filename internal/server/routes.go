package server

import (
	"net/http"
)

// Routes defines all our routes
func (s *Server) Routes() {
	s.Router.Path("/myapi").Methods("GET", "OPTIONS").
		HandlerFunc(corsHandler(s.handleReadMyAPI()))
	s.Router.Path("/error").Methods("GET", "OPTIONS").
		HandlerFunc(corsHandler(s.handleReadError()))
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
