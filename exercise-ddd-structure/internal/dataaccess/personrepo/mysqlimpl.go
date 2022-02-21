package personrepo

import (
	"github.com/jyangorch/hello-go/exercise-ddd-structure/pkg/domain/person"
)

type MySqlImpl struct {
}

func NewPersonRepositoryMySqlImpl() *MySqlImpl {
	return &MySqlImpl{}
}

func (r *MySqlImpl) GetPerson(personId string) (*person.Person, error) {
	panic("not implemented")
}

func (r *MySqlImpl) SavePerson(p *person.Person) error {
	panic("not implemented")
}
