package main

import (
	"homeworks/hw14/pkg/messages"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// Server - тип rpc-сервера.
type Server struct {
	mess   []messages.Message
	mu     sync.Mutex
	lastID int
}

// Send - добавляет список сообщений в хранилище.
func (s *Server) Send(req []messages.Message, _ *int) error {
	for _, m := range req {
		s.mu.Lock()
		m.ID = s.lastID + 1
		s.mess = append(s.mess, m)
		s.lastID++
		s.mu.Unlock()
	}
	return nil
}

// Messages - получает список сообщений из хранилища.
func (s *Server) Messages(_ []messages.Message, resp *[]messages.Message) error {
	*resp = s.mess
	return nil
}

func main() {
	srv := new(Server)
	err := rpc.Register(srv)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}
