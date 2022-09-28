package iter

import (
	"testing"
)

func TestNth(t *testing.T) {
	res := Nth(Window(Slice([]int{1, 2, 3, 4, 5}), 2), 8)
	if res.Next() {
		t.Fatal("unexpected res")
	}
}
