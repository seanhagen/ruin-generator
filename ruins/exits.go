package ruins

import (
	"fmt"

	"github.com/justinian/dice"
)

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
	// ExitTBranch are exits from a T-shaped corridor
	ExitTBranch
	// ExitXBranch are exits from a X-shaped corridor
	ExitXBranch
	// ExitShaft are exits on the side of a shaft
	ExitShaft
)

// Exit ...
type Exit struct {
	ID     int
	Type   ExitType
	Parent *Room
	Child  *Room
}

// String  ...
func (e Exit) String() string {
	et := "N"
	switch e.Type {
	case ExitTrapped:
		et = "T"
	case ExitSealed:
		et = "S"
	case ExitHidden:
		et = "H"
	case ExitRupture:
		et = "R"
	case ExitTBranch:
		et = "T"
	case ExitXBranch:
		et = "X"
	}

	if e.Child != nil {
		return fmt.Sprintf("Exit %v (T:%v L:%v) => %v", e.ID, et, e.Parent.Level, e.Child.ID)
	}

	return fmt.Sprintf("Exit %v (T:%v L:%v) => <<no chamber??>>", e.ID, et, e.Parent.Level)
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

	return _generateExits(0, res.Int(), ro)
}

func addInterstitialExits(ro *Room) error {
	res, _, err := dice.Roll("1d20")
	if err != nil {
		return err
	}

	r2, _, err := dice.Roll("1d10")
	if err != nil {
		return err
	}

	return _generateExits(0, res.Int()+r2.Int(), ro)
}

func _generateExits(nested, num int, ro *Room) error {
	if nested > 3 {
		// short-circuit for endless exits
		return nil
	}

	if ro.Level > MaxDepth {
		return nil
	}

	exits := []*Exit{}
	normal := 0
	sealed := 0
	trapped := 0
	hidden := 0

	for _, e := range exitList {
		if e.min <= num && e.max >= num {
			normal += e.normal
			sealed += e.sealed
			trapped += e.trapped
			hidden += e.hidden
			if e.reroll {
				res, _, err := dice.Roll("1d20")
				if err != nil {
					return err
				}

				if err := _generateExits(nested+1, res.Int(), ro); err != nil {
					return err
				}
			}
		}
	}

	for i := 0; i < normal; i++ {
		exits = append(exits, &Exit{nextExitID(), ExitNormal, ro, nil})
	}
	for i := 0; i < sealed; i++ {
		exits = append(exits, &Exit{nextExitID(), ExitSealed, ro, nil})
	}
	for i := 0; i < trapped; i++ {
		exits = append(exits, &Exit{nextExitID(), ExitTrapped, ro, nil})
	}
	for i := 0; i < hidden; i++ {
		exits = append(exits, &Exit{nextExitID(), ExitHidden, ro, nil})
	}

	ro.Exits = exits
	return nil
}
