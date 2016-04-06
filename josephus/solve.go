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
	fmt.Println("find sword")
	r := self.circle
	for i := 0; i < r.Len(); i++ {
		person := r.Value.(Person)
		if person.pos == sword_pos {
			r = r.Next()
			break
		}
	}
	self.circle = r
	fmt.Println(self.circle)
}

func makeJCircle(totalAlive int) *jCircle {
	j := jCircle{size: totalAlive, circle: ring.New(totalAlive)}
	return &j
}

func (self *jCircle) populate() {
	r := self.circle

	for i := 0; i < r.Len(); i++ {
		r.Value = Person{isAlive: true, pos: i}
		r = r.Next()
	}
	// r.Do(func(x interface{}) {
	// 	fmt.Println(x)
	// })
}

func (self *jCircle) killNext() int {
	var person Person
	current_person := self.circle
	done := false
	for done != true {
		person = current_person.Value.(Person)
		fmt.Println(person)
		if person.isAlive == true {
			current_person.Value = Person{isAlive: false, pos: person.pos}
			done = true
		}
		current_person = current_person.Next()

	}
	fmt.Println("kill from kill")
	fmt.Println(person.pos)
	return person.pos
}

func (self *jCircle) passSword() int {
	fmt.Println("from pass")
	var person Person
	current_person := self.circle
	done := false
	for done != true {
		person = current_person.Value.(Person)
		fmt.Println(person)
		if person.isAlive == true {
			done = true
		}
		current_person = current_person.Next()
	}
	return person.pos
}

func Simulate(totalAlive int) {
	j := makeJCircle(totalAlive)
	j.populate()
	// fmt.Println(j)
	// fmt.Println(j.size)
	sword_pos := 0
	for totalAlive > 1 {
		j.findSword(sword_pos)
		killed := j.killNext()
		fmt.Println("killed")
		fmt.Println(killed)
		j.findSword(killed)
		sword_pos = j.passSword()
		fmt.Println("passed")
		fmt.Println(sword_pos)
		totalAlive--
	}
	fmt.Println(j.circle.Value)
	return
}
