package iter

import "github.com/dgrr/tl"

type iterFilter[T any] struct {
	inner tl.Iter[T]
	fn    tl.CompareFunc[T]
}

func (iter *iterFilter[T]) Next() bool {
	for iter.inner.Next() {
		if iter.fn(iter.inner.Get()) {
			return true
		}
	}

	return false
}

func (iter *iterFilter[T]) Get() T {
	return iter.inner.Get()
}

func (iter *iterFilter[T]) GetPtr() *T {
	return iter.inner.GetPtr()
}

// Filter filters the values of iterators.
func Filter[T any](inner tl.Iter[T], fn tl.CompareFunc[T]) tl.Iter[T] {
	return &iterFilter[T]{
		inner: inner,
		fn:    fn,
	}
}
