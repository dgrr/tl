package tl

// Contains returns where `e` is present in `vs`.
func Contains[T comparable](vs []T, e T) bool {
	for i := range vs {
		if vs[i] == e {
			return true
		}
	}

	return false
}

// SearchFn iterates over `vs` comparing the values using `cmpFn`.
//
// Returns the index to the element if found, -1 otherwise.
func SearchFn[T any](vs []T, cmpFn CompareFunc[T]) int {
	for i := range vs {
		if cmpFn(vs[i]) {
			return i
		}
	}

	return -1
}

// ContainsFn returns where `vs` contains an element using `cmpFn`.
func ContainsFn[T any](vs []T, cmpFn CompareFunc[T]) bool {
	for i := range vs {
		if cmpFn(vs[i]) {
			return true
		}
	}

	return false
}

// Map maps the values of `set` using `fn`.
func Map[T, E any](set []T, fn func(T) E) []E {
	r := make([]E, len(set))
	for i := range set {
		r[i] = fn(set[i])
	}

	return r
}

// Filter filters a set of values using `cmpFn`.
func Filter[T any](set []T, cmpFn CompareFunc[T]) []T {
	r := make([]T, 0)
	for i := range set {
		if cmpFn(set[i]) {
			r = append(r, set[i])
		}
	}

	return r
}

// FilterInPlace filters the set in-place, meaning that it will not create a new slice like in Filter.
func FilterInPlace[T any](set []T, cmpFn CompareFunc[T]) []T {
	for i := 0; i < len(set); i++ {
		if !cmpFn(set[i]) {
			set = append(set[:i], set[i+1:]...)
			i--
		}
	}

	return set
}

// Delete removes an element from a set. Notice that this function only removes the first occurrence.
func Delete[T comparable](set []T, value T) []T {
	for i := 0; i < len(set); i++ {
		if set[i] == value {
			if i == len(set)-1 {
				set = set[:i]
			} else {
				set = append(set[:i], set[i+1:]...)
			}

			break
		}
	}

	return set
}

// DeleteAll deletes all occurrences of a matching value.
func DeleteAll[T comparable](set []T, value T) []T {
	for i := 0; i < len(set); i++ {
		if set[i] == value {
			if i == len(set)-1 {
				set = set[:i]
			} else {
				set = append(set[:i], set[i+1:]...)
			}

			i--
		}
	}

	return set
}

// Join returns a slice containing the elements that are present in `a` and `b`.
func Join[T comparable](a, b []T) []T {
	return JoinFn(a, b, func(a, b T) bool {
		return a == b
	})
}

// JoinFn performs an Join operation for non-comparable types.
//
// See Join for more.
func JoinFn[T any](a, b []T, cmpFn func(a, b T) bool) (r []T) {
	for i := range a {
		if ContainsFn(b, func(e T) bool {
			return cmpFn(e, a[i])
		}) {
			r = append(r, a[i])
		}
	}

	return r
}

// AntiJoin returns a slice containing the elements that are present in `a` but not in `b`.
func AntiJoin[T comparable](a, b []T) []T {
	return AntiJoinFn(a, b, func(a, b T) bool {
		return a == b
	})
}

// AntiJoinFn performs an AntiJoin operation for non-comparable types.
//
// See AntiJoin for more.
func AntiJoinFn[T any](a, b []T, cmpFn func(a, b T) bool) (r []T) {
	for i := range a {
		if !ContainsFn(b, func(e T) bool {
			return cmpFn(e, a[i])
		}) {
			r = append(r, a[i])
		}
	}

	return r
}

// Merge merges multiple slices into one avoiding duplicates.
//
// Returns a slice containing all the elements in a, and more... without duplicates.
func Merge[T comparable](a []T, more ...[]T) (r []T) {
	return MergeFn(func(a, b T) bool {
		return a == b
	}, a, more...)
}

// MergeFn performs a Merge operation for non-comparable types.
//
// See Merge for more.
func MergeFn[T any](cmpFn func(a, b T) bool, a []T, more ...[]T) (r []T) {
	r = append(r, a...)
	for _, b := range more {
		for i := range b {
			if !ContainsFn(a, func(e T) bool {
				return cmpFn(e, b[i])
			}) {
				r = append(r, b[i])
			}
		}
	}

	return
}
