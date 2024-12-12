package websocket

import (
	"IMChat_App/pkg/protocol"
	"github.com/gogo/protobuf/proto"
	"sync"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[string]*Client
	mutex     *sync.Mutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[string]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

func (s *Server) Start() {

	for {
		select {
		case conn := <-s.Register:
			s.Clients[conn.Name] = conn

			// Welcome message to the new client
			welcomeMsg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "Welcome to the chat room!",
			}
			protoMsg, _ := proto.Marshal(welcomeMsg)
			conn.Send <- protoMsg

		case conn := <-s.Ungister:
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			// Process the incoming message and broadcast it to the intended recipient(s)
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				// Send to a specific user
				if client, ok := s.Clients[msg.To]; ok {
					client.Send <- message
				}
			} else {
				// Broadcast to all clients
				for _, conn := range s.Clients {
					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.Name)
					}
				}
			}
		}
	}
}
