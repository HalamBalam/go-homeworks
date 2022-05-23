package main

import (
	"homeworks/hw14/pkg/messages"
	"log"
	"net"
	"net/rpc"
)

type Server int

var mess []messages.Message

func (s *Server) Send(req []messages.Message, _ *[]messages.Message) error {
	id := 1
	if len(mess) != 0 {
		id = mess[len(mess)-1].ID + 1
	}
	for _, m := range req {
		m.ID = id
		mess = append(mess, m)
		id++
	}
	return nil
}

func (s *Server) Messages(_ []messages.Message, resp *[]messages.Message) error {
	*resp = mess
	return nil
}

func main() {
	srv := new(Server)
	err := rpc.Register(srv)
	if err != nil {
		log.Fatal(err)
	}

	// регистрация сетевой службы RPC-сервера
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}
