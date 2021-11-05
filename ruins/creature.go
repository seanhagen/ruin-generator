package ruins

import "github.com/justinian/dice"

/**
 * File: creature.go
 * Date: 2021-11-04 14:39:23
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

type creatureFeature struct {
	baseFeature
}

// Apply ...
func (cf creatureFeature) Apply(exit *Exit) (*Room, error) {
	return createCreature(exit.Parent.Level+1, exit)
}

func addCreature(room *Room) error {
	cr, _, err := dice.Roll("1d100")
	if err != nil {
		return err
	}

	room.Traits = append(room.Traits, BasicTrait{name: "Creature", roll: cr.Int()})
	return nil
}

// createCreature ...
func createCreature(level int, parent *Exit) (*Room, error) {
	room, err := rollRoom(level, parent)
	if err != nil {
		return nil, err
	}

	addCreature(room)

	return room, nil
}
