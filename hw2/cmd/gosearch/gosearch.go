package main

import (
	"fmt"
	"bufio"
	"os"
	"homeworks/hw2/pkg/index"
	"homeworks/hw2/pkg/crawler/spider"
)

var scanner *spider.Service = spider.New()

func main() {
	urls := []string{"https://go.dev", "https://golang.org/"}

	ind := index.New()
	err := fillInd(ind, urls)
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

func fillInd(ind *index.Index, urls []string) error {
	for _, url := range urls {
		fmt.Println("Сканирование сайта " + url)

		docs, err := scanner.Scan(url, 2)
		if err != nil {
			return err
		}

		ind.AddDocs(docs)
	}

	return nil
}
