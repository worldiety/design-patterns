package main

import "fmt"

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	Firstname, Lastname string
	Age                 int
}

// NewPerson allocates a memory segment for a Person and returns a pointer to it.
// It is guaranteed, that neither first nor second name is the empty string or that age is negative.
// This is the actual pattern.
func NewPerson(first, second string, age int) (*Person, error) {
	if first == "" || second == "" {
		return nil, fmt.Errorf("names must not be empty")
	}

	if age < 0 {
		return nil, fmt.Errorf("invalid age")
	}

	return &Person{Firstname: first, Lastname: second}, nil
}

func main() {
	// fields are initialized with zero values
	p1 := &Person{}
	// factory function or "constructor" enforces validation
	p2, err := NewPerson("Torben", "Schinke", 38)
	if err != nil {
		// throwing a panic is only allowed for programming-errors
		panic(fmt.Errorf("this is unreachable code: %w", err))
	}

	fmt.Printf("%+v\n%+v\n", p1, p2)
}
