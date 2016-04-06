package josephus

import "fmt"

type Circle struct {
	size int
}

type Person struct {
	isAlive bool
	pos     int
}

type jCircle struct {
	size   int
	circle Circle
}

func makeJCircle(totalAlive int) *jCircle {
	j := jCircle{size: totalAlive, circle: Circle{size: totalAlive}}
	return &j
}

func (self *jCircle) killNext() int {
	fmt.Println(self.size)
	return 1
}

func Simulate(totalAlive int) {
	j := makeJCircle(totalAlive)
	fmt.Println(j)
	fmt.Println(j.size)
	for totalAlive > 1 {
		j.killNext()
		//jCircle.killNext()
		//jCircle.passSword()
		totalAlive--
	}
	return
}
