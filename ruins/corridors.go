package ruins

import (
	"github.com/justinian/dice"
)

/**
 * File: corridors.go
 * Date: 2021-11-03 16:11:45
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// corridorFeature ...
type corridorFeature struct {
	baseFeature
}

// Apply ...
func (cf corridorFeature) Apply(exit *Exit) (*Room, error) {
	return createCorridor(exit.Parent.Level+1, exit)
}

func createCorridor(level int, exit *Exit) (*Room, error) {
	return _createCorridor(0, level, exit)
}

func _createCorridor(nested, level int, exit *Exit) (*Room, error) {
	if nested > 5 {
		// short-circuit for unlimited corridors
		return createChamber(level, exit)
	}

	t, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	ct := t.Int()
	reroll := false

	if ct <= 7 {
		reroll = true
	}

	if ct == 17 {
		if exit.Parent != nil && (exit.Parent.Type == "Chamber" || exit.Parent.Type == "Corridor") {
			reroll = true
		}
	}

	corr := &Room{
		ID:     nextRoomID(),
		Level:  level,
		Type:   "Corridor",
		Traits: []Trait{BasicTrait{name: "Table Result", roll: ct}},
	}
	if exit != nil {
		if exit.Parent != nil {
			corr.Parent = exit.Parent
		}

		exit.Child = corr
	}

	exits := []*Exit{
		&Exit{ID: nextExitID(), Type: ExitNormal, Parent: corr},
	}

	if ct == 17 {
		exits = []*Exit{}
	}

	if ct == 19 {
		exits = []*Exit{
			&Exit{ID: nextExitID(), Type: ExitTBranch, Parent: corr, Child: nil},
			&Exit{ID: nextExitID(), Type: ExitTBranch, Parent: corr, Child: nil},
		}
	}

	if ct == 20 {
		exits = []*Exit{
			&Exit{ID: nextExitID(), Type: ExitXBranch, Parent: corr, Child: nil},
			&Exit{ID: nextExitID(), Type: ExitXBranch, Parent: corr, Child: nil},
			&Exit{ID: nextExitID(), Type: ExitXBranch, Parent: corr, Child: nil},
		}
	}

	corr.Exits = exits

	if reroll {
		if len(corr.Exits) <= 0 {
			corr.Exits = []*Exit{&Exit{ID: nextExitID(), Type: ExitNormal, Parent: corr}}
		}

		c2, err := _createCorridor(nested+1, level+1, corr.Exits[0])
		if err != nil {
			return nil, err
		}
		corr = c2
	}

	return corr, nil
}
