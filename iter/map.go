package iter

import "github.com/dgrr/tl"

type iterMap[T, V any] struct {
	current V
	inner   tl.Iter[T]
	conv    func(T) V
}

func (iter *iterMap[T, V]) Next() bool {
	if !iter.inner.Next() {
		return false
	}

	iter.current = iter.conv(iter.inner.Get())

	return true
}

func (iter *iterMap[T, V]) Get() V {
	return iter.current
}

func (iter *iterMap[T, V]) GetPtr() *V {
	return &iter.current
}

func Map[T, V any](inner tl.Iter[T], conv func(T) V) tl.Iter[V] {
	return &iterMap[T, V]{
		inner: inner,
		conv:  conv,
	}
}
