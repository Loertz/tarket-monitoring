package main

// SocketPool is the struct that keep track of all Client connection
type SocketPool struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// NewSocketPool : Factory for the socket pool
func NewSocketPool() *SocketPool {
	return &SocketPool{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Just a loop checking actionable on the different channel
// Don't worry, it's running concurently, and it's not blocking
// the main thread
func (s *SocketPool) run() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
		case message := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}
}
