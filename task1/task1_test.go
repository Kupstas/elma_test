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
		{row: "1+(2*3{3})", expected: true},
		{row: "(some test{s}[1]((aa)))", expected: true},
		{row: "и даже русский (текст) поддреживается", expected: true},
		{row: "({[]})", expected: true},
		{row: "{[()]}", expected: true},
		{row: "[({})]", expected: true},
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
