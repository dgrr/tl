package tl

// Iter defines an iterator interface.
type Iter[T any] interface {
	// Next is called to iterate to the next element.
	//
	// Returns true if a next element is available, false otherwise.
	Next() bool
	// Get returns the current element.
	Get() T
	// GetPtr returns a pointer to the current element.
	GetPtr() *T
}

// IterDrop implements an iterator which current element can be dropped.
type IterDrop[T any] interface {
	Iter[T]
	// Drop removes the current element from the iterator.
	Drop()
}

// IterBidir defines a bidirectional iterator.
type IterBidir[T any] interface {
	Iter[T]
	// Back is like next but for going backwards. It moves the iterator to the previous position.
	// For some implementations that might mean going in reverse mode (not going backwards).
	Back() bool
}

// IterDropBidir merges IterBidir and IterDrop in one interface.
type IterDropBidir[T any] interface {
	IterDrop[T]
	IterBidir[T]
}

// Advance advances iter a number of `count` positions.
func Advance[T any](iter Iter[T], count int) {
	for i := 0; i < count && iter.Next(); i++ {
	}
}
