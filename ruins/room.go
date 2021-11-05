package ruins

import (
	"fmt"
	"strings"
)

/**
 * File: room.go
 * Date: 2021-11-03 16:41:28
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// Room ...
type Room struct {
	ID     string
	Level  int
	Type   string
	Traits []Trait
	Exits  []*Exit
	Parent *Room
}

// String  ...
func (r Room) String() string {
	traits := ""
	// exits := ""

	indent := "\n"
	for i := 0; i <= r.Level; i++ {
		indent = fmt.Sprintf("%v*", indent)
	}

	for _, t := range r.Traits {
		if traits == "" {
			traits = fmt.Sprintf(" - Traits => %v", t.String())
		} else {
			traits = fmt.Sprintf("%v, %v", traits, t.String())
		}
	}

	msg := fmt.Sprintf("%v Room %v (D: %v, E: %v): %v", indent, r.ID, r.Level, len(r.Exits), r.Type)
	msg = fmt.Sprintf("%v\n%v\n", msg, traits)
	// msg = fmt.Sprintf("%v - Exits:\n%v", msg, exits)

	if len(r.Exits) > 0 {
		msg = fmt.Sprintf("%v - Exit List:\n", msg)
		msg = fmt.Sprintf("%v---------##%v##---------\n", msg, r.ID)

		list := []string{}
		for _, e := range r.Exits {
			list = append(list, fmt.Sprintf(" -- Room %v", e))
		}

		msg = fmt.Sprintf("%v%v", msg, strings.Join(list, "\n"))
		msg = fmt.Sprintf("%v\n+++++++++++%v+++++++++++\n", msg, r.ID)

		for _, e := range r.Exits {
			if e.Child == nil {
				msg = fmt.Sprintf("%v\n%v* nil child", msg, indent)
			} else {
				msg = fmt.Sprintf("%v%v", msg, e.Child)
			}
		}
	} else {
		msg = fmt.Sprintf("%v - No Exits", msg)
	}

	return msg
}
