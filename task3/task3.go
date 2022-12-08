package task3

import (
	"elma_test/internal"
)

func ReverseLinkedList(list internal.List) internal.List {
	newList := internal.NewList()

	back := list.Front()
	for back != nil {
		newList.PushFront(back.Value)
		back = back.Next
	}

	return newList
}
