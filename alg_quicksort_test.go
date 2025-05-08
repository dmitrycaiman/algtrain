package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alg_quicksort(t *testing.T) {
	cases := []struct{ input, output []int }{
		{input: []int{1, 2, 3, 4, 5, 6}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{3, 1, 4, 2, 5, 6}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{2, 4, 5, 6, 3, 1}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{3, 2, 1, 5, 4, 6}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{3, 1, 5, 6, 4, 2}, output: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, v := range cases {
		alg_quicksort(v.input)
		assert.Equal(t, v.output, v.input)
	}
}

// Требует упрощения.
func alg_quicksort(input []int) {
	i, j, pivot := 0, len(input)-1, input[0]
	for j > i {
		switch {
		case input[i] >= pivot && input[j] < pivot:
			input[i], input[j] = input[j], input[i]
			i++
			j--
		case input[i] < pivot && input[j] >= pivot:
			i++
			j--
		case input[i] < pivot:
			i++
		case input[j] >= pivot:
			j--
		}
	}
	for k, v := range input {
		if v >= pivot && k != 0 {
			alg_quicksort(input[:k])
			alg_quicksort(input[k:])
			return
		}
	}
}
