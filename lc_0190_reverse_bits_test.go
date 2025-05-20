package main

import (
	"math/bits"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0190_reverse_bits(t *testing.T) {
	for range 10000 {
		input := rand.Uint32()
		assert.Equal(t, bits.Reverse32(input), lc_0190_reverse_bits(input))
	}
}

// Проходимся по исходному числу от младшего бита к старшему посредством побитового умножения на соответствующую степень двойки. Таким
// образом мы определяем значение бита. Если оно равно 1 (ненулевой результат умножения), то мы прибавляем 1 к общему результату.
// До каждой итерации общий результат умножается на 2: так формируется сдвиг влево предыдущих полученных значений.
// Проходимся таким образом по всей длине 32-х битового числа.
func lc_0190_reverse_bits(num uint32) uint32 {
	var res, i uint32 = 0, 1
	for range 32 {
		res <<= 1
		if num&i != 0 {
			res++
		}
		i <<= 1
	}
	return res
}
