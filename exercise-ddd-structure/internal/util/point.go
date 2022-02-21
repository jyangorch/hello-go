package util

import "fmt"

// an value object, that is immutable
type Location struct {
	left  int
	right int
}

// as value object, constructor return value instead of pointer
func NewLocation(left int, right int) Location {
	l := Location{left, right}
	return l
}

// as value object, method references "value receiver" as we don't mutate the object
func (l Location) Left() int {
	return l.left
}

func (l Location) Right() int {
	return l.right
}

func (l Location) ToString() string {
	return fmt.Sprintf("[%d, %d]", l.left, l.right)
}

func (l Location) Add(distanceX int, distanceY int) Location {
	return NewLocation(l.left+distanceX, l.right+distanceY)
}
