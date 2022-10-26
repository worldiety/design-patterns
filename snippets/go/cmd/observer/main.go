package main

import "fmt"

// Person contains a first name, a last name and an age.
// Just a random model for demonstration.
type Person struct {
	firstname, lastname string
	age                 int // age is a private field

	lastHandle   ObserverHandle
	ageObservers map[ObserverHandle]Observer[int]
}

// NewPerson creates an observable person.
func NewPerson() *Person {
	return &Person{ageObservers: map[ObserverHandle]Observer[int]{}}
}

// ObserverHandle represents a context specific handle to identify an Observer.
// In Go, a function is not comparable. Although functions are pointer types, they are not comparable because
// the compiler may reuse functions in various situations, so that we cannot identify them safely.
type ObserverHandle int

// An Observer may be a generic function contract, but it may also be an interface.
type Observer[T any] func(oldValue, newValue T)

// AddAgeObserver observes the age field (or attribute).
func (p *Person) AddAgeObserver(observer Observer[int]) ObserverHandle {
	p.lastHandle++
	p.ageObservers[p.lastHandle] = observer
	return p.lastHandle
}

// RemoveAgeObserver delete the observer from the notification list.
func (p *Person) RemoveAgeObserver(hnd ObserverHandle) {
	delete(p.ageObservers, hnd)
}

// SetAge updates the age and notifies all registered observers.
func (p *Person) SetAge(newAge int) {
	old := p.age
	p.age = newAge
	for _, observer := range p.ageObservers {
		observer(old, newAge)
	}
}

func main() {
	p := NewPerson()
	p.SetAge(12) // nobody has been registered, so nothing special happens
	hnd := p.AddAgeObserver(func(oldValue, newValue int) {
		fmt.Printf("%d -> %d\n", oldValue, newValue)
	})

	// this will call the observer above
	p.SetAge(38)

	// unregister the observer
	p.RemoveAgeObserver(hnd)

	// nothing is printed anymore
	p.SetAge(39)
}
