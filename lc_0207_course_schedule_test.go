package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/course-schedule
func Test_lc_0207_course_schedule(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		result        bool
	}{
		{3, [][]int{{1, 0}, {0, 2}, {2, 1}}, false},
		{3, [][]int{{1, 0}, {0, 2}, {2, 0}}, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0207_course_schedule(c.numCourses, c.prerequisites))
	}
}

// Собираем в мапу для каждого курса массив пререквизитов. Итерируемся по одному из курсов относительно его пререквизитов.
// Если в пререквизитах или в их пререквизитах (и так далее рекурсивно) встречается наш курс, то на лицо зацикливание,
// которое сразу даёт отрицательный ответ. Ищем циклы.
func lc_0207_course_schedule(_ int, p [][]int) bool {
	m := map[int][]int{}
	for _, v := range p {
		m[v[0]] = append(m[v[0]], v[1])
	}
	for i := range m {
		if !lc_0207(i, map[int]struct{}{}, m) {
			return false
		}
	}
	return true
}

func lc_0207(i int, seen map[int]struct{}, m map[int][]int) bool {
	s, ok := m[i]
	if !ok { // Данный курс не имеет пререквизитов.
		return true
	}
	if len(m) != 0 { // Если мы только начали проход, то не попадаем сюда.
		for _, v := range s {
			if _, ok := seen[v]; ok { // Нашли цикл, возвращаем отрицательный ответ.
				return false
			}
		}
	}
	seen[i] = struct{}{}  // То, что мы уже видели, в том числе и текущий курс. Если встретим данные значения, то цикл найден.
	for _, v := range s { // Смотрим пререквизиты текущего курса на предмет циклов.
		if !lc_0207(v, seen, m) {
			return false // Нашли любой цикл где-то в пререквизитах.
		}
	}
	delete(seen, i) // Удаляем курс i из мапы увиденного: в других проходах вне этой функции он не участвует.
	delete(m, i)    // Удаляем курс i из мапы на проверку: мы его уже проверили.
	return true
}
