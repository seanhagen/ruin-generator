package ruins

import "github.com/justinian/dice"

/**
 * File: weird.go
 * Date: 2021-11-04 15:43:12
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// weirdFeature ...
type weirdFeature struct {
	baseFeature
}

// Apply ...
func (lf weirdFeature) Apply(exit *Exit) (*Room, error) {
	return createWeird(exit.Parent.Level+1, exit)
}

// createWeird ...
func createWeird(level int, exit *Exit) (*Room, error) {
	l, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	room := exit.Parent

	room.Traits = append(room.Traits,
		BasicTrait{name: "Weird Event", roll: l.Int()},
	)

	return room, nil
}
