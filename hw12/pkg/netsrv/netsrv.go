package netsrv

import (
	"bufio"
	"homeworks/hw12/pkg/index"
	"net"
	"strings"
	"time"
)

type Service struct {
	ind *index.Index
}

func New(ind *index.Index) Service {
	return Service{ind: ind}
}

func (s *Service) Start() error {
	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handler(conn, s)
	}
}

func handler(conn net.Conn, s *Service) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Minute))

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		word := strings.ToLower(string(msg))

		findUrls := s.ind.Find(word)
		for _, url := range findUrls {
			_, err = conn.Write([]byte(url + "\n"))
			if err != nil {
				return
			}
		}
		_, err = conn.Write([]byte("\n"))
		if err != nil {
			return
		}
	}
}
