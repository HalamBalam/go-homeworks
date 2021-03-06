package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type player struct {
	name   string
	points int
}

var failPercent = 20
var finalScore = 11

func game(p *player, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		if val == "end" {
			break
		}
		time.Sleep(time.Second / 4)

		if val == "begin" {
			fmt.Println("--------------")
			fmt.Println(p.name + ": ping")
			ch <- "ping"
		}

		if val == "stop" {
			fmt.Println(p.name + ": ГОЛ!")
			p.points++
			if p.points == finalScore {
				ch <- "end"
				break
			} else {
				ch <- "begin"
			}
		}

		if val == "ping" {
			if rand.Intn(100) < failPercent {
				ch <- "stop"
			} else {
				fmt.Println(p.name + ": pong")
				ch <- "pong"
			}
		}

		if val == "pong" {
			if rand.Intn(100) < failPercent {
				ch <- "stop"
			} else {
				fmt.Println(p.name + ": ping")
				ch <- "ping"
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	p1 := player{name: "Роланд Якоби"}
	p2 := player{name: "Виктор Барна"}
	defer showScore(&p1, &p2)

	var wg sync.WaitGroup

	ch := make(chan string)

	wg.Add(2)
	go game(&p1, ch, &wg)
	go game(&p2, ch, &wg)

	ch <- "begin"
	wg.Wait()
}

func showScore(p1 *player, p2 *player) {
	winner := ""
	if p1.points == p2.points {
		winner = "Ничья"
	} else if p1.points > p2.points {
		winner = "Победил: " + p1.name
	} else {
		winner = "Победил: " + p2.name
	}
	fmt.Printf("Игра завершена! Счет: %d:%d\n", p1.points, p2.points)
	fmt.Println(winner)
}
