package main

import (
	"log"
	"fmt"
	"flag"
	"strings"
	"homeworks/hw1/pkg/crawler"
	"homeworks/hw1/pkg/crawler/spider"
)

var s *string = flag.String("s", "", "search text")
var scanner *spider.Service = spider.New()

func main() {
	flag.Parse()

	var allDocs []crawler.Document
	sites := [...]string{"https://go.dev", "https://golang.org/"}
	for _, site := range sites {
		docs, err := scan(site)
		if err != nil {
			log.Println(err)
			return
		}
		for _, doc := range docs {
			allDocs = append(allDocs, doc)
		}
	}

	if *s != "" {
		refs := findRefs(allDocs)
		for _, ref := range refs {
			fmt.Println(ref)
		}
		fmt.Printf("Найдено %d ссылок", len(refs))
	}
}

func scan(site string) ([]crawler.Document, error) {
	log.Println("Сканирование сайта " + site)

	return scanner.Scan(site, 2)
}

func findRefs(docs []crawler.Document) ([]string) {
	var res []string
	for _, doc := range docs {
		if strings.Contains(doc.Title, *s) || strings.Contains(doc.Body, *s) {
			res = append(res, doc.URL)
		}
	}
	return res
}
