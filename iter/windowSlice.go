package iter

import "github.com/dgrr/tl"

type iterWindowSlice[T any] struct {
	size int
	win  []T
	vs   []T
}

func (iter *iterWindowSlice[T]) Next() bool {
	if len(iter.vs) < iter.size {
		return false
	}

	iter.win = iter.vs[:iter.size]
	iter.vs = iter.vs[1:]

	return true
}

func (iter *iterWindowSlice[T]) Get() []T {
	return iter.win
}

func (iter *iterWindowSlice[T]) GetPtr() *[]T {
	return &iter.win
}

func WindowSlice[T any](vs []T, n int) tl.Iter[[]T] {
	return &iterWindowSlice[T]{
		vs:   vs,
		size: n,
	}
}
