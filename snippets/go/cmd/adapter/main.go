package main

import "fmt"

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	Firstname, Lastname string
}

// Nameable is a behavior which is expected from a subsystem.
type Nameable interface {
	Name() string
}

// PersonWrapper is here to document, that the "wrapper pattern" is just the same as the "classic" adapter pattern.
type PersonWrapper = PersonNameableAdapter

// PersonNameableAdapter wraps and delegates to a Person. It also implements the Nameable interface, which actually
// makes this type an Adapter.
//
// WARNING: There are other definitions of an Adapter like the Adapter-View-Pattern or the
// Secondary and Primary-Adapter-Pattern (DDD) which have a more specialized meaning of "adapting".
type PersonNameableAdapter struct {
	delegate Person
}

func NewPersonNameableAdapter(person Person) PersonNameableAdapter {
	return PersonNameableAdapter{
		delegate: person,
	}
}

// Name is implemented by the adapter and wraps a Person.
// It delegates calls to the Person to assemble the actual name.
func (p PersonNameableAdapter) Name() string {
	return fmt.Sprintf("%s, %s", p.delegate.Lastname, p.delegate.Firstname)
}

func main() {
	p := Person{
		Firstname: "Torben",
		Lastname:  "Schinke",
	}

	adapter := NewPersonNameableAdapter(p)
	wrapper := adapter // often used for the same meaning

	fmt.Printf("%s\n", wrapper.Name())
}
