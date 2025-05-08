package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0228_summary_ranges(t *testing.T) {
	assert.Equal(t, []string{"1->5"}, lc_0228_summary_ranges([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []string{"1->5", "7", "9", "11->12"}, lc_0228_summary_ranges([]int{1, 2, 3, 4, 5, 7, 9, 11, 12}))
	assert.Equal(t, []string{"0->1", "3->4", "6->7", "11->13"}, lc_0228_summary_ranges([]int{0, 1, 3, 4, 6, 7, 11, 12, 13}))

}

func lc_0228_summary_ranges(nums []int) []string {
	result := []string{}
	i := 0
	for j := range len(nums) {
		if j == len(nums)-1 || nums[j+1]-nums[j] > 1 {
			if j-i >= 1 {
				result = append(result, fmt.Sprintf("%v->%v", nums[i], nums[j]))
			} else {
				result = append(result, fmt.Sprint(nums[i]))
			}
			i = j + 1
		}
	}
	return result
}
