package tl

type ListElement[T any] struct {
	value      T
	prev, next *ListElement[T]
	list       *List[T]
}

func (e *ListElement[T]) Get() T {
	return e.value
}

func (e *ListElement[T]) GetPtr() *T {
	return &e.value
}

func (e *ListElement[T]) Drop() {
	if e.prev != nil {
		e.prev.next = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	}

	e.list.size--
}

type List[T any] struct {
	root ListElement[T]
	size int
}

func (list *List[T]) Size() int {
	return list.size
}

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

func (list *List[T]) Front() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)
	}

	return
}

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

func (list *List[T]) PopFront() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)

		list.root.next.Drop()
	}

	return
}

func (list *List[T]) PopBack() (opt OptionalPtr[T]) {
	if list.root.prev != nil {
		opt.Set(&list.root.prev.value)

		list.root.prev.Drop()
	}

	return
}

func (list *List[T]) Reset() {
	list.root.next = nil
	list.root.prev = nil
}

type iterList[T any] struct {
	root       *ListElement[T]
	current    *ListElement[T]
	prev, next *ListElement[T]
}

func (list *List[T]) Iter() IterDropBidir[T] {
	return &iterList[T]{
		root: &list.root,
		next: list.root.next,
		prev: list.root.prev,
	}
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
