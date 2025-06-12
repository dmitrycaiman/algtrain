package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array
func Test_lc_3423_maximum_difference_between_adjacent_elements_in_a_circular_array(t *testing.T) {
	cases := []struct{}{}
	for _, c := range cases {
		_ = c
		//lc_3423_maximum_difference_between_adjacent_elements_in_a_circular_array()
		assert.Equal(t, 1, 1)
	}
}

// Проитерироваться по всем разницам и найти наибольшую. Учесть краевой случай со с сравнением первого и последнего элементов.
func lc_3423_maximum_difference_between_adjacent_elements_in_a_circular_array(nums []int) int {
	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}
	res, nums := 0, append(nums, nums[0])
	for i := range len(nums) - 1 {
		if dif := abs(nums[i] - nums[i+1]); dif > res {
			res = dif
		}
	}
	return res
}
