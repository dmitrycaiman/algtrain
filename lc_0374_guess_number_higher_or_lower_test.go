package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lc_0374_guess func(num int) int

func Test_lc_0374_guess_number_higher_or_lower(t *testing.T) {
	for i := range 100 {
		lc_0374_guess = func(num int) int {
			switch {
			case num > i:
				return -1
			case num < i:
				return 1
			default:
				return 0
			}
		}
		assert.Equal(t, i, lc_0374_guess_number_higher_or_lower(100))
	}
}

// Обыкновенный бинарный поиск.
func lc_0374_guess_number_higher_or_lower(n int) int {
	a, b := 0, n
	for {
		if a == b {
			return a
		}
		g := (b-a)/2 + a + 1
		switch lc_0374_guess(g) {
		case 0:
			return g
		case 1:
			a = g + 1
		case -1:
			b = g - 1
		}
	}
}
