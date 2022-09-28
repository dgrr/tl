package iter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	res := ToSlice(
		Filter(
			Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			func(n int) bool {
				return n < 5
			},
		),
	)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4}) {
		t.Fatalf("differs: %v <> %v", res, []int{1, 2, 3, 4})
	}
}
