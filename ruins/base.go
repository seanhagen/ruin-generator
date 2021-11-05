package ruins

import (
	"strings"
)

/**
 * File: base.go
 * Date: 2021-11-04 14:11:03
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// MaxDepth ...
const MaxDepth int = 10

var currentExit int = 1
var currentRoom int = 1

func roomToID(i int) string {
	if i <= 0 {
		return "-"
	}

	j := (i - 1) % 26
	l := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := roomToID((i - j) / 26)
	final := s + string(l[j])

	if final == "-" {
		return "-"
	}

	return strings.Replace(final, "-", "", -1)
}

func nextRoomID() string {
	old := currentRoom
	currentRoom++
	return roomToID(old)
}

func nextExitID() int {
	old := currentExit
	currentExit++
	return old
}
