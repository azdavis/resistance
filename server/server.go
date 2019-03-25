package main

import (
	"context"
	"net/http"
	"time"
)

// Server is the server.
type Server struct {
	HTTPServer *http.Server
}

// NewServer returns a new Server.
func NewServer() *Server {
	txLobbyMap := make(chan ToLobbyMap, 3)
	txWelcomer := make(chan *Client, 3)
	go runLobbyMap(txLobbyMap)
	go runWelcomer(txLobbyMap, txWelcomer)
	hs := &http.Server{
		Handler:      NewHub(txWelcomer),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s := &Server{hs}
	return s
}

// Close shuts down the Server. It should only be called once.
func (s *Server) Close() {
	s.HTTPServer.Shutdown(context.Background())
}
