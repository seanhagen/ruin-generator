package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/ruin-generator/ruins"
)

/**
 * File: main.go
 * Date: 2021-11-03 15:16:33
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func main() {
	ent := ruins.GetEntrance()

	// feat, lvl, err := ruins.RollMainFeature(0)
	// spew.Dump(feat, lvl, err)

	spew.Dump(ent)
}
