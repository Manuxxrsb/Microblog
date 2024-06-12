// internal/server/server.go
package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}
