package greet

import (
	"fmt"

	"github.com/jyangorch/hello-go/exercise-ddd-structure/pkg/domain/person"
)

// A service object, that operates on domain object
type Greeter struct {
}

func (g *Greeter) MoveAndGreet(persons []*person.Person, moveByX int, moveByY int) {
	for _, p := range persons {
		p.MoveBy(moveByX, moveByY)
		fmt.Printf("Person %s is currently at location: %s\n",
			p.Name(), p.CurrentLocation().ToString())
	}
}
