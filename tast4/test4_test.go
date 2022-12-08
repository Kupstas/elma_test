package tast4

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	buf := bytes.Buffer{}

	stack := NewStack()
	stack.Add(3)
	stack.PrintMax(&buf)
	assert.Equal(t, strconv.Itoa(3), buf.String())

	stack.Add(5)
	buf.Reset()
	stack.PrintMax(&buf)
	assert.Equal(t, strconv.Itoa(5), buf.String())

	stack.Add(2)
	buf.Reset()
	stack.PrintMax(&buf)
	assert.Equal(t, strconv.Itoa(5), buf.String())

	stack.Remove()
	stack.Add(3)
	buf.Reset()
	stack.PrintMax(&buf)
	assert.Equal(t, strconv.Itoa(5), buf.String())

	stack.Add(6)
	buf.Reset()
	stack.PrintMax(&buf)
	assert.Equal(t, strconv.Itoa(6), buf.String())

	buf.Reset()
	stack.Remove()
	stack.PrintMax(&buf)
	assert.Equal(t, strconv.Itoa(5), buf.String())
}
