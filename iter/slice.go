package iter

import "github.com/dgrr/tl"

type iterSlice[T any] struct {
	vs      []T
	current *T
}

func (iter *iterSlice[T]) Next() bool {
	next := len(iter.vs) != 0
	if !next {
		iter.current = nil
	} else {
		iter.current = &iter.vs[0]
		iter.vs = iter.vs[1:]
	}

	return next
}

func (iter *iterSlice[T]) Get() T {
	return *iter.current
}

func (iter *iterSlice[T]) GetPtr() *T {
	return iter.current
}

func Slice[T any](vs []T) tl.Iter[T] {
	return &iterSlice[T]{
		vs: vs,
	}
}

func AppendTo[T any](vs []T, iter tl.Iter[T]) []T {
	for iter.Next() {
		vs = append(vs, iter.Get())
	}

	return vs
}

func ToSlice[T any](iter tl.Iter[T]) []T {
	newSlice := make([]T, 0)

	for iter.Next() {
		newSlice = append(newSlice, iter.Get())
	}

	return newSlice
}
