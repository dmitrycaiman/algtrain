package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0169_majority_element(t *testing.T) {
	assert.Equal(t, 1, lc_0169_majority_element([]int{1, 1, 3, 1, 7, 1, 9, 9, 6, 1, 1, 1}))
	assert.Equal(t, 1, lc_0169_majority_element([]int{1}))
	assert.Equal(t, 1, lc_0169_majority_element([]int{1, 1, 3, 3, 1}))
	assert.Equal(t, 1, lc_0169_majority_element([]int{1, 1, 1}))
}

// Мажоритарный элемент в массиве длиной n встречается n/2+1 и более раз.
func lc_0169_majority_element(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	mark := len(nums)/2 + 1
	for k, v := range m {
		if v >= mark {
			return k
		}
	}
	return 0
}
