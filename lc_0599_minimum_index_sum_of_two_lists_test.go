package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-index-sum-of-two-lists
func Test_lc_0599_minimum_index_sum_of_two_lists(t *testing.T) {
	cases := []struct {
		l1, l2, result []string
	}{
		{[]string{"a", "b", "c"}, []string{"b", "d", "c", "a"}, []string{"b"}},
		{[]string{"a", "b", "c"}, []string{"q", "d", "c", "a"}, []string{"a"}},
		{[]string{"a", "b", "c"}, []string{"q", "d", "c", "k", "a"}, []string{"a", "c"}},
	}
	for _, c := range cases {
		_ = c
		assert.ElementsMatch(t, c.result, lc_0599_minimum_index_sum_of_two_lists(c.l1, c.l2))
	}
}

// Собираем один из массивов в хеш-таблицу, где ключ — строка, значение — индекс.
// Проходимся по второму массиву, и если строка есть в хеш-таблице, то считаем сумму индексов.
// Запоминаем сумму и добавляем к результату строки с такой же суммой.
// При нахождении меньшей суммы начинаем сбор заново. Суммы, которые больше, пропускаем.
// В итоге останется массив строк, суммы которых были наилучшими.
func lc_0599_minimum_index_sum_of_two_lists(list1 []string, list2 []string) []string {
	bestSum, result, m := 2000, []string{}, map[string]int{}
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}
	for i := range list1 {
		m[list1[i]] = i
	}
	for j := range list2 {
		i, ok := m[list2[j]]
		if !ok {
			continue
		}
		sum := i + j
		switch {
		case sum < bestSum:
			bestSum = sum
			result = []string{list2[j]}
		case sum == bestSum:
			result = append(result, list2[j])
		}
	}
	return result
}
