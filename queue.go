package tl

type element[T any] struct {
	data T
	next *element[T]
}

// Queue defines a queue data structure.
type Queue[T any] struct {
	first *element[T]
	last  *element[T]
}

// Reset resets the queue.
func (q *Queue[T]) Reset() {
	q.first = nil
	q.last = nil
}

// Front returns the first element if any.
func (q *Queue[T]) Front() (v Optional[T]) {
	if q.first != nil {
		v.Set(q.first.data)
	}

	return
}

// PushFront pushes `data` to the beginning of the queue.
func (q *Queue[T]) PushFront(data T) {
	if q.first == nil {
		q.first = &element[T]{
			data: data,
		}
		q.last = q.first
	} else {
		e := &element[T]{
			data: data,
			next: q.first,
		}
		q.first = e
	}
}

// PushBack pushes to the back of the queue.
func (q *Queue[T]) PushBack(data T) {
	if q.first == nil {
		q.first = &element[T]{
			data: data,
		}
		q.last = q.first
	} else {
		q.last.next = &element[T]{
			data: data,
		}
		q.last = q.last.next
	}
}

// Pop pops from the beginning of the queue.
func (q *Queue[T]) Pop() (v Optional[T]) {
	if q.first != nil {
		v.Set(q.first.data)
		q.first = q.first.next
	}

	return
}
