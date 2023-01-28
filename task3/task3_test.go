package task3

import (
	"testing"

	"github.com/Kupstas/elma_test/internal/list"

	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	l := list.New[string]().
		Add("1").
		Add("2").
		Add("3").
		Add("4")

	res := list.New[string]().
		Add("4").
		Add("3").
		Add("2").
		Add("1")

	ReverseLinkedList(l)

	reversedItems := make([]string, 0, l.Len())
	for v := range l.Iterate() {
		reversedItems = append(reversedItems, v)
	}

	originItems := make([]string, 0, res.Len())
	for v := range res.Iterate() {
		originItems = append(originItems, v)
	}

	assert.Equal(t, reversedItems, originItems)
}
