package ruins

import (
	"fmt"

	"github.com/justinian/dice"
)

/**
 * File: machine.go
 * Date: 2021-11-04 15:29:04
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// machineFeature  ...
type machineFeature struct {
	baseFeature
}

// Apply ...
func (mf machineFeature) Apply(exit *Exit) (*Room, error) {
	return createMachine(exit.Parent.Level+1, exit)
}

// machineTrait ...
type machineTrait struct {
	BasicTrait
	req string
}

// String ...
func (mt machineTrait) String() string {
	return fmt.Sprintf("%v: %v (%v)", mt.name, mt.req, mt.roll)
}

// createMachine ...
func createMachine(level int, exit *Exit) (*Room, error) {
	ch, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}
	r := ch.Int()

	var room *Room

	if r <= 33 {
		room, err = createCorridor(level, exit)
	}

	if r > 33 && r <= 95 {
		room, err = createChamber(level, exit)
	}

	if r > 95 {
		room, err = createVault(level, exit)
	}

	rep, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	if rep.Int() <= 30 {
		room.Traits = append(room.Traits,
			machineTrait{
				BasicTrait: BasicTrait{name: "Repair/Component Required", roll: rep.Int()},
				req:        "Yes",
			},
		)
	} else {
		room.Traits = append(room.Traits,
			machineTrait{
				BasicTrait: BasicTrait{name: "Repair/Component Required", roll: rep.Int()},
				req:        "No",
			},
		)
	}

	im, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	room.Traits = append(room.Traits,
		BasicTrait{name: "Machine", roll: im.Int()},
	)

	return room, err
}
