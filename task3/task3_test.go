package task3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"elma_test/internal"
)

func TestReverseLinkedList(t *testing.T) {
	list := internal.NewList()
	list.PushFront("1")
	list.PushFront("2")
	list.PushFront("3")
	list.PushFront("4")

	res := internal.NewList()
	res.PushFront("4")
	res.PushFront("3")
	res.PushFront("2")
	res.PushFront("1")

	reversed := ReverseLinkedList(list)
	item := reversed.Front()
	item2 := res.Front()

	for item != nil {
		assert.Equal(t, item.Value, item2.Value)
		item = item.Next
		item2 = item2.Next
	}
}
