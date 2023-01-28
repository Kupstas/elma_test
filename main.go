package main

import (
	"bytes"
	"fmt"

	"github.com/Kupstas/elma_test/internal/list"
	"github.com/Kupstas/elma_test/task1"
	"github.com/Kupstas/elma_test/task2"
	"github.com/Kupstas/elma_test/task3"
	"github.com/Kupstas/elma_test/tast4"
)

func main() {
	t1()
	t2()
	t3()
	t4()
}

func t1() {
	row := "(some test{s}[1]((aa)))"
	check := task1.CheckBracketSequence(row)
	res := "NO"
	if check {
		res = "YES"
	}
	fmt.Println("task1:")
	fmt.Printf("\trow: %s\n", row)
	fmt.Printf("\tresult: %s\n", res)
}

func t2() {
	data := []int{1, 2, 3, 4, 5}
	shift := 3

	fmt.Println("task2:")
	fmt.Printf("\tinput: %v\n", data)
	fmt.Printf("\tshift: %d\n", shift)
	fmt.Printf("\tresult: %v\n", task2.SliceShift(data, shift))
}

func t3() {
	list := list.New[string]()
	list.Add("1")
	list.Add("2")
	list.Add("3")

	fmt.Println("task3:")
	fmt.Printf("\tlist: %s\n", list)
	task3.ReverseLinkedList(list)
	fmt.Printf("\tres: %s\n", list)
}

func t4() {
	buf := bytes.Buffer{}
	stack := tast4.NewStack()
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)
	stack.PrintMax(&buf)

	fmt.Println("task4:")
	fmt.Printf("\tstack: %s\n", stack)
	fmt.Printf("\tmax: %s\n", buf.String())

	stack.Add(5)
	fmt.Printf("\tadd item 5: %s\n", stack)

	buf.Reset()
	stack.PrintMax(&buf)
	fmt.Printf("\tmax: %s\n", buf.String())

	buf.Reset()
	stack.Remove()
	stack.PrintMax(&buf)
	fmt.Printf("\tremove last element: %s\n", stack)
	fmt.Printf("\tremove last element: %s\n", buf.String())
}
