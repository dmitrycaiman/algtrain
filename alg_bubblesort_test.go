package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alg_bubblesort(t *testing.T) {
	cases := []struct{ input, output []int }{
		{input: []int{1, 2, 3, 4, 5, 6}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{3, 1, 4, 2, 5, 6}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{2, 4, 5, 6, 3, 1}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{3, 2, 1, 5, 4, 6}, output: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{3, 1, 5, 6, 4, 2}, output: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, v := range cases {
		alg_bubblesort(v.input)
		assert.Equal(t, v.output, v.input)
	}
}

func alg_bubblesort(in []int) []int {
	for {
		swap := false
		for i := 0; i <= len(in)-2; i++ {
			if in[i+1] < in[i] {
				in[i], in[i+1] = in[i+1], in[i]
				swap = true
			}
		}
		if !swap {
			return in
		}
	}
}
