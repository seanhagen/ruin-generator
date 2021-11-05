package ruins

import (
	"math/rand"

	"github.com/justinian/dice"
)

/**
 * File: interstitial.go
 * Date: 2021-11-04 14:42:46
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// interstitialFeature ...
type interstitialFeature struct {
	baseFeature
}

// Apply ...
func (inf interstitialFeature) Apply(exit *Exit) (*Room, error) {
	return createInterstitial(exit.Parent.Level+1, exit)
}

// createInterstitial ...
func createInterstitial(level int, parent *Exit) (*Room, error) {
	width := 1000
	depth := 500

	if rand.Intn(2) == 1 {
		if rand.Intn(2) == 0 {
			width += rand.Intn(501)
		} else {
			width -= rand.Intn(501)
		}
	}

	if rand.Intn(2) == 1 {
		if rand.Intn(2) == 0 {
			depth += rand.Intn(251)
		} else {
			depth -= rand.Intn(251)
		}
	}

	res, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}
	ro := &Room{
		ID:    nextRoomID(),
		Level: level,
		Type:  "Interstitial cavity",
		Traits: []Trait{
			BasicTrait{"Width", width},
			BasicTrait{"Depth", depth},
			BasicTrait{"Interstitial Cavity", res.Int()},
		},
	}

	err = addInterstitialExits(ro)
	return ro, err
}
