package iter

import (
	"testing"
)

func TestSlice(t *testing.T) {
	res := Nth(
		Filter(
			Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			func(n int) bool {
				return n > 2
			},
		), 0,
	)

	if !res.Next() && res.Get() != 3 {
		t.Fatalf("differs: %v <> %v", res.Get(), 3)
	}
}
