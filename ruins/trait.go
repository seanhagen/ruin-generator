package ruins

import "fmt"

/**
 * File: trait.go
 * Date: 2021-11-03 16:58:30
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// Trait contains information about a room, such as size, shape, and any
// special information about what else may be found in that room
type Trait interface {
	Name() string
	String() string
}

// BasicTrait
type BasicTrait struct {
	name string
	roll int
}

// Name ...
func (bt BasicTrait) Name() string {
	return string(bt.name)
}

// Roll ...
func (bt BasicTrait) Roll() int {
	return bt.roll
}

// String ...
func (bt BasicTrait) String() string {
	return fmt.Sprintf("%v: %v", bt.name, bt.roll)
}

// StringTrait ...
type StringTrait struct {
	name, value string
}

// Name ...
func (st StringTrait) Name() string {
	return st.name
}

// String ...
func (st StringTrait) String() string {
	return fmt.Sprintf("%v: %v", st.name, st.value)
}
