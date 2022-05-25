package main

import (
	"fmt"
	"homeworks/hw14/pkg/messages"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var req []messages.Message
	for i := 1; i <= 5; i++ {
		m := messages.Message{
			Time: time.Now(),
			Text: fmt.Sprintf("Message number: %d", i),
		}
		req = append(req, m)
		time.Sleep(time.Second)
	}

	err = client.Call("Server.Send", req, nil)
	if err != nil {
		log.Fatal(err)
	}

	var resp []messages.Message
	err = client.Call("Server.Messages", new([]messages.Message), &resp)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range resp {
		fmt.Println(m.String())
	}
}
