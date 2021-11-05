package ruins

import (
	"github.com/justinian/dice"
)

/**
 * File: shaft.go
 * Date: 2021-11-04 14:44:33
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// shaftFeature ...
type shaftFeature struct {
	baseFeature
}

// Apply ...
func (inf shaftFeature) Apply(exit *Exit) (*Room, error) {
	return createShaft(exit.Parent.Level+1, exit)
}

// createShaft ...
func createShaft(level int, parent *Exit) (*Room, error) {
	s, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	traits := []Trait{
		BasicTrait{name: "Type", roll: s.Int()},
	}

	ro := &Room{
		ID:     nextRoomID(),
		Level:  level,
		Type:   "Shaft",
		Traits: traits,
	}

	addExits(ro)

	for _, e := range ro.Exits {
		e.Type = ExitShaft
	}

	ro.Exits = append(ro.Exits, &Exit{nextExitID(), ExitNormal, ro, nil})

	return ro, nil
}
