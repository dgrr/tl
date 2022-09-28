package tl

type Iter[T any] interface {
	Next() bool
	Get() T
	GetPtr() *T
}

type IterDrop[T any] interface {
	Iter[T]

	Drop()
}

type IterBidir[T any] interface {
	Iter[T]

	Back() bool
}

type IterDropBidir[T any] interface {
	IterDrop[T]
	IterBidir[T]
}

func Advance[T any](iter Iter[T], count int) {
	for i := 0; i < count && iter.Next(); i++ {
	}
}
