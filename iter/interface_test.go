package iter

import (
	"github.com/dgrr/tl"
	"testing"
)

func TestAdvance(t *testing.T) {
	it := Slice([]int{1, 2, 3})
	it.Next()

	tl.Advance(it, 1)

	if !it.Next() || it.Get() != 3 {
		t.Fatalf("unexpected value: %d", it.Get())
	}
}
