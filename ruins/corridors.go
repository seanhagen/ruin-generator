package ruins

import "github.com/justinian/dice"

/**
 * File: corridors.go
 * Date: 2021-11-03 16:11:45
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

var corridorTable = rollingTable{
	diceRoll: "1d20",
	results: []feature{
		{1, 7},
	},
}

func createCorridor(level int, parent *Exit) (*Room, error) {
	t, _, err := dice.Roll("1d20")

	ct := t.Int()

	reroll := false

}
