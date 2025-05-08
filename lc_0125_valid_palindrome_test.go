package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0125_valid_palindrome(t *testing.T) {
	assert.True(t, lc_0125_valid_palindrome("abba"))
	assert.True(t, lc_0125_valid_palindrome(".a"))
	assert.True(t, lc_0125_valid_palindrome(""))
	assert.True(t, lc_0125_valid_palindrome("1"))
	assert.True(t, lc_0125_valid_palindrome("A man, a plan, a canal: Panama"))
	assert.True(t, lc_0125_valid_palindrome("ABba"))
	assert.True(t, lc_0125_valid_palindrome(`;Ab\b;B/.a\\\`))
	assert.True(t, lc_0125_valid_palindrome(`abc\\\cba`))
	assert.True(t, lc_0125_valid_palindrome(`;]]]'`))
	assert.False(t, lc_0125_valid_palindrome(`race a car`))
}

func lc_0125_valid_palindrome(s string) bool {
	i, j, b, dif := 0, len(s)-1, []byte(s), byte('a'-'A')
	check := func(n int) bool {
		switch {
		case b[n] >= '0' && b[n] <= '9':
		case b[n] >= 'a' && b[n] <= 'z':
		case b[n] >= 'A' && b[n] <= 'Z':
			b[n] += dif
		default:
			return false
		}
		return true
	}
	for {
		if i >= j {
			return true
		}
		for !check(i) {
			i++
			if i == j {
				return true
			}
		}
		for !check(j) {
			j--
		}
		if b[i] != b[j] {
			return false
		}
		i++
		j--
	}
}
