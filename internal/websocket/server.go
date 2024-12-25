package websocket

import (
	pro "IMChat_App/pkg/proto"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"sync"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[int32]*Client
	mutex     *sync.RWMutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.RWMutex{},
		Clients:   make(map[int32]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

func (s *Server) Start() {

	for {
		select {
		case conn := <-s.Register:
			s.mutex.Lock()
			s.Clients[conn.Id] = conn
			s.mutex.Unlock()

			// Welcome message to the new client
			welcomeMsg := "Welcome!"
			conn.Send <- []byte(welcomeMsg)

		case conn := <-s.Ungister:
			s.mutex.RLock()
			if _, ok := s.Clients[conn.Id]; ok {
				close(conn.Send)
				s.mutex.RUnlock()
				s.mutex.Lock()
				delete(s.Clients, conn.Id)
				s.mutex.Unlock()
			} else {
				s.mutex.RUnlock()
			}

		case message := <-s.Broadcast:
			fmt.Println("来消息了")

			msg := &pro.Message{}
			err := proto.Unmarshal(message, msg)
			if err != nil {
				fmt.Println("反序列化失败")
			}

			// 发送消息到某个client
			s.mutex.RLock()
			if client, ok := s.Clients[msg.Receiver]; ok {
				client.Send <- message
				s.mutex.RUnlock()
			} else {
				s.mutex.RUnlock()
			}
		}
	}
}
