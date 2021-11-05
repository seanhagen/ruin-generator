package ruins

import (
	"fmt"

	"github.com/justinian/dice"
)

/**
 * File: relic.go
 * Date: 2021-11-04 15:44:47
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// relicFeature ...
type relicFeature struct {
	baseFeature
}

// Apply ...
func (lf relicFeature) Apply(exit *Exit) (*Room, error) {
	return createRelic(exit.Parent.Level+1, exit)
}

type vaultTrait struct {
	BasicTrait
	door string
}

// String ...
func (vt vaultTrait) String() string {
	return fmt.Sprintf("Vault Door (L6)")
}

// createRelic ...
func createRelic(level int, exit *Exit) (*Room, error) {
	ro := &Room{
		ID:    nextRoomID(),
		Level: level,
		Type:  "Relic Chamber",
	}
	traits := getSizeAndShape()

	t, _, err := dice.Roll("1d100")
	if err != nil {
		return nil, err
	}

	if t.Int() >= 21 {
		err = addCreature(ro)
		if err != nil {
			return nil, err
		}
	}

	if t.Int() >= 60 {
		traits = append(traits, vaultTrait{})
	}

	an, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	traits = append(traits, BasicTrait{name: "Relic Anatomy", roll: an.Int()})

	qu, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}
	traits = append(traits, BasicTrait{name: "Relic Quality", roll: qu.Int()})

	ro.Traits = traits

	return ro, nil
}
