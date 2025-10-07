package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/summary-ranges
func Test_lc_0228_summary_ranges(t *testing.T) {
	assert.Equal(t, []string{"1->5"}, lc_0228_summary_ranges([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []string{"1->5", "7", "9", "11->12"}, lc_0228_summary_ranges([]int{1, 2, 3, 4, 5, 7, 9, 11, 12}))
	assert.Equal(t, []string{"0->1", "3->4", "6->7", "11->13"}, lc_0228_summary_ranges([]int{0, 1, 3, 4, 6, 7, 11, 12, 13}))

}

// Два указателя. Первый движется до момента, пока не встретит число, которое отличается от следующего на единицу и более.
// В этом случае записываем в массив результата промежуток: в зависимости от его размера — одно или два числа.
// После сдвигаем второй указатель на позицию первого и повторяем так до конца массива.
func lc_0228_summary_ranges(nums []int) []string {
	result := []string{}
	for i, j := 0, 0; j < len(nums); j++ {
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
