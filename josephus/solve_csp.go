package main

import "fmt"

type action int

type person struct {
	pos   int
	alive bool
	left  chan action
	right chan action
}

const (
	DIE  action = 0
	PASS action = 1
)

func run(p person, totalAlive chan int, done chan bool) {
	fmt.Println(p)
	checkOneLeft(totalAlive, done)
}

func trackKills(totalAlive chan int) {
	left := <-totalAlive
	left--
	totalAlive <- left
}

func checkOneLeft(totalAlive chan int, done chan bool) {
	left := <-totalAlive
	if left == 1 {
		done <- true
	}
}

func solve(total int) {
	done := make(chan bool)
	totalAlive := make(chan int)
	go trackKills(totalAlive)
	var p person
	fmt.Println("reaches here", total)
	totalAlive <- total

	fmt.Println("never reaches here")
	var persons = make([]person, total)

	for i := 0; i < total; i++ {
		if i == 0 {
			p = person{
				pos:   i,
				alive: true,
				left:  make(chan action),
				right: make(chan action),
			}
		} else {
			p = person{
				pos:   i,
				alive: true,
				left:  make(chan action),
				right: persons[i-1].left,
			}
		}
		persons[i] = p

		persons[0].right = persons[total-1].left
	}

	for i := 0; i < total; i++ {
		go run(persons[i], totalAlive, done)
	}

	go checkOneLeft(totalAlive, done)
	finish := <-done
	fmt.Println("solve")
	fmt.Println(finish)
}

func main() {
	solve(6)
}
