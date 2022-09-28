package tl

import (
	"reflect"
	"sort"
	"testing"

	"github.com/dgrr/tl"
)

func TestAntiJoin(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{0, 2, 3, 9}
	r := []int{1, 4, 5, 6}

	got := tl.AntiJoin(a, b)

	if !reflect.DeepEqual(got, r) {
		t.Fatalf("unexpected: %v <> %v", r, got)
	}

	got = tl.AntiJoin(b, []int{})

	if !reflect.DeepEqual(got, b) {
		t.Fatalf("unexpected: %v <> %v", r, got)
	}
}

func TestJoin(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{0, 2, 3, 9}
	r := []int{2, 3}

	got := tl.Join(a, b)

	if !reflect.DeepEqual(got, r) {
		t.Fatalf("unexpected: %v <> %v", r, got)
	}
}

func TestMerge(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{0, 2, 3, 9}
	r := []int{0, 1, 2, 3, 4, 5, 6, 9}

	got := tl.Merge(a, b)
	sort.Ints(got)

	if !reflect.DeepEqual(got, r) {
		t.Fatalf("unexpected: %v <> %v", r, got)
	}
}
