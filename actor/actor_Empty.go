package actor

import (
	"fmt"
	"github.com/GUMI-golang/gumi"
)

// Actor::Empty
//
// Empty exists only for the GUMI Tree as an element that does nothing
type Empty struct {
	gumi.ParentGUMI
}

func (s Empty) MaxChildrun() int {
	return 0
}
func (s Empty) String() string {
	return fmt.Sprintf("%s", "Empty")
}

// Constructor
func Empty0() *Empty {
	return &Empty{}
}
