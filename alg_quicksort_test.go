package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alg_quicksort(t *testing.T) {
	testFlow := func(f func(input []int)) {
		cases := []struct{ input, output []int }{
			{input: []int{20, 23, 27, 32, 38, 4, 12, 20}, output: []int{4, 12, 20, 20, 23, 27, 32, 38}},
			{input: []int{1, 2, 3, 4, 1, 5}, output: []int{1, 1, 2, 3, 4, 5}},
			{input: []int{1, 2, 3, 4, 5, 6}, output: []int{1, 2, 3, 4, 5, 6}},
			{input: []int{6, 1, 4, 3, 5, 2}, output: []int{1, 2, 3, 4, 5, 6}},
			{input: []int{2, 4, 5, 6, 3, 1}, output: []int{1, 2, 3, 4, 5, 6}},
			{input: []int{3, 2, 1, 5, 4, 6}, output: []int{1, 2, 3, 4, 5, 6}},
			{input: []int{3, 1, 5, 6, 4, 2}, output: []int{1, 2, 3, 4, 5, 6}},
			{input: []int{28, 37, 2, 7, 20, 45, 1, 13, 40}, output: []int{1, 2, 7, 13, 20, 28, 37, 40, 45}},
		}
		for _, v := range cases {
			f(v.input)
			assert.Equal(t, v.output, v.input)
		}
		for i := 1; i <= 1000; i++ {
			sorted, nonSorted, n, m := []int{}, []int{}, 0, map[int]struct{}{}
			for j := range i {
				n += rand.Intn(10)
				sorted = append(sorted, n)
				m[j] = struct{}{}
			}
			for k := range m {
				nonSorted = append(nonSorted, sorted[k])
			}
			f(nonSorted)
			assert.Equal(t, sorted, nonSorted)
		}
	}
	testFlow(alg_quicksort_1)
	testFlow(alg_quicksort_2)
}

func alg_quicksort_1(input []int) {
	switch len(input) {
	case 0, 1:
		return
	case 2:
		if input[0] > input[1] {
			input[0], input[1] = input[1], input[0]
		}
		return
	}
	i, j, pivot := 0, len(input)-1, input[len(input)/2]
	for i <= j {
		for input[i] < pivot {
			i++
		}
		for input[j] > pivot {
			j--
		}
		if i <= j {
			input[i], input[j] = input[j], input[i]
			i++
			j--
		}
	}
	alg_quicksort_1(input[:i])
	alg_quicksort_1(input[i:])
}

func alg_quicksort_2(input []int) {
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
			alg_quicksort_2(input[:k])
			alg_quicksort_2(input[k:])
			return
		}
	}
}
