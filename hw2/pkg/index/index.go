package index

import (
	"log"
	"sort"
	"strings"
	"homeworks/hw2/pkg/crawler"
	"homeworks/hw2/pkg/crawler/spider"
)

type Index struct{
	Docs []crawler.Document
	Data map[string][]int
}

var scanner *spider.Service = spider.New()

func New() *Index {
	i := Index{}
	i.Data = make(map[string][]int)
	return &i
}

func (i *Index) Scan(urls []string) error {
	id := 0
	for _, url := range urls {
		docs, err := scan(url)
		if err != nil {
			return err
		}
		for _, doc := range docs {
			id++
			doc.ID = id
			i.Docs = append(i.Docs, doc)

			words := strings.Split(doc.Title, " ")
			for _, word := range words {
				if len(word) > 1 && find(i.Data[word], doc.ID) == -1 {
					i.Data[word] = append(i.Data[word], doc.ID)
				}
			}
		}
	}

	sort.Slice(i.Docs, func(a, b int) bool {
	  return i.Docs[a].ID < i.Docs[b].ID
	})

	return nil
}

func (i *Index) Find(word string) []string {
	var res []string
	numbers := i.Data[word]
	for _, num := range numbers {
		doc, ok := findDoc(i.Docs, num)
		if ok {
			res = append(res, doc.URL)
		}
	}
	return res
}

func scan(url string) ([]crawler.Document, error) {
	log.Println("Сканирование сайта " + url)

	return scanner.Scan(url, 2)
}

func find(data []int, item int) int {
	for i := range data {
		if data[i] == item {
			return i
		}
	}
	return -1
}

func findDoc(data []crawler.Document, item int) (crawler.Document, bool) {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].ID == item {
			return data[mid], true
		}
		if data[mid].ID < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return crawler.Document{}, false
}
