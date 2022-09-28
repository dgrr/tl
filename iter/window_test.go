package iter

import (
	"reflect"
	"testing"
)

func TestWindow(t *testing.T) {
	res := ToSlice(WindowCopy(Slice([]int{1, 2, 3, 4, 5}), 2))

	if !reflect.DeepEqual(res, [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}) {
		t.Fatalf("Unexpected: %v", res)
	}
}
