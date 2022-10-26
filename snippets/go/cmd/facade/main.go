package main

import "fmt"

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	Firstname, Lastname string
	Age                 int
}

type somePackagePrivateSubsystem struct{}

func (somePackagePrivateSubsystem) doStuff(Person) {}

type someEvenMoreComplexSubsystem struct{}

func (someEvenMoreComplexSubsystem) doOtherStuff(Person) {}

type cannotBeUsedFromTheOutsideButImportantToPartOfYourProblemSolutions struct{}

func (cannotBeUsedFromTheOutsideButImportantToPartOfYourProblemSolutions) doDifferentStuff(Person) {}

// PersonFacade encapsulates a bunch of sub systems from this package and provides an easy-to-use API (application
// programming interface).
type PersonFacade struct {
}

// NewPersonFacade creates a new default person facade to marry people with each other using an arbitrary complex
// process.
func NewPersonFacade() *PersonFacade {
	return &PersonFacade{}
}

// Marry engages two persons with each other. This is part of your exported API layer.
func (p *PersonFacade) Marry(personA, personB Person) {
	var system1 somePackagePrivateSubsystem
	var system2 someEvenMoreComplexSubsystem
	var system3 cannotBeUsedFromTheOutsideButImportantToPartOfYourProblemSolutions

	system1.doStuff(personA)
	system2.doOtherStuff(personB)
	system3.doDifferentStuff(personA)

	fmt.Printf("%+v is married to %+v\n", personA, personB)
}

func main() {
	facade := NewPersonFacade()
	facade.Marry(Person{Firstname: "Sam"}, Person{Firstname: "Mary"})
}
