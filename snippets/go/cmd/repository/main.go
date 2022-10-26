package main

import "fmt"

// Person contains a first name, a last name and an age.
// This is our entity (see also the Domain-Driven Design definition).
// This is just a random model for demonstration.
type Person struct {
	ID                  int
	Firstname, Lastname string
	Age                 int
}

// PersonRepository describes the behavior of a repository which manages a bunch of entities.
// An entity must have a unique identifier. Also note, that each method may return an error, required for
// I/O based implementations (e.g. SQL).
// This is the central part of the pattern.
type PersonRepository interface {
	// FindAll is typical name, but it may be named with different verbs, like All or List.
	// It is usually better to return a value type instead of a pointer type to ensure
	// immutability of the returned entities. Otherwise, you have to create defensive copies.
	FindAll() ([]Person, error)

	// FindByID returns the person if found, otherwise an error.
	FindByID(id int) (Person, error)

	// many more management methods can be put here like
	// DeleteByID
	// RemoveAll, Clear
	// FindByLastname
	// FindOldest
}

// ExamplePersonRepository is a hardcoded immutable in-memory repository implementation.
// You don't know what the future brings, there may be implementations which work on MySQL, sqlite, PostgreSQL
// S3, REST and many more. Therefore, you separate behavior from implementation.
type ExamplePersonRepository struct {
	entities []Person
}

func NewExamplePersonRepository() *ExamplePersonRepository {
	return &ExamplePersonRepository{entities: []Person{
		{
			ID:        1,
			Firstname: "Frodo",
			Lastname:  "Beutlin",
			Age:       50,
		},
		{
			ID:        2,
			Firstname: "Bilbo",
			Lastname:  "Beutlin",
			Age:       111,
		},
		{
			ID:        3,
			Firstname: "Sam",
			Lastname:  "Gamdschie",
			Age:       38,
		},
	}}
}

// FindAll returns a copy of the internal slice, to ensure immutability.
func (r *ExamplePersonRepository) FindAll() ([]Person, error) {
	clonedSlice := make([]Person, 0, len(r.entities))
	for _, entity := range r.entities {
		clonedSlice = append(clonedSlice, entity)
	}

	return clonedSlice, nil
}

// FindByID searches for a specific entity.
func (r *ExamplePersonRepository) FindByID(id int) (Person, error) {
	for _, entity := range r.entities {
		if entity.ID == id {
			return entity, nil
		}
	}

	return Person{}, fmt.Errorf("person not found: %d", id)
}

func main() {
	// this is the interface, which you are always coding against
	var repo PersonRepository

	// initialize the abstract type with a concrete instance
	repo = NewExamplePersonRepository()

	// work with it, but use the interface
	// find many
	entities, err := repo.FindAll()
	if err != nil {
		panic(fmt.Errorf("unreachable: %w", err))
	}

	for _, entity := range entities {
		fmt.Printf("%+v\n", entity)
	}

	// find one
	bilbo, err := repo.FindByID(2)
	if err != nil {
		panic(fmt.Errorf("unreachable: %w", err))
	}

	fmt.Printf("bilbo is here: %+v\n", bilbo)
}
