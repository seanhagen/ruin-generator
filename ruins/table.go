package ruins

import (
	"fmt"

	"github.com/justinian/dice"
)

/**
 * File: table.go
 * Date: 2021-11-03 17:08:01
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

type result interface {
	// IsResult returns true if the number passed in would return this result on the table
	IsResult(int) bool
	// Name returns the name of the table result
	Name() string
	Apply(*Exit) (*Room, error)
}

// baseFeature ...
type baseFeature struct {
	min, max int
	name     string
}

// IsResult  ...
func (bf baseFeature) IsResult(i int) bool {
	return i >= bf.min && i <= bf.max
}

// Name  ...
func (bf baseFeature) Name() string {
	return bf.name
}

type rollingTable struct {
	diceRoll string
	results  []result
}

// getResult ...
func (rt rollingTable) getResult() (result, error) {
	r, _, err := dice.Roll(rt.diceRoll)
	if err != nil {
		return nil, err
	}

	for _, f := range rt.results {
		if f.IsResult(r.Int()) {
			return f, nil
		}
	}
	return nil, fmt.Errorf("%v is not possible when rolling %v", r, rt.diceRoll)
}
