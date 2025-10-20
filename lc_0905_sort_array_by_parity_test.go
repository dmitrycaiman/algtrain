package main

import (
	"testing"
)

// https://leetcode.com/problems/sort-array-by-parity
func Test_lc_0905_sort_array_by_parity(t *testing.T) {
	cases := []struct{ input []int }{
		{[]int{2, 2, 3, 3}},
		{[]int{2, 2, 3, 3, 1}},
		{[]int{2, 3, 1, 1, 1, 2, 3, 3, 1}},
	}
	for _, c := range cases {
		odd := false
		for _, v := range lc_0905_sort_array_by_parity(c.input) {
			switch {
			case !odd && v%2 == 0:
			case !odd && v%2 == 1:
				odd = true
			case odd && v%2 == 1:
			default:
				t.Fatalf("%v is on wrong position", v)
			}
		}
	}
}

// 100/85
// Два указателя. Один в начале, другой в конце.
// Двигаем указатели, если на их местах стоят правильные элементы.
// В противном случае меняем элементы местами.
// Выходим, когда указатели встречаются.
func lc_0905_sort_array_by_parity(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		switch {
		case nums[i]%2 == 0:
			i++
		case nums[j]%2 == 1:
			j--
		default:
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
