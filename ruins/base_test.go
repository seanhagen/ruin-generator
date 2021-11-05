package ruins

import (
	"fmt"
	"testing"
)

/**
 * File: base_test.go
 * Date: 2021-11-04 16:37:19
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func TestRuins_Base(t *testing.T) {
	tests := []struct {
		in int
		ex string
	}{
		{-1, "-"},
		{0, "-"},
		{1, "A"},
		{2, "B"},
		{10, "J"},
		{20, "T"},
		{26, "Z"},
		{27, "AA"},
		{52, "AZ"},
		{53, "BA"},
		{60, "BH"},
		{78, "BZ"},
		{79, "CA"},
		{676, "YZ"},
		{677, "ZA"},
		{702, "ZZ"},
		{703, "AAA"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test%v", i), func(t *testing.T) {
			out := roomToID(tt.in)
			if out != tt.ex {
				t.Errorf("Wrong output for %v, expected '%v' got '%v'", tt.in, tt.ex, out)
			}
		})
	}
}
