package tl

import (
	"testing"

	"github.com/dgrr/tl"
	"github.com/dgrr/tl/iter"
)

func TestList(t *testing.T) {
	var list tl.List[int]

	type expectedValue struct {
		fn       func(int) *tl.ListElement[int]
		forward  bool
		expected []int
	}

	values := iter.ToSlice(iter.Range(0, 10, 1))

	es := []expectedValue{
		{
			fn:       list.PushBack,
			forward:  true,
			expected: iter.ToSlice(iter.Range(0, 10, 1)),
		}, {
			fn:       list.PushFront,
			forward:  true,
			expected: iter.ToSlice(iter.Range(9, -1, -1)),
		}, {
			fn:       list.PushBack,
			forward:  false,
			expected: iter.ToSlice(iter.Range(9, -1, -1)),
		}, {
			fn:       list.PushFront,
			forward:  false,
			expected: iter.ToSlice(iter.Range(0, 10, 1)),
		},
	}

	for idx, e := range es {
		for _, v := range values {
			e.fn(v)
		}

		iter := list.Iter()
		next := iter.Next

		if !e.forward {
			next = iter.Back
		}

		for i := 0; i < len(e.expected); i++ {
			next()

			if e.expected[i] != iter.Get() {
				t.Fatalf("unexpected value on %d,%d: %d <> %d", idx, i, e.expected[i], iter.Get())
			}
		}

		if next() {
			t.Fatal("unexpected")
		}

		list.Reset()
	}
}

func TestListPop(t *testing.T) {
	var list tl.List[int]

	type expectedValue struct {
		fn       func(int)
		iter     func() tl.IterDropBidir[int]
		expected []int
	}

	iter := list.Iter()
	if iter.Next() {
		t.Fatal("unexpected")
	}

	{
		list.PushBack(2)
		list.PushFront(4)
		list.PopFront()
		list.PopBack()
		list.PushBack(3)
		list.PopFront()
		list.PushFront(2)
		list.PushFront(4)
		list.PopBack()
		list.PopFront()

		iter = list.Iter()
		if iter.Next() {
			t.Fatal("unexpected")
		}
	}
	{
		list.PushFront(2)
		list.PopBack()

		iter := list.Iter()
		if iter.Next() {
			t.Fatal("unexpected")
		}
	}
}

func TestIterDrop(t *testing.T) {
	var list tl.List[int]

	values := iter.ToSlice(iter.Range(1, 11, 1))

	for _, v := range values {
		list.PushBack(v)
	}

	it := list.Iter()

	it.Next()
	values = values[1:]

	for it.Next() {
		if it.Get() != values[0] {
			t.Fatalf("got %d, expected %d", it.Get(), values[0])
		}

		it.Drop()

		values = values[1:]
	}

	if len(values) != 0 {
		t.Fatalf("unexpected: %d", len(values))
	}

	if it.Next() {
		t.Fatalf("unexpected: %d", it.Get())
	}

	if list.Size() != 1 {
		t.Fatalf("Unexpected: %d", list.Size())
	}
}

func TestPushDrop(t *testing.T) {
	var list tl.List[int]

	values := iter.ToSlice(iter.Range(1, 11, 1))

	vars := make([]*tl.ListElement[int], 0)

	for _, v := range values {
		vars = append(vars, list.PushBack(v))
	}

	for len(vars) != 0 {
		if vars[0].Get() != values[0] {
			t.Fatalf("got %d, expected %d", vars[0].Get(), values[0])
		}

		vars[0].Drop()

		values = values[1:]
		vars = vars[1:]
	}

	if len(values) != 0 {
		t.Fatalf("unexpected: %d", len(values))
	}

	if list.Iter().Next() {
		t.Fatal("unexpected")
	}

	if list.Size() != 0 {
		t.Fatalf("Unexpected: %d", list.Size())
	}
}

func TestIterDropReverse(t *testing.T) {
	var list tl.List[int]

	values := iter.ToSlice(iter.Range(1, 11, 1))

	for _, v := range values {
		list.PushBack(v)
	}

	values = iter.ToSlice(iter.Range(10, 0, -1))

	it := list.Iter()

	it.Back()
	values = values[1:]

	for it.Back() {
		if it.Get() != values[0] {
			t.Fatalf("got %d, expected %d", it.Get(), values[0])
		}

		it.Drop()

		values = values[1:]
	}

	if len(values) != 0 {
		t.Fatalf("unexpected: %d", len(values))
	}

	if it.Back() {
		t.Fatalf("unexpected: %d", it.Get())
	}

	if list.Size() != 1 {
		t.Fatalf("Unexpected: %d", list.Size())
	}
}
