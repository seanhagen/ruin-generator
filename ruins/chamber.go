package ruins

import "github.com/justinian/dice"

/**
 * File: chamber.go
 * Date: 2021-11-03 17:44:38
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func createChamber(level int, parent *Exit) (*Room, error) {
	sz, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}
	sh, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	size := BasicTrait{"Size", sz.Int()}
	shape := BasicTrait{"Shape", sh.Int()}

	traits := []Trait{size, shape}

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
		Level:  level + 1,
		Type:   "Chamber",
		Traits: traits,
	}

	err = addExits(ro)
	if err != nil {
		return nil, err
	}

	parent.Child = ro

	return ro, nil
}
