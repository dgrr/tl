package iter

import (
	"github.com/dgrr/tl"
	"golang.org/x/exp/constraints"
)

type iterSum[T constraints.Integer | constraints.Float] struct {
	result T
	inner  tl.Iter[T]
}

func (iter *iterSum[T]) Next() bool {
	if !iter.inner.Next() {
		return false
	}

	iter.result += iter.inner.Get()

	return true
}

func (iter *iterSum[T]) Get() T {
	return iter.result
}

func (iter *iterSum[T]) GetPtr() *T {
	return &iter.result
}

func Sum[T constraints.Integer | constraints.Float](inner tl.Iter[T]) tl.Iter[T] {
	return &iterSum[T]{
		inner: inner,
	}
}
