package main

import (
	"fmt"
	"bufio"
	"os"
	"homeworks/hw2/pkg/index"
)

func main() {
	urls := []string{"https://go.dev", "https://golang.org/"}

	ind := index.New()
	err := ind.Scan(urls)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Найти-> ")
	for scanner.Scan() {
		word := scanner.Text()
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
