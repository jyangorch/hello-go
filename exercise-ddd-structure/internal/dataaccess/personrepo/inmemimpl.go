package personrepo

import (
	"errors"

	"github.com/jyangorch/hello-go/exercise-ddd-structure/pkg/domain/person"
)

type InMemoryImpl struct {
	storage map[string]*person.Person
}

func NewInMemoryImpl() *InMemoryImpl {
	r := InMemoryImpl{}
	r.storage = make(map[string]*person.Person)
	return &r
}

func (r *InMemoryImpl) GetPerson(personId string) (*person.Person, error) {
	if result, ok := r.storage[personId]; ok {
		return result, nil
	}
	return nil, errors.New("not found")
}

func (r *InMemoryImpl) SavePerson(p *person.Person) error {
	r.storage[p.Id()] = p
	return nil
}
