package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0977_squares_of_a_sorted_array(t *testing.T) {
	cases := []struct {
		nums, result []int
	}{
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{-7, -7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 49, 121}},
		{[]int{-100, -1, 1, 10}, []int{1, 1, 100, 10000}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0977_squares_of_a_sorted_array(c.nums))
	}
}

func lc_0977_squares_of_a_sorted_array(nums []int) []int {
	result, counter := make([]int, len(nums)), len(nums)-1
	for i, j := 0, len(nums)-1; i <= j; {
		r, l := nums[i]*nums[i], nums[j]*nums[j]
		switch {
		case i == j:
			fallthrough
		case r > l:
			result[counter] = r
			i++
		case r < l:
			result[counter] = l
			j--
		default:
			result[counter], result[counter-1] = r, l
			counter--
			i++
			j--
		}
		counter--
	}
	return result
}
