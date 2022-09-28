package iter

import "github.com/dgrr/tl"

func Index[T any](iter tl.Iter[T], cmpFn tl.CompareFunc[T]) int {
	i := 0
	for ; iter.Next(); i++ {
		if cmpFn(iter.Get()) {
			return i
		}
	}

	return -1
}

func Search[T any](iter tl.Iter[T], cmpFn tl.CompareFunc[T]) tl.Iter[T] {
	for iter.Next() {
		if cmpFn(iter.Get()) {
			return iter
		}
	}

	return nil
}
