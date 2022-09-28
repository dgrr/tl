package iter

import "github.com/dgrr/tl"

type iterFlatten[T any] struct {
	inner   tl.Iter[[]T]
	current T
	vs      []T
}

func (iter *iterFlatten[T]) Next() bool {
	if len(iter.vs) == 0 {
		if !iter.inner.Next() {
			return false
		}

		iter.vs = iter.inner.Get()
	}

	iter.current = iter.vs[0]
	iter.vs = iter.vs[1:]

	return true
}

func (iter *iterFlatten[T]) Get() T {
	return iter.current
}

func (iter *iterFlatten[T]) GetPtr() *T {
	return &iter.current
}

func Flatten[T any](inner tl.Iter[[]T]) tl.Iter[T] {
	return &iterFlatten[T]{
		inner: inner,
	}
}
