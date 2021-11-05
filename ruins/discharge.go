package ruins

import "github.com/justinian/dice"

/**
 * File: discharge.go
 * Date: 2021-11-04 15:42:15
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// dischargeFeature ...
type dischargeFeature struct {
	baseFeature
}

// Apply ...
func (lf dischargeFeature) Apply(exit *Exit) (*Room, error) {
	return createDischarge(exit.Parent.Level+1, exit)
}

// createDischarge ...
func createDischarge(level int, exit *Exit) (*Room, error) {
	room, err := rollRoom(level, exit)
	if err != nil {
		return nil, err
	}

	l, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	room.Traits = append(room.Traits,
		BasicTrait{name: "Energy Discharge", roll: l.Int()},
	)

	return room, nil
}
