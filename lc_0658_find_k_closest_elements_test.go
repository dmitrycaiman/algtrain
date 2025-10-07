package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-k-closest-elements
func Test_lc_0658_find_k_closest_elements(t *testing.T) {
	cases := []struct {
		arr    []int
		k, x   int
		result []int
	}{
		{[]int{1}, 1, 1, []int{1}},
		{[]int{1, 5, 9, 15}, 2, -1, []int{1, 5}},
		{[]int{1, 5, 9, 15}, 2, 5, []int{1, 5}},
		{[]int{1, 5, 9, 15}, 3, 5, []int{1, 5, 9}},
		{[]int{1, 1, 1, 10, 10, 10}, 1, 9, []int{10}},
		{[]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4, []int{0, 1, 1, 1, 2, 3, 6, 7, 8}},
		{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
		{[]int{-2, -1, 1, 2, 3, 4, 5}, 7, 3, []int{-2, -1, 1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2, []int{1, 3}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0658_find_k_closest_elements(c.arr, c.k, c.x))
	}
}

// С помощью бинарного поиска ищем либо сам целевой элемент,
// либо самый близкий к нему (если выбирать из соседей, то берём меньшего).
// Далее посредством двух указателей расширяемся влево и вправо, в зависимости от того,
// какой из соседних элементов ближе по значению к целевому.
func lc_0658_find_k_closest_elements(arr []int, k int, x int) []int {
	a, b, index := 0, len(arr)-1, 0
	for {
		m := a + (b-a)/2
		switch {
		case x < arr[a]: // Элемент за левой границей.
			index = a
		case x > arr[b]: // Элемент за правой границей.
			index = b
		case b-a <= 1: // Случай для длин 1 и 2. В обоих случаях будет выбран ближайший меньший элемент.
			if arr[b]-x >= x-arr[a] {
				index = a
			} else {
				index = b
			}
		case arr[a] <= x && x <= arr[m]:
			b = m
			continue
		case arr[m] <= x && x <= arr[b]:
			a = m
			continue
		}
		break
	}
	a, b = index, index+1 // При таких значениях границ мы можем отдать даже единственный индекс, когда k == 1.
	for k != 1 {          // Можем потратить k единиц для формирования ответа. Одна уходит на индекс, найденный выше.
		// По условию задачи k не превышает длину массива, поэтому при невалидности одной границы смело двигаем другую.
		switch {
		case b == len(arr):
			a--
		case a == 0:
			b++
		case arr[b]-x >= x-arr[a-1]: // Сравниваем расстояние валидных границ до целевого элемента.
			a--
		default:
			b++
		}
		k-- // Потратили одну из k единиц на передвижение границ.
	}
	return arr[a:b]
}
