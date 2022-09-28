package tl

// ListElement is an element in a List.
type ListElement[T any] struct {
	value      T
	prev, next *ListElement[T]
	list       *List[T]
}

// Get returns the value of the element.
func (e *ListElement[T]) Get() T {
	return e.value
}

// GetPtr returns a pointer to the value of the element.
func (e *ListElement[T]) GetPtr() *T {
	return &e.value
}

// Drop drops the current element from the list.
func (e *ListElement[T]) Drop() {
	if e.prev != nil {
		e.prev.next = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	}

	e.list.size--
}

// List defines a doubly linked list.
type List[T any] struct {
	root ListElement[T]
	size int
}

// Size returns the size of the linked list (number of elements inside the list).
func (list *List[T]) Size() int {
	return list.size
}

// PushBack appends an element to the back of the queue.
func (list *List[T]) PushBack(v T) *ListElement[T] {
	e := &ListElement[T]{
		value: v,
		prev:  list.root.prev,
		next:  &list.root,
		list:  list,
	}

	if e.prev == nil {
		e.prev = &list.root
	}

	if list.root.prev != nil {
		list.root.prev.next = e
	} else {
		list.root.next = e
	}

	list.root.prev = e

	list.size++

	return e
}

// PushFront pushes an element to the front of the queue.
func (list *List[T]) PushFront(v T) *ListElement[T] {
	e := &ListElement[T]{
		value: v,
		next:  list.root.next,
		prev:  &list.root,
		list:  list,
	}

	if e.next == nil {
		e.next = &list.root
	}

	if list.root.next != nil {
		list.root.next.prev = e
	} else {
		list.root.prev = e
	}

	list.root.next = e

	list.size++

	return e
}

// Front returns an optional to the first element of the queue.
func (list *List[T]) Front() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)
	}

	return
}

// Back returns an optional to the last element of the queue.
func (list *List[T]) Back() (opt OptionalPtr[T]) {
	if list.root.prev != nil {
		opt.Set(&list.root.prev.value)
	}

	return
}

// func Print[T any](list *listElement[T]) {
// 	fmt.Printf("%p - %p %p\n", list, list.prev, list.next)
// 	for next := list.next; next != nil && next != list; next = next.next {
// 		fmt.Printf("%p (%v) = %p - %p\n", next, next.value, next.prev, next.next)
// 	}
// 	println("-------")
// }

// PopFront pops the first element if any.
func (list *List[T]) PopFront() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)

		list.root.next.Drop()
	}

	return
}

// PopBack pops the last element if any.
func (list *List[T]) PopBack() (opt OptionalPtr[T]) {
	if list.root.prev != nil {
		opt.Set(&list.root.prev.value)

		list.root.prev.Drop()
	}

	return
}

// Reset resets the list.
func (list *List[T]) Reset() {
	list.root.next = nil
	list.root.prev = nil
}

// Iter returns an iterator for the List.
func (list *List[T]) Iter() IterDropBidir[T] {
	return &iterList[T]{
		root: &list.root,
		next: list.root.next,
		prev: list.root.prev,
	}
}

type iterList[T any] struct {
	root       *ListElement[T]
	current    *ListElement[T]
	prev, next *ListElement[T]
}

func (iter *iterList[T]) Drop() {
	iter.current.Drop()
}

func (iter *iterList[T]) Back() bool {
	iter.current = iter.prev

	if iter.prev != nil && iter.current != iter.root {
		iter.next = iter.prev
		iter.prev = iter.prev.prev

		return true
	}

	return false
}

func (iter *iterList[T]) Next() bool {
	iter.current = iter.next

	if iter.next != nil && iter.current != iter.root {
		iter.prev = iter.next
		iter.next = iter.next.next

		return true
	}

	return false
}

func (iter *iterList[T]) Get() T {
	return iter.current.value
}

func (iter *iterList[T]) GetPtr() *T {
	return &iter.current.value
}
