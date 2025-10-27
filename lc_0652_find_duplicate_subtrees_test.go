package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-duplicate-subtrees
func Test_lc_0652_find_duplicate_subtrees(t *testing.T) {
	cases := []struct {
		root   string
		result []string
	}{
		{"1,2,3,4,null,2,4,null,null,4", []string{"2,4", "4"}},
		{"10,2,22,1,12,1,1", []string{"1"}},
		{"2,2,2,3,null,3,null", []string{"2,3", "3"}},
		{"1,2,3,2,3,2,1,1,1,2,1", []string{"3,2,1", "1", "2"}},
		{"1,2,2", []string{"2"}},
	}
	for _, c := range cases {
		result := []string{}
		for _, v := range lc_0652_find_duplicate_subtrees(NewTree(c.root)) {
			result = append(result, v.String())
		}
		assert.ElementsMatch(t, c.result, result)
	}
}

// 68/65
// Сериализуем каждое поддерево при postorder-обходе и считаем частоту встреченных поддеревьев.
// В результат добавляем те, которые встречаются больше одного раза.
func lc_0652_find_duplicate_subtrees(root *TreeNode) []*TreeNode {
	storage := map[string][]*TreeNode{}
	_ = lc_0652_rec(root, storage)
	result := []*TreeNode{}
	for _, roots := range storage {
		if len(roots) > 1 {
			result = append(result, roots[0])
		}
	}
	return result
}

func lc_0652_rec(node *TreeNode, storage map[string][]*TreeNode) []byte {
	if node == nil {
		return []byte{255} // Обозначаем пустой узел.
	}
	r := append(
		append(
			lc_0652_rec(node.Left, storage),     // Значение левой части.
			lc_0652_rec(node.Right, storage)..., // Значение правой части.
		),
		byte((node.Val)), // Значение узла.
	)
	if node.Val < 0 {
		r = append(r, 254) // Обозначаем знак минуса при значении узла.
	}
	r = append(r, 253) // Экранируем узел.
	key := string(r)
	storage[key] = append(storage[key], node) // Добавляем корень поддерева в хранилище.
	return r
}
