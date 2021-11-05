package ruins

import "github.com/justinian/dice"

/**
 * File: matter_leak.go
 * Date: 2021-11-04 15:39:08
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// leakFeature ...
type leakFeature struct {
	baseFeature
}

// Apply ...
func (lf leakFeature) Apply(exit *Exit) (*Room, error) {
	return createLeak(exit.Parent.Level+1, exit)
}

// createLeak ...
func createLeak(level int, exit *Exit) (*Room, error) {
	room, err := rollRoom(level, exit)
	if err != nil {
		return nil, err
	}

	l, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	room.Traits = append(room.Traits,
		BasicTrait{name: "Matter Leak", roll: l.Int()},
	)

	return room, nil
}
