package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0027_remove_element(t *testing.T) {
	cases := []struct {
		input, output []int
		val, res      int
	}{
		{input: []int{2, 5, 4, 3, 3, 4, 6, 1, 9, 3}, output: []int{2, 5, 4, 4, 6, 1, 9}, val: 3, res: 7},
		{input: []int{2, 5, 4, 3, 3, 4, 6, 1, 9, 3}, output: []int{2, 5, 3, 3, 6, 1, 9, 3}, val: 4, res: 8},
		{input: []int{2, 5, 4, 3, 3, 4, 6, 1, 9, 3}, output: []int{5, 4, 3, 3, 4, 6, 1, 9, 3}, val: 2, res: 9},
		{input: []int{2, 5, 4, 3, 3, 4, 6, 1, 9, 3}, output: []int{2, 5, 4, 3, 3, 4, 6, 1, 3}, val: 9, res: 9},
		{input: []int{2, 5, 4, 3, 3, 4, 6, 1, 9, 3}, output: []int{2, 5, 4, 3, 3, 4, 6, 1, 3}, val: 9, res: 9},
		{input: []int{2, 5, 5, 5, 3, 4, 6, 5, 5, 3, 5}, output: []int{2, 3, 4, 6, 3}, val: 5, res: 5},
	}
	for _, v := range cases {
		assert.Equal(t, v.res, lc_0027_remove_element(v.input, v.val))
		for i := 0; i < v.res; i++ {
			assert.Equal(t, v.output[i], v.input[i])
		}
	}
}

func lc_0027_remove_element(nums []int, val int) int {
	// i — позиция, на которую мы ищем подходящего кандидата.
	// j — метка последовательного перебора всех значений.
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			// Кандидат найден. Ищем кандидата на следующую позицию.
			if i != j {
				nums[i] = nums[j]
			}
			i++
		}
	}
	return i
}

// В слайсе есть ямы — последовательности val. Между ямами находятся блоки.
// Нужно устранить все ямы перемещением блоков влево и вернуть общую длину всех блоков.
// func lc_0027_remove_element_2(nums []int, val int) int {
// 	// Функция перемещает блок [start;finish] на l позиций.
// 	shift := func(l, start, finish int) {
// 		for i := start; i <= finish; i++ {
// 			nums[i-l] = nums[i]
// 		}
// 	}

// 	// a — начало ямы, b — конец ямы, c — конец блока.
// 	a, b, c, i, sum := -1, -1, -1, 0, len(nums)
// 	for {
// 		for ; i < len(nums); i++ {
// 			switch {
// 			case a == -1:
// 				if nums[i] == val {
// 					// Найдено начало самой первой ямы.
// 					a = i
// 					sum--
// 				}
// 			case b == -1:
// 				if nums[i] == val {
// 					// Яма продолжается.
// 					sum--
// 				} else {
// 					// Яма закончилась.
// 					b = i
// 				}
// 			case c == -1:
// 				if nums[i] == val {
// 					sum--
// 					// Блок кончается перед первым элементом следующей ямы.
// 					c = i - 1
// 				}
// 			}
// 		}
// 		switch {
// 		case b == -1:
// 			// Не найден конец ямы, значит, нечего перемещать.
// 			return sum
// 		case c == -1:
// 			// Не найден конец блока, значит, перемещаем все оставшиеся элементы.
// 			shift(b-a, b, len(nums)-1)
// 			return sum
// 		}
// 		shift(b-a, b, c)
// 		// Следующая яма начинается со следующего элемента после конца уже перемещённого блока.
// 		a = c + 1 - (b - a)
// 		// Поиск конца следующей ямы начинаем со следующего элемента после конца блока до перемещения.
// 		i = c + 2
// 		b = -1
// 		c = -1
// 	}
// }
