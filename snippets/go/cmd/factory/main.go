package main

import "fmt"

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	Firstname, Lastname string
	Age                 int
}

// A PersonFactory demonstrate the most simple factory pattern (not the abstract factory pattern by GoF).
// Actually it just replaces a constructor. This is the actual pattern.
// The line between a builder and a factory is not clear, however a factory is usually less about optional
// configuration and more about a singleton lifetime.
type PersonFactory struct {
	lastAge int
}

// NewPersonFactory initializes a valid PersonFactory.
func NewPersonFactory() *PersonFactory {
	return &PersonFactory{lastAge: 1}
}

// CreatePerson makes some robots.
func (p *PersonFactory) CreatePerson() *Person {
	p.lastAge++
	return &Person{Firstname: "R2", Lastname: "D2", Age: p.lastAge}
}

func main() {
	factory := NewPersonFactory()
	for i := 0; i < 10; i++ {
		fmt.Printf("%+v\n", factory.CreatePerson())
	}
}
