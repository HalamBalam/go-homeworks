package main

import (
	"homeworks/hw11/pkg/netsrv"
	"log"
)

func main() {
	err := netsrv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
