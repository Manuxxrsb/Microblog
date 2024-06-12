// internal/server/server.go
package server

import "net/http"

// Server is a base server configuration.
type Server struct {
	server *http.Server
}
