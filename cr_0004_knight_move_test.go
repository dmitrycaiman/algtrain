package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://coderun.yandex.ru/problem/knight-move
func Test_cr_0004_knight_move(t *testing.T) {
	cases := []struct{ input, output string }{
		{"3 2\n", "1"},
		{"31 34\n", "293930"},
	}
	for _, c := range cases {
		w := bytes.NewBuffer(nil)
		cr_0004_knight_move(bytes.NewBufferString(c.input), w)
		result, err := io.ReadAll(w)
		assert.NoError(t, err)
		assert.Equal(t, c.output, string(result))
	}
}

// func main() {
//   cr_0004_knight_move(os.Stdin, os.Stdout)
// }

// Формируем поле в виде двумерного целочисленного массива. Итерируемся слева-направо сверху-вниз по клеткам поля
// и прибавляем к значению каждой клетки те значения, которые лежат на клетках, получающихся при двух вариантах
// движения коня "назад" из текущей клетки (т.е. один шаг влево и два вверх или два шага влево и один вверх).
// При этом в самую первую клетку помещаем единицу. Таким образом мы определим, сколько единиц "дошло" до нижней правой клетки.
func cr_0004_knight_move(r io.Reader, w io.Writer) {
	reader := bufio.NewReaderSize(r, 1<<20)
	writer := bufio.NewWriterSize(w, 1<<20)
	defer writer.Flush()

	var input [2]int
	for i := range input {
		b, _ := reader.ReadBytes(' ')
		input[i], _ = strconv.Atoi(string(b[:len(b)-1]))
	}

	board := make([][]int, input[0])
	for i := range board {
		board[i] = make([]int, input[1])
	}

	board[0][0] = 1
	for i, row := range board {
		for j := range row {
			if i > 0 && j > 1 {
				board[i][j] += board[i-1][j-2]
			}
			if i > 1 && j > 0 {
				board[i][j] += board[i-2][j-1]
			}
		}
	}
	writer.WriteString(strconv.Itoa(board[input[0]-1][input[1]-1]))
}

// Решение через полный перебор без мемоизации. Работает, но не проходит по времени.
// func cr_0004_knight_move(r io.Reader, w io.Writer) {
// 	reader := bufio.NewReaderSize(r, 1<<20)
// 	writer := bufio.NewWriterSize(w, 1<<20)
// 	defer writer.Flush()

// 	var input [2]int
// 	for i := range input {
// 		b, _ := reader.ReadBytes(' ')
// 		input[i], _ = strconv.Atoi(string(b[:len(b)-1]))
// 	}
// 	writer.WriteString(strconv.Itoa(cr_0004_rec(1, 1, input[0], input[1])))
// }

// func cr_0004_rec(i, j, m, n int) int {
// 	switch {
// 	case i == m && j == n:
// 		return 1
// 	case i > m || j > n:
// 		return 0
// 	}
// 	return cr_0004_rec(i+1, j+2, m, n) + cr_0004_rec(i+2, j+1, m, n)
// }
