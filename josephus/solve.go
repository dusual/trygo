package josephus

import "fmt"
import "container/ring"

type person struct {
	alive bool
	pos   int
}

func (self *person) isAlive() bool {
	return self.alive
}

func (self *person) getPos() int {
	return self.pos
}

func (self *person) kill() {
	self.alive = false
}

type jCircle struct {
	current *ring.Ring
	nAlive  int
}

func makeJCircle(totalAlive int) *jCircle {
	j := &jCircle{nAlive: totalAlive, current: ring.New(totalAlive)}
	j.populate()
	return j
}

func (self *jCircle) populate() {
	current := self.current
	for i := 0; i < current.Len(); i++ {
		current.Value = &person{alive: true, pos: i}
		current = current.Next()
	}
}

func (self *jCircle) killNext() {
	done := false
	for done != true {
		self.current = self.current.Next()
		next_person := self.current.Value.(*person)
		if next_person.isAlive() {
			next_person.kill()
			done = true
		}
	}
}

func (self *jCircle) passSword() {
	done := false
	for done != true {
		self.current = self.current.Next()
		next_person := self.current.Value.(*person)
		if next_person.isAlive() {
			done = true
		}
	}
}

func Simulate(totalAlive int) {
	j := makeJCircle(totalAlive)
	for j.nAlive > 1 {
		j.killNext()
		j.passSword()
		j.nAlive--
	}
	fmt.Println(j.current.Value.(*person).pos)
	return
}
