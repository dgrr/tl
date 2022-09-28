package tl

type Vec[T any] []T

func MakeVec[T any](elmnts ...T) Vec[T] {
	vc := Vec[T]{}
	vc.Append(elmnts...)

	return vc
}

func MakeVecSize[T any](size, capacity int) Vec[T] {
	return (Vec[T])(make([]T, size, capacity))
}

func (vc Vec[T]) Get(i int) T {
	return vc[i]
}

func (vc *Vec[T]) Resize(n int) {
	vc.Reserve(n)
	*vc = (*vc)[:n]
}

func (vc *Vec[T]) Reserve(n int) {
	if nSize := n - cap(*vc); nSize > 0 {
		*vc = append((*vc)[:cap(*vc)], make([]T, nSize)...)
	}
}

func (vc *Vec[T]) Append(elmnts ...T) {
	*vc = append(*vc, elmnts...)
}

func (vc *Vec[T]) Push(elmnts ...T) {
	*vc = append((*vc)[:len(elmnts)], *vc...)
	copy(*vc, elmnts)
}

func (vc Vec[T]) Front() (opt OptionalPtr[T]) {
	if len(vc) != 0 {
		opt.Set(&vc[0])
	}
	return opt
}

func (vc Vec[T]) Back() (opt OptionalPtr[T]) {
	if len(vc) != 0 {
		opt.Set(&vc[len(vc)-1])
	}

	return
}

func (vc *Vec[T]) PopBack() (opt OptionalPtr[T]) {
	if len(*vc) != 0 {
		opt.Set(&(*vc)[len(*vc)-1])
		*vc = (*vc)[:len(*vc)-1]
	}

	return
}

func (vc *Vec[T]) PopFront() (opt OptionalPtr[T]) {
	if len(*vc) != 0 {
		opt.Set(&(*vc)[0])
		*vc = append((*vc)[:0], (*vc)[1:]...)
	}

	return
}

func (vc Vec[T]) Len() int {
	return len(vc)
}

func (vc Vec[T]) Cap() int {
	return cap(vc)
}

type CompareFunc[T any] func(T) bool

func (vc *Vec[T]) Filter(cmpFn CompareFunc[T]) (val T, erased bool) {
	for i := 0; i < len(*vc); i++ {
		if cmpFn((*vc)[i]) {
			vc.DelByIndex(i)
			i--
		}
	}

	return
}

func (vc *Vec[T]) DelByIndex(i int) (val T, erased bool) {
	if vc.Len() <= i {
		return val, false
	}

	val = (*vc)[i]
	*vc = append((*vc)[:i], (*vc)[i+1:]...)

	return val, true
}

func (vc Vec[T]) Index(cmpFn CompareFunc[T]) int {
	for i := range vc {
		if cmpFn(vc[i]) {
			return i
		}
	}

	return -1
}

func (vc Vec[T]) Contains(cmpFn CompareFunc[T]) bool {
	return vc.Index(cmpFn) >= 0
}

func (vc Vec[T]) Search(cmpFn CompareFunc[T]) (v T, ok bool) {
	for i := range vc {
		if cmpFn(vc[i]) {
			return vc[i], true
		}
	}

	return
}
