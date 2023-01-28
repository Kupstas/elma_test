package list

type (
	List[T any] interface {
		Remove()
		Reverse()
		Add(v T) List[T]
		Len() int
		Get() (*T, bool)
		Iterate() chan T
	}
	listItem[T any] struct {
		value T
		next  *listItem[T]
	}

	list[T any] struct {
		size int
		head *listItem[T]
	}
)

func New[T any]() List[T] {
	return &list[T]{}
}

func (l *list[T]) Len() int {
	return l.size
}

func (l *list[T]) Add(v T) List[T] {
	newHead := &listItem[T]{value: v}
	newHead.next = l.head
	l.head = newHead

	l.size++
	return l
}

func (l *list[T]) Get() (*T, bool) {
	if l.head == nil {
		return nil, false
	}

	return &l.head.value, true
}

func (l *list[T]) Remove() {
	l.head = l.head.next
	l.size--
}

func (l *list[T]) Reverse() {
	curr := l.head
	var (
		prev *listItem[T]
		next *listItem[T]
	)

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *list[T]) Iterate() chan T {
	ch := make(chan T)
	go func() {
		for curr := l.head; curr != nil; curr = curr.next {
			ch <- curr.value
		}
		close(ch)
	}()

	return ch
}
