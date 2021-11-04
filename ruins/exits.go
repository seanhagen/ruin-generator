package ruins

import "github.com/justinian/dice"

/**
 * File: exits.go
 * Date: 2021-11-03 16:55:43
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// ExitType ...
type ExitType int

const (
	// ExitNormal is for exits that are just regular doorways/hatches/whatever
	ExitNormal ExitType = iota
	// ExitTrapped means the exit _looks_ like a sealed one, but is trapped
	ExitTrapped
	// ExitSealed are for exits that are sealed but can be opened
	ExitSealed
	// ExitHidden are for exits that are not immediately obvious as exits ( not necessarily 'secret' )
	ExitHidden
	// ExitRupture is for exits that are replaced by the rupture
	ExitRupture
)

// Exit ...
type Exit struct {
	Type   ExitType
	Parent *Room
	Child  *Room
}

var exitList = []struct {
	min     int
	max     int
	normal  int
	trapped int
	sealed  int
	hidden  int
	reroll  bool
}{
	{1, 4, 0, 0, 0, 0, false},
	{5, 12, 1, 0, 0, 0, false},
	{13, 14, 0, 0, 1, 0, false},
	{15, 15, 2, 0, 0, 0, false},
	{16, 16, 1, 0, 1, 0, false},
	{17, 17, 2, 0, 1, 0, false},
	{18, 18, 0, 0, 2, 0, false},
	{19, 19, 0, 1, 0, 0, true},
	{20, 20, 0, 0, 0, 1, true},
}

// addExits ...
func addExits(ro *Room) error {
	res, _, err := dice.Roll("1d20")
	if err != nil {
		return err
	}

	normal := 0
	sealed := 0
	trapped := 0
	hidden := 0

	r := res.Int()
	for _, e := range exitList {
		if e.min <= r && e.max >= r {
			normal += e.normal
			sealed += e.sealed
			trapped += e.trapped
			hidden += e.hidden
			if e.reroll {
				if err := addExits(ro); err != nil {
					return err
				}
			}
		}
	}

	exits := []Exit{}

	for i := 0; i < normal; i++ {
		exits = append(exits, Exit{ExitNormal, ro, nil})
	}
	for i := 0; i < sealed; i++ {
		exits = append(exits, Exit{ExitSealed, ro, nil})
	}
	for i := 0; i < trapped; i++ {
		exits = append(exits, Exit{ExitTrapped, ro, nil})
	}
	for i := 0; i < hidden; i++ {
		exits = append(exits, Exit{ExitHidden, ro, nil})
	}

	ro.Exits = exits
	return nil
}
