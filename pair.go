package tl

type Pair[T, U any] struct {
	first  T
	second U
}

func MakePair[T, U any](t T, u U) Pair[T, U] {
	return Pair[T, U]{
		first:  t,
		second: u,
	}
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return MakePair(p.second, p.first)
}

func (p Pair[T, U]) First() T {
	return p.first
}

func (p Pair[T, U]) Second() U {
	return p.second
}

func (p Pair[T, U]) Both() (T, U) {
	return p.first, p.second
}
