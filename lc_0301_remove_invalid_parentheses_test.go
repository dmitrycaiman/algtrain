package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0301_remove_invalid_parentheses(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{"()())()", []string{"(())()", "()()()"}},
		{")(", []string{""}},
		{"(((k()((", []string{"k()", "(k)"}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.output, lc_0301_remove_invalid_parentheses(c.input))
	}
}

func lc_0301_remove_invalid_parentheses(s string) []string {
	as, bs := []int{}, []int{}
	ac, bc := 0, 0
	for i, v := range s {
		switch v {
		case '(':
			as = append(as, i)
			ac++
		case ')':
			bs = append(bs, i)
			if ac == 0 {
				bc++
			} else {
				ac--
			}
		}
	}

	aps, bps := lc_0301_permutate(ac, as), lc_0301_permutate(bc, bs)
	switch {
	case len(aps) == 0:
		aps = [][]int{nil}
	case len(bps) == 0:
		bps = [][]int{nil}
	}
	results, keys := map[string]struct{}{}, []string{}
	for _, ap := range aps {
		resa := []byte(s)
		for _, p := range ap {
			resa[p] = '-'
		}
		for _, bp := range bps {
			resb := make([]byte, len(resa))
			copy(resb, resa)
			for _, p := range bp {
				resb[p] = '-'
			}
			if lc_0301_validate(resb) {
				result := []byte{}
				for _, v := range resb {
					if v != '-' {
						result = append(result, v)
					}
				}
				key := string(result)
				if _, ok := results[key]; !ok {
					results[key] = struct{}{}
					keys = append(keys, key)
				}
			}
		}
	}

	return keys
}

func lc_0301_validate(scheme []byte) bool {
	counter := 0
	for _, v := range scheme {
		switch v {
		case '(':
			counter++
		case ')':
			if counter == 0 {
				return false
			}
			counter--
		}
	}
	return counter == 0
}

func lc_0301_permutate(size int, set []int) [][]int {
	return lc_0301_permutateRec(0, size, []int{}, set, [][]int{})
}

func lc_0301_permutateRec(index, size int, result, set []int, results [][]int) [][]int {
	switch {
	case len(result) == size:
		return append(results, result)
	case index == len(set):
		return results
	}
	resultCopy := make([]int, len(result))
	copy(resultCopy, result)
	return lc_0301_permutateRec(
		index+1,
		size,
		append(resultCopy, set[index]),
		set,
		lc_0301_permutateRec(index+1, size, result, set, results),
	)
}
