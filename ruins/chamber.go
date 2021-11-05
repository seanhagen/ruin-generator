package ruins

import (
	"github.com/justinian/dice"
)

/**
 * File: chamber.go
 * Date: 2021-11-03 17:44:38
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// chamberFeature ...
type chamberFeature struct {
	baseFeature
}

// Apply ...
func (cf chamberFeature) Apply(exit *Exit) (*Room, error) {
	return createChamber(exit.Parent.Level+1, exit)
}

func getSizeAndShape() []Trait {
	sz, _, err := dice.Roll("1d20")
	if err != nil {
		return []Trait{}
	}
	sh, _, err := dice.Roll("1d20")
	if err != nil {
		return []Trait{}
	}

	size := BasicTrait{"Size", sz.Int()}
	shape := BasicTrait{"Shape", sh.Int()}

	return []Trait{size, shape}
}

func createChamber(level int, exit *Exit) (*Room, error) {
	traits := getSizeAndShape()

	e, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	if e.Int() > 15 {
		cf, _, err := dice.Roll("1d100")
		if err != nil {
			return nil, err
		}
		traits = append(traits, BasicTrait{"Chamber Feature", cf.Int()})
	}

	ro := &Room{
		ID:     nextRoomID(),
		Level:  level,
		Type:   "Chamber",
		Traits: traits,
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
