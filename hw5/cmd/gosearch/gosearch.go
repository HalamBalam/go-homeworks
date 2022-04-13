package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"homeworks/hw2/pkg/crawler/spider"
	"homeworks/hw2/pkg/index"
	"io"
	"os"
	"strings"
)

var scanner = spider.New()
var indPath = "./index.json"

func main() {
	urls := []string{"https://go.dev", "https://golang.org/"}

	var ind *index.Index
	var err error
	if _, err = os.Stat(indPath); err == nil {
		ind, err = readIndFromFile(indPath)
	} else {
		ind, err = fillInd(urls)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	console := bufio.NewScanner(os.Stdin)
	fmt.Print("Найти-> ")
	for console.Scan() {
		word := strings.ToLower(console.Text())
		if word == "exit" {
			break
		}

		findUrls := ind.Find(word)
		for _, url := range findUrls {
			fmt.Println(url)
		}
		fmt.Printf("Найдено %d ссылок\n", len(findUrls))
		fmt.Print("Найти-> ")
	}
}

func readIndFromFile(indPath string) (*index.Index, error) {
	file, err := os.Open(indPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return readInd(file)
}

func readInd(r io.Reader) (*index.Index, error) {
	ind := index.New()

	var data []byte
	var buf = make([]byte, 10)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return ind, err
		}
		if n > 0 {
			data = append(data, buf[0:n]...)
		}
	}

	err := json.Unmarshal(data, ind)
	return ind, err
}

func fillInd(urls []string) (*index.Index, error) {
	ind := index.New()
	for _, url := range urls {
		fmt.Println("Сканирование сайта " + url)

		docs, err := scanner.Scan(url, 2)
		if err != nil {
			continue
		}
		ind.AddDocs(docs)
	}

	file, err := os.Create(indPath)
	if err != nil {
		return ind, err
	}
	err = writeInd(ind, file)
	return ind, err
}

func writeInd(ind *index.Index, w io.Writer) error {
	data, err := json.Marshal(ind)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
