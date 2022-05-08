package netsrv

import (
	"bufio"
	"fmt"
	"homeworks/hw11/pkg/crawler/spider"
	"homeworks/hw11/pkg/index"
	"net"
	"strings"
	"time"
)

var scanner = spider.New()
var ind *index.Index

func Start() error {
	urls := []string{"https://go.dev", "https://golang.org/"}

	ind = createInd(urls)

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
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Minute))

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		word := strings.ToLower(string(msg))

		findUrls := ind.Find(word)
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

func createInd(urls []string) *index.Index {
	ind := index.New()
	for _, url := range urls {
		fmt.Println("Сканирование сайта " + url)

		docs, err := scanner.Scan(url, 2)
		if err != nil {
			continue
		}
		ind.AddDocs(docs)
	}

	return ind
}
