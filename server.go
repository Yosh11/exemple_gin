package web

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server pointer on server
type Server struct {
	httpServer *http.Server
}

// Run start server
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	fmt.Printf("[INFO] Listening and serving HTTP on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown ...
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
