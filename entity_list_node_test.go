package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	cases := []struct {
		scheme, noCycle, withCycle string
		cyclePos                   int
	}{
		{"1,3,5,7,9", "1,3,5,7,9", "1,3,5,7,9,3", 1},
		{"1,3,3,3,3", "1,3,3,3,3", "1,3,3,3,3,1", 0},
		{"1,1,1,0", "1,1,1,0", "1,1,1,0", -1},
		{"1", "1", "1,1", 0},
		{"1", "1", "1", 1},
		{"", "", "", 0},
		{"", "", "", 10},
	}
	for _, c := range cases {
		assert.Equal(t, c.noCycle, NewList(c.scheme).String())
		assert.Equal(t, c.withCycle, NewListWithCycle(c.scheme, c.cyclePos).String())
	}
}
