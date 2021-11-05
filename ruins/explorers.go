package ruins

import (
	"fmt"

	"github.com/justinian/dice"
)

/**
 * File: explorers.go
 * Date: 2021-11-04 14:47:52
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// explorerFeature ...
type explorerFeature struct {
	baseFeature
}

// Apply ...
func (ef explorerFeature) Apply(exit *Exit) (*Room, error) {
	return createExplorer(exit.Parent.Level+1, exit)
}

// createExplorer ...
func createExplorer(level int, parent *Exit) (*Room, error) {
	room, err := rollRoom(level, parent)
	if err != nil {
		return nil, err
	}

	room.Type = fmt.Sprintf("%v (with Explorers)", room.Type)

	num, _, err := dice.Roll("1d6+1")
	if err != nil {
		return nil, err
	}

	sit, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	room.Traits = append(room.Traits,
		BasicTrait{name: "# of Explorers", roll: num.Int()},
		BasicTrait{name: "Situation", roll: sit.Int()},
	)

	return room, nil
}
