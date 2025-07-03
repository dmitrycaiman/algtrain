package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	cases := []struct{ scheme, str string }{
		{"1,3,5,7,9,1", "1,3,5,7,9,1"},
		{"1,3,5,7,9,1", "1,3,5,7,9,1"},
		{"1,3,5,7,9", "1,3,5,7,9"},
		{"1,1,1", "1,1"},
		{"", ""},
	}
	for _, c := range cases {
		assert.Equal(t, c.str, NewList(c.scheme).String())
	}
}
