package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequences(t *testing.T) {
	for i := range listSize {
		assert.Equal(t, fibonacciNumbersList[i], fibonacciNumbers(i))
		assert.Equal(t, fibonacciNumbersList[i], fibonacciNumbersMemo(i))
		assert.Equal(t, tribonacciNumbersList[i], tribonacciNumbers(i))
		assert.Equal(t, tetranacciNumbersList[i], tetranacciNumbers(i))
		assert.Equal(t, pellNumbersList[i], pellNumbers(i))
		assert.Equal(t, padovanNumbersList[i], padovanNumbers(i))
		assert.Equal(t, jacobsthalNumbersList[i], jacobsthalNumbers(i))
	}
}
