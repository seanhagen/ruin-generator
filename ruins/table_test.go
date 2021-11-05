package ruins

import (
	"fmt"
	"testing"
)

/**
 * File: table_test.go
 * Date: 2021-11-04 14:28:12
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

func TestRuins_BaseFeatureIsResult(t *testing.T) {
	tests := []struct {
		a, b, c int
		t       bool
	}{
		{1, 2, 1, true},
		{1, 1, 2, false},
		{59, 70, 63, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test%v", i), func(t *testing.T) {
			b := baseFeature{min: tt.a, max: tt.b}
			res := b.IsResult(tt.c)
			if res != tt.t {
				t.Errorf("wrong result, for range of [%v:%v], %v should be %v, got %v", tt.a, tt.b, tt.c, tt.t, res)
			}
		})
	}
}
