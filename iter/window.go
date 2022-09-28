package iter

import "github.com/dgrr/tl"

type iterWindow[T any] struct {
	size  int
	win   []T
	inner tl.Iter[T]
}

func (iter *iterWindow[T]) Next() bool {
	if len(iter.win) != 0 {
		iter.win = append(iter.win[:0], iter.win[1:]...)
	}

	for len(iter.win) < iter.size {
		if !iter.inner.Next() {
			return false
		}

		iter.win = append(iter.win, iter.inner.Get())
	}

	return true
}

func (iter *iterWindow[T]) Get() []T {
	return iter.win
}

func (iter *iterWindow[T]) GetPtr() *[]T {
	return &iter.win
}

// Window returns an iterator containing the last `n` values.
func Window[T any](inner tl.Iter[T], n int) tl.Iter[[]T] {
	return &iterWindow[T]{
		inner: inner,
		win:   make([]T, 0, n),
		size:  n,
	}
}
