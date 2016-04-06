package josephus

import "fmt"
import "container/ring"

type Person struct {
	isAlive bool
	pos     int
}

type jCircle struct {
	size   int
	circle *ring.Ring
	sword  int
}

func (self *jCircle) findSword(sword_pos int) {
	r := self.circle
	for i := 0; i < r.Len(); i++ {
		person := r.Value.(Person)
		if person.pos == sword_pos {
			r = r.Next()
			break
		}
	}
	self.circle = r
}

func makeJCircle(totalAlive int) *jCircle {
	j := jCircle{size: totalAlive, circle: ring.New(totalAlive)}
	return &j
}

func (self *jCircle) populate() {
	//size := self.size
	r := self.circle

	for i := 0; i < r.Len(); i++ {
		r.Value = Person{isAlive: true, pos: i}
		r = r.Next()
	}
	r.Do(func(x interface{}) {
		fmt.Println(x)
	})
}

func (self *jCircle) killNext() int {
	var person Person
	current_person := self.circle
	done := false
	for done != true {
		current_person.Next()
		person = current_person.Value.(Person)
		if person.isAlive == true {
			person.isAlive = false
			done = true
		}

	}
	return person.pos
}

func (self *jCircle) passSword() int {
	var person Person
	current_person := self.circle
	done := false
	for done != true {
		current_person.Next()
		person = current_person.Value.(Person)
		if person.isAlive == true {
			done = true
		}

	}
	return person.pos
}

func Simulate(totalAlive int) {
	j := makeJCircle(totalAlive)
	j.populate()
	fmt.Println(j)
	fmt.Println(j.size)
	sword_pos := 0
	for totalAlive > 1 {
		j.findSword(sword_pos)
		//		killed = j.killNext()
		//		j.findSword(killed)
		//		sword_pos = j.passSword()
		//jCircle.killNext()
		//jCircle.passSword()
		totalAlive--
	}
	fmt.Println(j.circle.Value)
	return
}
