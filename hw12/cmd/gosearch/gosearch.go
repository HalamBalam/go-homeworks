package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"homeworks/hw12/pkg/crawler/spider"
	"homeworks/hw12/pkg/index"
	"homeworks/hw12/pkg/netsrv"
	"homeworks/hw12/pkg/webapp"
	"net/http"
)

var scanner = spider.New()
var ind *index.Index

func main() {
	urls := []string{"https://go.dev", "https://golang.org/"}

	ind = createInd(urls)

	srv := netsrv.New(ind)
	go srv.Start()

	r := mux.NewRouter()
	webapp.New(ind, r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
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
