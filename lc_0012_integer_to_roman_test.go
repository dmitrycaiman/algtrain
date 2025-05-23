package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0012_integer_to_roman(t *testing.T) {
	assert.Equal(t, "I", lc_0012_integer_to_roman(1))
	assert.Equal(t, "II", lc_0012_integer_to_roman(2))
	assert.Equal(t, "III", lc_0012_integer_to_roman(3))
	assert.Equal(t, "IV", lc_0012_integer_to_roman(4))
	assert.Equal(t, "V", lc_0012_integer_to_roman(5))
	assert.Equal(t, "VI", lc_0012_integer_to_roman(6))
	assert.Equal(t, "VII", lc_0012_integer_to_roman(7))
	assert.Equal(t, "VIII", lc_0012_integer_to_roman(8))
	assert.Equal(t, "IX", lc_0012_integer_to_roman(9))
	assert.Equal(t, "X", lc_0012_integer_to_roman(10))
	assert.Equal(t, "XI", lc_0012_integer_to_roman(11))
	assert.Equal(t, "XII", lc_0012_integer_to_roman(12))
	assert.Equal(t, "MMMDCCXLIX", lc_0012_integer_to_roman(3749))

}

func lc_0012_integer_to_roman(num int) string {
	sym := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	val := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000, 10000}
	res := ""
	for num != 0 {
		for i := range val {
			if num < val[i] {
				res += sym[i-1]
				num -= val[i-1]
				break
			}
		}
	}
	return res
}
