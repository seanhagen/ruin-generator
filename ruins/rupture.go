package ruins

import (
	"github.com/justinian/dice"
)

/**
 * File: rupture.go
 * Date: 2021-11-04 14:45:13
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// ruptureFeature ...
type ruptureFeature struct {
	baseFeature
}

// Apply ...
func (inf ruptureFeature) Apply(exit *Exit) (*Room, error) {
	return createRupture(exit.Parent.Level+1, exit)
}

// createRupture ...
func createRupture(level int, exit *Exit) (*Room, error) {
	exit.Type = ExitRupture

	ru, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	ro := &Room{
		ID:     nextRoomID(),
		Level:  level,
		Type:   "Rupture",
		Traits: []Trait{BasicTrait{name: "Rupture", roll: ru.Int()}},
	}

	if exit != nil {
		if exit.Parent != nil {
			ro.Parent = exit.Parent
		}
		exit.Child = ro
	}

	err = addExits(ro)
	return ro, err
}
