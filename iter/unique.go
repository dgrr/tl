package iter

import "github.com/dgrr/tl"

type iterUnique[T any] struct {
	inner tl.Iter[T]
	eq    func(a, b T) bool
	prev  []T
}

func (iter *iterUnique[T]) Next() bool {
	for iter.inner.Next() {
		next := iter.inner.Get()
		if !tl.ContainsFn(iter.prev, func(a T) bool {
			return iter.eq(next, a)
		}) {
			iter.prev = append(iter.prev, next)
			return true
		}
	}

	return false
}

func (iter *iterUnique[T]) Get() T {
	return iter.inner.Get()
}

func (iter *iterUnique[T]) GetPtr() *T {
	return iter.inner.GetPtr()
}

func UniqueFn[T any](inner tl.Iter[T], eq func(a, b T) bool) tl.Iter[T] {
	return &iterUnique[T]{
		inner: inner,
		eq:    eq,
	}
}

func Unique[T comparable](inner tl.Iter[T]) tl.Iter[T] {
	return UniqueFn(inner, func(a, b T) bool {
		return a == b
	})
}

func Get[T any](iter tl.Iter[T]) T {
	for iter.Next() {
	}

	return iter.Get()
}
