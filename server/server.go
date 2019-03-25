package main

import (
	"context"
	"net/http"
	"time"
)

// Server is the server.
type Server struct {
	q          chan<- struct{}
	HTTPServer *http.Server
}

// NewServer returns a new Server.
func NewServer() *Server {
	txLobbyMap := make(chan ToLobbyMap, 3)
	txWelcomer := make(chan *Client, 3)
	q := make(chan struct{})
	go runLobbyMap(txLobbyMap, q)
	go runWelcomer(txLobbyMap, txWelcomer, q)
	hs := &http.Server{
		Handler:      NewHub(txWelcomer),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s := &Server{q, hs}
	return s
}

// Close shuts down the Server. It should only be called once.
func (s *Server) Close() error {
	err := s.HTTPServer.Shutdown(context.Background())
	close(s.q)
	return err
}
