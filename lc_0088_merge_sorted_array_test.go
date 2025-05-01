package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0088_merge_sorted_array(t *testing.T) {
	cases := []struct {
		nums1, nums2, res []int
		m, n              int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, []int{1, 2, 2, 3, 5, 6}, 3, 3},
		{[]int{1, 2, 5, 9, 0, 0, 0, 0}, []int{2, 7, 7, 8}, []int{1, 2, 2, 5, 7, 7, 8, 9}, 4, 4},
		{[]int{0, 2, 2, 3, 6, 9, 9, 9, 0, 0, 0, 0, 0, 0}, []int{4, 4, 6, 6, 8, 9}, []int{0, 2, 2, 3, 4, 4, 6, 6, 6, 8, 9, 9, 9, 9}, 8, 6},
		{[]int{0, 0, 0}, []int{1, 1, 1}, []int{1, 1, 1}, 0, 3},
	}
	for _, v := range cases {
		lc_0088_merge_sorted_array(v.nums1, v.m, v.nums2, v.n)
		assert.Equal(t, v.res, v.nums1)
	}
}

func lc_0088_merge_sorted_array(nums1 []int, m int, nums2 []int, n int) {
	totalInjects := 0
	inject := func(pos, val int) {
		for i := m + totalInjects; i > pos; i-- {
			nums1[i] = nums1[i-1]
		}
		nums1[pos] = val
		totalInjects++
	}

	i, j := 0, 0
	for i < m+n-1 && j < n {
		if nums1[i] >= nums2[j] {
			inject(i, nums2[j])
			i++
			j++
		} else {
			i++
		}
	}

	if rest := n - totalInjects; rest != 0 {
		for i := 0; i < rest; i++ {
			nums1[m+n-1-i] = nums2[n-1-i]
		}
	}
}
