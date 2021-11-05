package ruins

import (
	"github.com/justinian/dice"
)

/**
 * File: accessway.go
 * Date: 2021-11-04 15:00:41
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// accesswayFeature ...
type accesswayFeature struct {
	baseFeature
}

// Apply ...
func (ef accesswayFeature) Apply(exit *Exit) (*Room, error) {
	return createAccessway(exit.Parent.Level+1, exit)
}

func createAccessway(level int, exit *Exit) (*Room, error) {
	return _createAccessway(0, exit.Parent.Level+1, exit)
}

func _createAccessway(nested, level int, exit *Exit) (*Room, error) {
	t, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	ct := t.Int()
	reroll := false

	ac := &Room{
		ID:     nextRoomID(),
		Level:  level,
		Type:   "Accessway",
		Traits: []Trait{BasicTrait{name: "Table Result", roll: ct}},
	}

	if exit != nil {
		ac.Parent = exit.Parent
	}

	exits := []*Exit{
		&Exit{ID: nextExitID(), Type: ExitNormal, Parent: ac},
	}

	if ct <= 9 {
		// connects back to previous corridor or chamber
		exits[0].Child = exit.Parent
	}

	if (ct >= 10 && ct <= 12) || (ct >= 18 && ct <= 19) {
		reroll = true
	}

	if ct == 20 {
		exits = []*Exit{
			&Exit{ID: nextExitID(), Type: ExitTBranch, Parent: ac, Child: nil},
			&Exit{ID: nextExitID(), Type: ExitTBranch, Parent: ac, Child: nil},
		}
	}

	ac.Exits = exits

	if reroll {
		if len(ac.Exits) <= 0 {
			ac.Exits = []*Exit{&Exit{ID: nextExitID(), Type: ExitNormal, Parent: ac}}
		}

		ac2, err := _createAccessway(nested+1, level+1, ac.Exits[0])
		if err != nil {
			return nil, err
		}
		ac = ac2
	}

	return ac, nil
}
