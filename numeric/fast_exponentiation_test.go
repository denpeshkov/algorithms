package numeric_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/denpeshkov/algorithms/numeric"
)

func TestFastExp(t *testing.T) {
	tests := []struct {
		a    uint64
		n    uint64
		want string
	}{
		{1, 1, "1"},
		{2, 1, "2"},
		{1, 10000, "1"},
		{2, 10, "1024"},
		{3, 5, "243"},
		{0, 0, "1"},
		{0, 10, "0"},
		{0, 4, "0"},
		{0, 1, "0"},
		{123, 14, "0x24A2E681594EED4183DD5EEE9"},
		{11, 101, "1515867357380449720253017088929860045389076432004231307646112530737048373050753844463412754472043790906011"},
		{1, 10000000, "1"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d^%d", tt.a, tt.n), func(t *testing.T) {
			got := numeric.FastExp(tt.a, tt.n)
			want, _ := new(big.Int).SetString(tt.want, 0)

			if got.Cmp(want) != 0 {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
