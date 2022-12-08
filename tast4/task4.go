package tast4

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type (
	Stack struct {
		max  int
		val  int
		size int
		elem *item
	}

	item struct {
		val  int
		next *item
	}
)

func (s *Stack) Add(elem int) {
	if elem > s.max {
		s.max = elem
	}

	str := &item{
		val:  elem,
		next: s.elem,
	}

	s.elem = str
	s.size++
}

func (s *Stack) Remove() {
	s.elem = s.elem.next
	el := s.elem
	max := el.val

	for el != nil {
		if el.val > max {
			max = el.val
		}

		el = el.next
	}
	s.max = max
	s.size--
}

func (s *Stack) PrintMax(writer io.Writer) {
	fmt.Fprint(writer, s.max)
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) String() string {
	builder := strings.Builder{}
	el := s.elem
	builder.WriteString("{")

	for el != nil {
		builder.WriteString(strconv.Itoa(el.val))
		el = el.next
		if el != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}
