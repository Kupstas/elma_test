package task2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceShift(t *testing.T) {
	type test struct {
		input  []int
		output []int
		shift  int
	}

	tests := []test{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{5, 1, 2, 3, 4},
			shift:  1,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{4, 5, 1, 2, 3},
			shift:  2,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{3, 4, 5, 1, 2},
			shift:  3,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{2, 3, 4, 5, 1},
			shift:  4,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
			shift:  5,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{5, 1, 2, 3, 4},
			shift:  6,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.output, SliceShift(tt.input, tt.shift))
	}
}
