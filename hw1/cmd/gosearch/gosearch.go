package main

import (
	"log"
	"fmt"
	"flag"
	"strings"
	"hw1/pkg/crawler"
	"hw1/pkg/crawler/spider"
)

var s *string

func main() {
	flag.Parse()

	docs := make([]crawler.Document, 30)
	scan(&docs, "https://go.dev")
	scan(&docs, "https://golang.org/")

	if *s != "" {
		find_refs(&docs)
	}
}

func init() {
	s = flag.String("s", "", "search text")
}

func scan(docs *[]crawler.Document, site string) {
	log.Println("Сканирование сайта " + site)

	scanner := spider.New()
	new_docs, err := scanner.Scan(site, 2)
	if err != nil {
		log.Println(err)
		return
	}

	for _, doc := range new_docs {
		*docs = append(*docs, doc)
	}
}

func find_refs(docs *[]crawler.Document) {
	i := 0
	for _, doc := range *docs {
		if strings.Contains(doc.Title, *s) || strings.Contains(doc.Body, *s) {
			fmt.Println(doc.URL)
			i++
		}
	}

	fmt.Printf("Найдено %d ссылок", i)
}
