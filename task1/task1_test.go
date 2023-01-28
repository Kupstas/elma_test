package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBracketSequence(t *testing.T) {
	type test struct {
		row      string
		expected bool
	}

	tests := []test{
		{row: "([{}])", expected: true},
		{row: "[{()}]", expected: true},
		{row: "{[()]}", expected: true},
		{row: "()", expected: true},
		{row: "{}", expected: true},
		{row: "[]", expected: true},
		{row: "[](){}", expected: true},
		{row: "{", expected: false},
		{row: "(", expected: false},
		{row: "[", expected: false},
		{row: "}", expected: false},
		{row: ")", expected: false},
		{row: "]", expected: false},
		{row: "({[]))", expected: false},
		{row: "{{()]}", expected: false},
		{row: "[({])]", expected: false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, CheckBracketSequence(tt.row))
	}
}
