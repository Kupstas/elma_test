package internal

import "strings"

type (
	List interface {
		Len() int
		Front() *ListItem
		Back() *ListItem
		PushFront(v string) *ListItem
		PushBack(v string) *ListItem
		Remove(i *ListItem)
		MoveToFront(i *ListItem)
	}
	ListItem struct {
		Value string
		Next  *ListItem
		Prev  *ListItem
	}

	list struct {
		size int
		head *ListItem
		tail *ListItem
	}
)

func (l *list) String() string {
	builder := strings.Builder{}
	el := l.head
	builder.WriteString("{")

	for el != nil {
		builder.WriteString(el.Value)
		el = el.Next
		if el != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v string) *ListItem {
	newElem := &ListItem{Value: v}
	head := l.Front()

	if head != nil {
		newElem.Next = head
		head.Prev = newElem
	} else {
		l.tail = newElem
	}

	l.head = newElem
	l.size++

	return newElem
}

func (l *list) PushBack(v string) *ListItem {
	newElem := &ListItem{Value: v}
	tail := l.Back()

	if tail != nil {
		newElem.Prev = tail
		tail.Next = newElem
	} else {
		l.head = newElem
	}

	l.tail = newElem
	l.size++

	return newElem
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	if next != nil {
		next.Prev = prev
	}

	if prev != nil {
		prev.Next = next
	}

	if next == nil {
		l.tail = prev
	}

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.Front() {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func NewList() List {
	return &list{}
}
