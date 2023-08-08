package personality

import "fmt"

type Human interface {
	Talk()
}

type Person struct {
	Name   string
	Age    int
	Origin string
}

func (p Person) Talk() {
	fmt.Printf("%v of %v years old is from %v", p.Name, p.Age, p.Origin)
}

func Id(h Human) {
	h.Talk()
}

func Details() {
	person := Person{
		Name:   "Emong",
		Age:    58,
		Origin: "Kumi",
	}

	person.Talk()
}
