package main

import (
	"fmt"
	"homeworks/hw12/pkg/crawler/spider"
	"homeworks/hw12/pkg/index"
	"homeworks/hw12/pkg/netsrv"
	"homeworks/hw12/pkg/webapp"
)

var scanner = spider.New()
var ind *index.Index

func main() {
	urls := []string{"https://go.dev", "https://golang.org/"}

	ind = createInd(urls)

	netsrv.Ind = ind
	go netsrv.Start()

	webapp.Ind = ind
	err := webapp.Start()
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
