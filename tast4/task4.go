package tast4

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Kupstas/elma_test/internal/list"
)

type (
	Stack struct {
		max int
		l   list.List[int]
	}
)

func NewStack() *Stack {
	return &Stack{
		l: list.New[int](),
	}
}

func (s *Stack) Add(elem int) {
	if elem > s.max {
		s.max = elem
	}

	s.l.Add(elem)
}

func (s *Stack) Remove() {
	s.l.Remove()

	v, _ := s.l.Get()
	newMax := *v
	for v := range s.l.Iterate() {
		if v > newMax {
			newMax = v
		}
	}

	s.max = newMax
	fmt.Println(s)
}

func (s *Stack) PrintMax(writer io.Writer) {
	fmt.Fprint(writer, s.max)
}

func (s *Stack) String() string {
	builder := strings.Builder{}
	builder.WriteString("{")

	for v := range s.l.Iterate() {
		builder.WriteString(strconv.Itoa(v))
		builder.WriteString(", ")
	}

	builder.WriteString("}")
	return builder.String()
}
