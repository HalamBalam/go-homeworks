package index

import (
	"sort"
	"strings"
	"homeworks/hw2/pkg/crawler"
)

type Index struct{
	Docs []crawler.Document
	Data map[string][]int
	LastID int
}

func New() *Index {
	i := Index{}
	i.Data = make(map[string][]int)
	i.LastID = 0
	return &i
}

func (i *Index) AddDocs(docs []crawler.Document) {
	for _, doc := range docs {
		i.LastID++
		doc.ID = i.LastID
		i.Docs = append(i.Docs, doc)

		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			if len(word) > 1 && findID(i.Data[word], doc.ID) == -1 {
				i.Data[word] = append(i.Data[word], doc.ID)
			}
		}
	}

	sort.Slice(i.Docs, func(a, b int) bool {
	  return i.Docs[a].ID < i.Docs[b].ID
	})
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

func findID(data []int, ID int) int {
	for i := range data {
		if data[i] == ID {
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
