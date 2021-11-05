package ruins

import (
	"math/rand"

	"github.com/justinian/dice"
)

/**
 * File: vault.go
 * Date: 2021-11-04 15:31:56
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// vaultFeature ...
type vaultFeature struct {
	baseFeature
}

// Apply ...
func (vf vaultFeature) Apply(exit *Exit) (*Room, error) {
	return createVault(exit.Parent.Level+1, exit)
}

// createVault  ...
func createVault(level int, exit *Exit) (*Room, error) {
	ro := &Room{
		ID:    nextRoomID(),
		Level: level,
		Type:  "Relic Chamber",
	}
	traits := getSizeAndShape()

	if rand.Intn(2) == 1 {
		traits = append(traits, StringTrait{name: "Security", value: "Level 7"})
	}

	c, _, err := dice.Roll("1d20")
	if err != nil {
		return nil, err
	}

	traits = append(traits, BasicTrait{name: "Vault Contents", roll: c.Int()})
	ro.Traits = traits

	return ro, nil
}
