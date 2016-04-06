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

func makeJCircle(totalAlive int, sword int) *jCircle {
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

}

func Simulate(totalAlive int) {
	j := makeJCircle(totalAlive)
	j.populate()
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
