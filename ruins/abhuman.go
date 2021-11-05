package ruins

import (
	"fmt"

	"github.com/justinian/dice"
)

/**
 * File: abhuman.go
 * Date: 2021-11-04 15:15:48
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// abhumanFeature ...
type abhumanFeature struct {
	baseFeature
}

// Apply ...
func (af abhumanFeature) Apply(exit *Exit) (*Room, error) {
	return createAbhuman(exit.Parent.Level+1, exit)
}

// createAbhuman ...
func createAbhuman(level int, exit *Exit) (*Room, error) {
	room, err := rollRoom(level, exit)
	if err != nil {
		return nil, err
	}

	room.Type = fmt.Sprintf("%v (with Abhuman Colony)", room.Type)

	d6t, _, err := dice.Roll("1d6")
	if err != nil {
		return nil, err
	}

	d10t, _, err := dice.Roll("1d10")
	if err != nil {
		return nil, err
	}

	d20t, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	spc, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	num, _, err := dice.Roll("1d20+8")
	if err != nil {
		return nil, err
	}

	shin, _, err := dice.Roll("2d20")
	if err != nil {
		return nil, err
	}

	room.Traits = append(room.Traits,
		BasicTrait{name: "Type (Corebook)", roll: d6t.Int()},
		BasicTrait{name: "Type (Bestiary)", roll: d10t.Int()},
		BasicTrait{name: "Type (Core+Bestiary)", roll: d20t.Int()},
		BasicTrait{name: "Specific", roll: spc.Int()},
		BasicTrait{name: "# of Abhumans", roll: num.Int()},
		BasicTrait{name: "# of shins", roll: shin.Int()},
	)

	return room, nil
}
