package main

import (
	"fmt"
	"strings"
)

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	Firstname, Lastname string
	Age                 int
}

// A PersonBuilder helps with creating a person for different use cases using the fluent builder pattern.
type PersonBuilder struct {
	name   string
	bornAt int
}

// SetName takes an arbitrary text as a name which is parsed later.
func (b *PersonBuilder) SetName(text string) *PersonBuilder {
	b.name = text
	return b
}

// SetBornAt takes a year and calculates the age later.
func (b *PersonBuilder) SetBornAt(year int) *PersonBuilder {
	b.bornAt = year
	return b
}

// Build creates a new Person instance if correctly configured.
func (b *PersonBuilder) Build() (*Person, error) {
	parts := strings.SplitN(b.name, " ", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid name format: must have at least one empty space: %s", b.name)
	}

	return &Person{
		Firstname: parts[0],
		Lastname:  parts[1],
		Age:       2022 - b.bornAt,
	}, nil
}

func main() {
	var builder PersonBuilder
	p, err := builder.
		SetName("Torben Schinke").
		SetBornAt(1984).
		Build()

	if err != nil {
		panic(fmt.Errorf("unreachable state: %w", err))
	}

	fmt.Printf("%+v", p)
}
