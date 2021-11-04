package ruins

import (
	"fmt"
)

/**
 * File: room.go
 * Date: 2021-11-03 16:41:28
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// Room ...
type Room struct {
	Level  int
	Type   string
	Traits []Trait
	Exits  []Exit
}

func CreateRoom(parent *Exit) (*Room, error) {

	return nil, fmt.Errorf("not yet")
}
