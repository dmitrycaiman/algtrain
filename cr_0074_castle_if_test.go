package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://coderun.yandex.ru/problem/castle-if
func Test_cr_0074_castle_if(t *testing.T) {
	cases := []struct{ input, output string }{
		{
			`1
1
1
1
1`,
			`YES`,
		},
		{
			`2
2
2
1
1`,
			`NO`,
		},
		{
			`1
5
2
5
2`,
			`YES`,
		},
		{
			`2
5
2
6
1`,
			`NO`,
		},
	}
	for _, c := range cases {
		w := bytes.NewBuffer(nil)
		cr_0074_castle_if(bytes.NewBufferString(c.input), w)
		result, err := io.ReadAll(w)
		assert.NoError(t, err)
		assert.Equal(t, c.output, string(result))
	}
}

// func main() {
//   $(os.Stdin, os.Stdout)
// }

// Среди сторон кирпича ищем такую, размеры которой меньше или равны отверстию в стене, независимо от её угла поворота.
// Соответственно ищем во входных данных два таких размера a и b, что относительно размеров дыры m, n:
// a <= m и b <= n или b <= m и a <= n.
func cr_0074_castle_if(r io.Reader, w io.Writer) {
	reader := bufio.NewReaderSize(r, 1<<20)
	writer := bufio.NewWriterSize(w, 1<<20)
	defer writer.Flush()

	input := [5]int{}
	for i := range input {
		line, _, _ := reader.ReadLine()
		input[i], _ = strconv.Atoi(string(line))
	}
	var a, b int
	for i := range 3 {
		switch i {
		case 0:
			a, b = input[0], input[1]
		case 1:
			a, b = input[0], input[2]
		case 2:
			a, b = input[1], input[2]
		}
		if (a <= input[3] && b <= input[4]) || (a <= input[4] && b <= input[3]) {
			writer.WriteString("YES")
			return
		}
	}
	writer.WriteString("NO")
}
