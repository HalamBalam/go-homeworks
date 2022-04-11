package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"homeworks/hw2/pkg/crawler/spider"
	"homeworks/hw2/pkg/index"
	"io"
	"os"
)

var scanner = spider.New()
var indPath = "./index.json"

func main() {
	urls := []string{"https://go.dev", "https://golang.org/"}

	ind := index.New()
	var err error
	if _, err = os.Stat(indPath); err == nil {
		err = readIndFromFile(ind)
	} else {
		err = fillInd(ind, urls)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	console := bufio.NewScanner(os.Stdin)
	fmt.Print("Найти-> ")
	for console.Scan() {
		word := console.Text()
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

func readIndFromFile(ind *index.Index) error {
	file, err := os.Open(indPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return readInd(ind, file)
}

func readInd(ind *index.Index, r io.Reader) error {
	var data []byte
	var buf = make([]byte, 10)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if n > 0 {
			data = append(data, buf[0:n]...)
		}
	}

	return json.Unmarshal(data, ind)
}

func fillInd(ind *index.Index, urls []string) error {
	for _, url := range urls {
		fmt.Println("Сканирование сайта " + url)

		docs, err := scanner.Scan(url, 2)
		if err != nil {
			return err
		}
		ind.AddDocs(docs)
	}

	file, err := os.Create(indPath)
	if err != nil {
		return err
	}
	return writeInd(ind, file)
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
