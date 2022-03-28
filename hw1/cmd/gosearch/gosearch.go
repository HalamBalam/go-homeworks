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

	docs := []crawler.Document{}
	sites := [...]string{"https://go.dev", "https://golang.org/"}
	for i := 0; i < len(sites); i++ {
		siteDocs, err := scan(sites[i])
		if err != nil {
			log.Println(err)
			return
		}
		for _, doc := range siteDocs {
			docs = append(docs, doc)
		}
	}

	if *s != "" {
		refs := findRefs(docs)
		for _, ref := range refs {
			fmt.Println(ref)
		}
		fmt.Printf("Найдено %d ссылок", len(refs))
	}
}

func init() {
	s = flag.String("s", "", "search text")
}

func scan(site string) ([]crawler.Document, error) {
	log.Println("Сканирование сайта " + site)

	scanner := spider.New()
	return scanner.Scan(site, 2)
}

func findRefs(docs []crawler.Document) ([]string) {
	res := []string{}
	for _, doc := range docs {
		if strings.Contains(doc.Title, *s) || strings.Contains(doc.Body, *s) {
			res = append(res, doc.URL)
		}
	}
	return res
}
