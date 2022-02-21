package person

import "github.com/jyangorch/hello-go/exercise-ddd-structure/internal/util"

// an aggregate object, that is mutable
type Person struct {
	id              string
	name            string
	currentLocation util.Location
}

// as aggregate object, constructor return pointer
func NewPerson(id string, name string) *Person {
	initLocation := util.NewLocation(0, 0)
	p := Person{id, name, initLocation}
	return &p
}

// as aggregate object, method references pointer recevier
func (p *Person) Id() string {
	return p.id
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) CurrentLocation() util.Location {
	return p.currentLocation
}

// as aggregate object, method references pointer recevier, as this method mutates object
func (p *Person) MoveBy(x int, y int) *Person {
	p.currentLocation = p.currentLocation.Add(x, y)
	return p
}
