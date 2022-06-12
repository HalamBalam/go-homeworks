package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"homeworks/hw18/pkg/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api.New(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
