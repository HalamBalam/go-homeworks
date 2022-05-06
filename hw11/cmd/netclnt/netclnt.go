package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	console := bufio.NewScanner(os.Stdin)
	fmt.Print("Найти-> ")
	for console.Scan() {
		word := strings.ToLower(console.Text())
		if word == "exit" {
			break
		}

		_, err := conn.Write([]byte(word + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(conn)
		for {
			b, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}
			if len(b) == 1 {
				break
			}
			fmt.Print(string(b))
		}
		fmt.Print("Найти-> ")
	}
}
