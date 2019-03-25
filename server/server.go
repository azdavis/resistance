package main

// Server is the server.
type Server struct {
	C chan<- *Client
	q chan<- struct{}
}

// NewServer returns a new Server.
func NewServer() *Server {
	txLobbyMap := make(chan ToLobbyMap, 3)
	txWelcomer := make(chan *Client, 3)
	q := make(chan struct{})
	go runLobbyMap(txLobbyMap, q)
	go runWelcomer(txLobbyMap, txWelcomer, q)
	s := &Server{txWelcomer, q}
	return s
}

// Close shuts down the Server. It should only be called once.
func (s *Server) Close() {
	close(s.q)
}
