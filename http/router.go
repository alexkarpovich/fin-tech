package http

import (
	"net/http"
	"time"

	"github.com/Sharykhin/fin-tech/http/handlers/users"

	"github.com/Sharykhin/fin-tech/http/handlers"
	"github.com/gorilla/mux"
)

func router() http.Handler {
	r := mux.NewRouter()

	s := r.StrictSlash(true).PathPrefix("/api").Subrouter()
	s.HandleFunc("/users", users.Index).Methods("GET")
	s.HandleFunc("/_healthcheck", handlers.HealthCheck).Methods("GET")

	return s
}

// ListenAndServe starts listening http service on a provided addr
func ListenAndServe(addr string) error {
	s := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router(),
		Addr:         addr,
	}

	return s.ListenAndServe()
}