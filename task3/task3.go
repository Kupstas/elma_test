package task3

import "github.com/Kupstas/elma_test/internal/list"

func ReverseLinkedList[T any](l list.List[T]) {
	l.Reverse()
}
