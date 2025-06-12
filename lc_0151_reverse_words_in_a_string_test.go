package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0151_reverse_words_in_a_string(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{"ke lo bo mda", "mda bo lo ke"},
		{"     ke lo bo mda", "mda bo lo ke"},
		{"     ke lo bo mda      ", "mda bo lo ke"},
		{"     ke    lo    bo      mda      ", "mda bo lo ke"},
		{"ke", "ke"},
		{"ke ko", "ko ke"},
		{"k", "k"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0151_reverse_words_in_a_string(c.input))
	}
}

// Проходим массив. В процессе запоминаем позиции начала последнего встреченного слова и последнего ведущего пробела.
// Если был встречен пробел, то:
// 1. Запоминаем его позицию и пропускаем все следующие пробелы.
// 2. Разворачиваем часть от начала последнего встреченного слова до данного пробела невключительно.
// Если был встречен иной символ, то:
// 1. Запоминаем его позицию как начало последнего встреченного слова.
// 2. Вырезаем из массива часть от последнего ведущего пробела до данного символа невключительно.
// В конце разворачиваем весь получившийся массив.
// По итогу мы удалили все пробелы, развернули каждое слово по отдельности, а потом развернули всё ещё раз, восстановив порядок букв в словах, но изменив их порядок.
func lc_0151_reverse_words_in_a_string(s string) string {
	b := []byte(s)
	reverse := func(i, j int) {
		for ; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	metSpace, spacePos, w := b[0] == 32, -1, 0
	for i := 0; i < len(b); i++ {
		switch {
		case metSpace && b[i] != 32:
			if i-spacePos > 1 {
				b = append(b[:spacePos+1], b[i:]...)
				i = spacePos + 1
			}
			metSpace = false
			w = i
		case !metSpace && b[i] == 32:
			spacePos = i
			metSpace = true
			reverse(w, i-1)
		}
	}
	if metSpace {
		b = b[:spacePos]
	} else {
		reverse(w, len(b)-1)
	}
	reverse(0, len(b)-1)
	return string(b)
}
