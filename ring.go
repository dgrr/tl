package tl

import (
	"sync/atomic"
)

// Ring is a Multiple-Producer Multiple-Consumer (MPMC) circular buffer data structure.
type Ring[T any] struct {
	head   uint32
	tail   uint32
	mask   uint32
	index  []uint32
	values []T
}

func roundNearest(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

// NewRing returns a ring of type T with the size rounded up to the nearest power of 2.
func NewRing[T any](size int) *Ring[T] {
	size = roundNearest(size)

	ring := &Ring[T]{
		values: make([]T, size),
		index:  make([]uint32, size),
		mask:   uint32(size - 1),
	}
	for i := range ring.index {
		ring.index[i] = uint32(i)
	}

	return ring
}

// Push adds a value to the ring.
//
// If the ring is full Push returns false. If the task has been pushed it will return true.
func (ring *Ring[T]) Push(value T) bool {
	for {
		head := atomic.LoadUint32(&ring.head)
		tail := atomic.LoadUint32(&ring.tail)

		if head-tail == ring.mask {
			return false
		}

		index := atomic.LoadUint32(&ring.index[head&ring.mask])
		if index == head {
			if atomic.CompareAndSwapUint32(&ring.head, head, head+1) {
				ring.values[head&ring.mask] = value

				atomic.StoreUint32(&ring.index[head&ring.mask], head+1)

				return true
			}
		}
	}
}

// Pop takes a value from the ring. Returns the (value, true) or (value, false)
func (ring *Ring[T]) Pop() (value T, ok bool) {
	for {
		head := ring.head
		tail := atomic.LoadUint32(&ring.tail)

		if tail == head { // nothing in the queue
			return
		}

		index := atomic.LoadUint32(&ring.index[tail&ring.mask])
		if index == tail+1 {
			// advance the tail value
			if atomic.CompareAndSwapUint32(&ring.tail, tail, tail+1) {
				value = ring.values[tail&ring.mask]
				atomic.StoreUint32(&ring.index[tail&ring.mask], tail+ring.mask+1)
				return value, true
			}
		}
	}
}
