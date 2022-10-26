package main

import "fmt"

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	Firstname, Lastname string
	Age                 int
	Address             Address
}

// Address is just an example of the most simple composite pattern.
type Address struct {
	Street  string
	ZipCode int
}

func main() {
	p := Person{
		Firstname: "Torben",
		Lastname:  "Schinke",
		Age:       38,
		Address: Address{
			Street:  "Marie-Curie-Str. 1",
			ZipCode: 26129,
		},
	}

	fmt.Printf("%+v", p)
}
