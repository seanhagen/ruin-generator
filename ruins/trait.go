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

// TraitType ...
type TraitType string

const (
	// TypeSize ...
	TypeSize TraitType = "Size"
	// TypeShape ...
	TypeShape = "Shape"
	// TypeCreature ...
	TypeCreature = "Creature"
)

// BasicTrait
type BasicTrait struct {
	name TraitType
	roll int
}

// Name ...
func (bt BasicTrait) Name() string {
	return bt.name
}

// String ...
func (bt BasicTrait) String() string {
	return fmt.Sprintf("%v: %v", bt.name, bt.roll)
}
