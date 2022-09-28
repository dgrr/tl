package tl

// Pair defines a pair of values.
type Pair[T, U any] struct {
	first  T
	second U
}

// MakePair returns a Pair instance using (t, u).
func MakePair[T, U any](t T, u U) Pair[T, U] {
	return Pair[T, U]{
		first:  t,
		second: u,
	}
}

// Swap returns a new Pair swapping the place of the values.
func (p Pair[T, U]) Swap() Pair[U, T] {
	return MakePair(p.second, p.first)
}

// First returns the first value.
func (p Pair[T, U]) First() T {
	return p.first
}

// Second returns the second value.
func (p Pair[T, U]) Second() U {
	return p.second
}

// Both returns both values at the same time.
func (p Pair[T, U]) Both() (T, U) {
	return p.first, p.second
}
