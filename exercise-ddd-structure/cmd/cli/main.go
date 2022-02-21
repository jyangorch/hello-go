package main

import (
	"github.com/jyangorch/hello-go/exercise-ddd-structure/internal/dataaccess/personrepo"
	"github.com/jyangorch/hello-go/exercise-ddd-structure/pkg/domain/greet"
	"github.com/jyangorch/hello-go/exercise-ddd-structure/pkg/domain/person"
)

func main() {
	var repo person.PersonRepository = personrepo.NewInMemoryImpl()

	// save to repo
	alice := person.NewPerson("001", "Alice")
	bob := person.NewPerson("002", "Bob")
	repo.SavePerson(alice)
	repo.SavePerson(bob)

	// get from repo
	p1, _ := repo.GetPerson("001")
	p2, _ := repo.GetPerson("002")
	persons := []*person.Person{p1, p2}

	// invoke a service
	var greeter greet.Greeter = greet.Greeter{}
	greeter.MoveAndGreet(persons, 100, 200)
}
