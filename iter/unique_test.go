package iter

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	res := ToSlice(
		Unique(
			Slice([]int{1, 1, 4, 8, 7, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		),
	)

	expect := []int{1, 4, 8, 7, 11, 2, 3, 5, 6, 9, 10}

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("differs: %v <> %v", res, expect)
	}
}
