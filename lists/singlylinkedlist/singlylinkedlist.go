package singlylinkedlist

type List[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

type node[T any] struct {
	data T
	next *node[T]
}

func New[T any](values ...T) *List[T] {
	l := List[T]{}

	if len(values) > 0 {
		for _, v := range values {
			l.Append(v)
		}
	}

	return &l
}

func (l *List[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var t T
		return t, false
	}

	current := l.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data, true
}

func (l *List[T]) Prepend(data T) {
	n := node[T]{data: data}

	if l.size == 0 {
		l.head = &n
		l.tail = &n
	} else {
		n.next = l.head
		l.head = &n
	}

	l.size++
}

func (l *List[T]) Append(data T) {
	n := node[T]{data: data}

	if l.size == 0 {
		l.head = &n
		l.tail = &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}

	l.size++
}

func (l *List[T]) Add(index int, data T) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.Prepend(data)
		return true
	}

	if index == l.size {
		l.Append(data)
		return true
	}

	n := node[T]{data: data}
	current := l.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	next := current.next
	current.next = &n
	n.next = next
	l.size++

	return true
}

func (l *List[T]) Delete(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}

	if index == 0 {
		return l.deleteHead()
	}

	if index == l.size-1 {
		return l.deleteTail()
	}

	current := l.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = current.next.next
	l.size--

	return true
}

func (l *List[T]) deleteHead() bool {
	l.head = l.head.next
	l.size--

	return true
}

func (l *List[T]) deleteTail() bool {
	current := l.head

	for i := 0; i < l.size-2; i++ {
		current = current.next
	}

	l.tail = current
	l.tail.next = nil
	l.size--

	return true
}

func (l *List[T]) Values() []T {
	values := make([]T, l.size)

	current := l.head

	for i := 0; i < l.size; i++ {
		values[i] = current.data
		current = current.next
	}

	return values
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Size() int {
	return l.size
}
