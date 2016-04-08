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

func run() {

}

func trackKills(totalAlive chan int) {
	left := <-totalAlive
	left--
	totalAlive <- left
}

func checkOneLeft(totalAlive chan int) int {
	left := <-totalAlive
	return left
}

func solve(total int) {
	var done chan person
	var totalAlive chan int
	totalAlive <- total
	p := <-done
	fmt.Println("solve")
}

func main() {
	solve(6)

}
