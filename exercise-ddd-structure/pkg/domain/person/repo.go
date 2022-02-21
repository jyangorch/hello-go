package person

// repository interface for person
type PersonRepository interface {

	// Get person by id
	GetPerson(personId string) (*Person, error)

	// Save person
	SavePerson(person *Person) error
}
