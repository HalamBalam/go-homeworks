package index

import (
	"homeworks/hw13/pkg/crawler"
	"sort"
	"strings"
)

// Index - обратный индекс документов.
type Index struct {
	Docs   []crawler.Document
	Data   map[string][]int
	LastID int
}

// New - конструктор обратного индекса.
func New() *Index {
	i := Index{}
	i.Data = make(map[string][]int)
	i.LastID = 0
	return &i
}

// AddDocs - добавляет список документов к существующему индексу.
func (i *Index) AddDocs(docs []crawler.Document) {
	for _, doc := range docs {
		i.LastID++
		doc.ID = i.LastID
		i.Docs = append(i.Docs, doc)

		words := strings.Split(doc.Title, " ")
		for _, word := range words {
			indWord := strings.ToLower(word)
			if len(indWord) > 1 && findID(i.Data[indWord], doc.ID) == -1 {
				i.Data[indWord] = append(i.Data[indWord], doc.ID)
			}
		}
	}

	sort.Slice(i.Docs, func(a, b int) bool {
		return i.Docs[a].ID < i.Docs[b].ID
	})
}

// Find - ищет в индексе слово, возвращает список найденных url.
func (i *Index) Find(word string) []string {
	var res = make([]string, 0)
	ids := i.Data[word]
	for _, id := range ids {
		if _, doc, ok := i.FindDoc(id); ok {
			res = append(res, doc.URL)
		}
	}
	return res
}

// FindDoc - ищет в индексе документ по идентификатору. Возвращает позицию и сам документ.
func (i *Index) FindDoc(id int) (int, *crawler.Document, bool) {
	low, high := 0, len(i.Docs)-1
	for low <= high {
		mid := (low + high) / 2
		if i.Docs[mid].ID == id {
			return mid, &i.Docs[mid], true
		}
		if i.Docs[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, nil, false
}

func findID(data []int, ID int) int {
	for i := range data {
		if data[i] == ID {
			return i
		}
	}
	return -1
}
